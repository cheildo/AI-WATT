// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {IERC4626} from "@openzeppelin/contracts/interfaces/IERC4626.sol";

import {IWEVQueue} from "../interfaces/IWEVQueue.sol";

/// @title WEVQueue
/// @notice sWATT redemption queue for large withdrawals that exceed the sWattUSD threshold.
///
///         Flow:
///           1. User calls requestRedeem or requestPriorityRedeem — sWATT transferred to
///              this contract, request created in QUEUED state.
///           2. User may cancel while QUEUED: sWATT (and any priority fee) returned.
///           3. Backend keeper (PROCESSOR_ROLE) calls processBatch with a list of requestIds
///              in desired processing order. For each: sWATT redeemed via sWattUSD,
///              WATT delivered to the user, request marked FULFILLED.
///
///         Priority queue:
///           - User pays >= 0.5% of sWattAmount in WATT as a priority fee.
///           - Fee held in contract; added to _protocolFees on fulfillment, refunded on cancel.
///           - Off-chain keeper respects priority ordering; the contract enforces fee payment only.
///
///         sWattUSD integration:
///           - sWattUSD.setWEVQueue(address(this)) is called by the deploy script.
///           - sWattUSD bypasses its large-redemption guard when the owner/caller is this contract,
///             allowing WEVQueue to redeem any amount on behalf of queued users.
///
/// @dev UUPS upgradeable. Storage layout — only append new variables.
contract WEVQueue is
    Initializable,
    AccessControlUpgradeable,
    PausableUpgradeable,
    UUPSUpgradeable,
    IWEVQueue
{
    using SafeERC20 for IERC20;

    // ── Roles ─────────────────────────────────────────────────────────────────

    /// @notice Backend keeper wallet — authorized to call processBatch.
    bytes32 public constant PROCESSOR_ROLE = keccak256("PROCESSOR_ROLE");

    /// @notice Protocol admin — withdraw fees, pause, update parameters.
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    /// @notice Allows pausing/unpausing the queue.
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    /// @notice Allows authorizing contract upgrades.
    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");

    // ── Constants ─────────────────────────────────────────────────────────────

    /// @inheritdoc IWEVQueue
    uint256 public constant PRIORITY_FEE_BPS = 50;          // 0.5%

    /// @inheritdoc IWEVQueue
    uint256 public constant STANDARD_WAIT = 30 days;

    /// @inheritdoc IWEVQueue
    uint256 public constant PRIORITY_WAIT = 3 days;

    // ── Storage ───────────────────────────────────────────────────────────────

    /// @notice sWattUSD vault — sWATT shares are held here and redeemed on fulfillment.
    address public sWattUSD;

    /// @notice WattUSD token — priority fees held and disbursed on fulfillment/cancel.
    address public wattUSD;

    mapping(bytes32 => RedemptionRequest) private _requests;
    mapping(bytes32 => bool) private _requestExists;
    mapping(address => bytes32[]) private _userRequests;

    /// @dev Count of currently QUEUED (active) requests.
    uint256 private _queueDepth;

    /// @dev Accumulated priority fees (WATT, 6 decimals) pending withdrawal.
    uint256 private _protocolFees;

    /// @dev Nonce for request ID uniqueness across same block / same user / same amount.
    uint256 private _requestNonce;

    /// @dev Inline reentrancy guard: 1 = not entered, 2 = entered.
    uint8 private _reentrancyStatus;

    error ReentrantCall();

    // ── Constructor ───────────────────────────────────────────────────────────

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    // ── Initializer ───────────────────────────────────────────────────────────

    /// @param admin     Address granted all admin roles.
    /// @param sWattUSD_ sWattUSD proxy address.
    /// @param wattUSD_  WattUSD proxy address.
    function initialize(
        address admin,
        address sWattUSD_,
        address wattUSD_
    ) external initializer {
        if (admin == address(0) || sWattUSD_ == address(0) || wattUSD_ == address(0))
            revert ZeroAddress();

        __AccessControl_init();
        __Pausable_init();
        _reentrancyStatus = 1;

        _grantRole(DEFAULT_ADMIN_ROLE, admin);
        _grantRole(ADMIN_ROLE, admin);
        _grantRole(PAUSER_ROLE, admin);
        _grantRole(UPGRADER_ROLE, admin);

        sWattUSD = sWattUSD_;
        wattUSD  = wattUSD_;
    }

    // ── Modifiers ─────────────────────────────────────────────────────────────

    modifier nonReentrant() {
        if (_reentrancyStatus == 2) revert ReentrantCall();
        _reentrancyStatus = 2;
        _;
        _reentrancyStatus = 1;
    }

    // ── IWEVQueue — write ─────────────────────────────────────────────────────

    /// @inheritdoc IWEVQueue
    function requestRedeem(uint256 sWattAmount)
        external nonReentrant whenNotPaused returns (bytes32 requestId)
    {
        if (sWattAmount == 0) revert ZeroAmount();

        requestId = _nextRequestId();

        IERC20(sWattUSD).safeTransferFrom(msg.sender, address(this), sWattAmount);

        _requests[requestId] = RedemptionRequest({
            requestId:   requestId,
            user:        msg.sender,
            sWattAmount: sWattAmount,
            priorityFee: 0,
            requestedAt: block.timestamp,
            status:      RequestStatus.QUEUED
        });
        _requestExists[requestId] = true;
        _userRequests[msg.sender].push(requestId);
        _queueDepth++;

        emit RedemptionRequested(requestId, msg.sender, sWattAmount, false);
    }

    /// @inheritdoc IWEVQueue
    function requestPriorityRedeem(uint256 sWattAmount, uint256 priorityFee)
        external nonReentrant whenNotPaused returns (bytes32 requestId)
    {
        if (sWattAmount == 0) revert ZeroAmount();

        uint256 minFee = sWattAmount * PRIORITY_FEE_BPS / 10_000;
        if (priorityFee < minFee) revert InsufficientPriorityFee(priorityFee, minFee);

        requestId = _nextRequestId();

        IERC20(sWattUSD).safeTransferFrom(msg.sender, address(this), sWattAmount);
        IERC20(wattUSD).safeTransferFrom(msg.sender, address(this), priorityFee);

        _requests[requestId] = RedemptionRequest({
            requestId:   requestId,
            user:        msg.sender,
            sWattAmount: sWattAmount,
            priorityFee: priorityFee,
            requestedAt: block.timestamp,
            status:      RequestStatus.QUEUED
        });
        _requestExists[requestId] = true;
        _userRequests[msg.sender].push(requestId);
        _queueDepth++;

        emit RedemptionRequested(requestId, msg.sender, sWattAmount, true);
    }

    /// @inheritdoc IWEVQueue
    function cancelRequest(bytes32 requestId) external nonReentrant {
        if (!_requestExists[requestId]) revert RequestNotFound(requestId);

        RedemptionRequest storage req = _requests[requestId];
        if (req.user != msg.sender)               revert NotRequestOwner(requestId);
        if (req.status != RequestStatus.QUEUED)   revert NotQueued(requestId, req.status);

        req.status = RequestStatus.CANCELLED;
        _queueDepth--;

        // Return sWATT to user
        IERC20(sWattUSD).safeTransfer(req.user, req.sWattAmount);

        // Refund priority fee in WATT if the request was a priority request
        if (req.priorityFee > 0) {
            IERC20(wattUSD).safeTransfer(req.user, req.priorityFee);
        }

        emit RedemptionCancelled(requestId, req.user);
    }

    /// @inheritdoc IWEVQueue
    /// @dev sWattUSD bypasses its large-redemption guard when the caller/owner is this
    ///      contract, so batch sizes are not restricted by the WEV threshold.
    function processBatch(bytes32[] calldata requestIds)
        external onlyRole(PROCESSOR_ROLE) nonReentrant
    {
        uint256 len = requestIds.length;
        if (len == 0) revert NothingToProcess();

        for (uint256 i = 0; i < len; i++) {
            bytes32 requestId = requestIds[i];
            if (!_requestExists[requestId]) revert RequestNotFound(requestId);

            RedemptionRequest storage req = _requests[requestId];
            if (req.status != RequestStatus.QUEUED) revert NotQueued(requestId, req.status);

            req.status = RequestStatus.FULFILLED;
            _queueDepth--;

            // Accumulate priority fee as protocol revenue
            if (req.priorityFee > 0) {
                _protocolFees += req.priorityFee;
            }

            // Redeem sWATT shares owned by this contract → WATT sent directly to user.
            // WEVQueue is the owner of the shares; sWattUSD bypasses the WEV guard for us.
            uint256 wattOut = IERC4626(sWattUSD).redeem(req.sWattAmount, req.user, address(this));

            emit RedemptionFulfilled(requestId, req.user, wattOut);
        }

        emit BatchProcessed(len);
    }

    /// @inheritdoc IWEVQueue
    function withdrawFees(address to) external onlyRole(ADMIN_ROLE) nonReentrant {
        if (to == address(0)) revert ZeroAddress();
        uint256 fees = _protocolFees;
        if (fees == 0) revert NoFeesAvailable();
        _protocolFees = 0;
        IERC20(wattUSD).safeTransfer(to, fees);
        emit FeesWithdrawn(to, fees);
    }

    // ── Admin ─────────────────────────────────────────────────────────────────

    function pause()   external onlyRole(PAUSER_ROLE) { _pause(); }
    function unpause() external onlyRole(PAUSER_ROLE) { _unpause(); }

    // ── IWEVQueue — view ──────────────────────────────────────────────────────

    /// @inheritdoc IWEVQueue
    function getRequest(bytes32 requestId) external view returns (RedemptionRequest memory) {
        if (!_requestExists[requestId]) revert RequestNotFound(requestId);
        return _requests[requestId];
    }

    /// @inheritdoc IWEVQueue
    function getUserRequests(address user) external view returns (bytes32[] memory) {
        return _userRequests[user];
    }

    /// @inheritdoc IWEVQueue
    function getQueueDepth() external view returns (uint256) {
        return _queueDepth;
    }

    /// @inheritdoc IWEVQueue
    function nextProcessingTimestamp() external view returns (uint256) {
        return block.timestamp + STANDARD_WAIT;
    }

    /// @inheritdoc IWEVQueue
    function getProtocolFees() external view returns (uint256) {
        return _protocolFees;
    }

    // ── Internal ──────────────────────────────────────────────────────────────

    /// @dev Returns a unique request ID using a monotonically increasing nonce.
    function _nextRequestId() internal returns (bytes32) {
        return keccak256(abi.encodePacked(msg.sender, block.timestamp, _requestNonce++));
    }

    /// @dev Only UPGRADER_ROLE can authorize an upgrade.
    function _authorizeUpgrade(address newImplementation)
        internal override onlyRole(UPGRADER_ROLE)
    {}
}
