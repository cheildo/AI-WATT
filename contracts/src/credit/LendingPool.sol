// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import {ILendingPool} from "../interfaces/ILendingPool.sol";
import {IAssetRegistry} from "../interfaces/IAssetRegistry.sol";
import {IHealthAttestation} from "../interfaces/IHealthAttestation.sol";
import {IWattUSD} from "../interfaces/IWattUSD.sol";
import {IsWattUSD} from "../interfaces/IsWattUSD.sol";

/// @title LendingPool
/// @notice Engine 2 on-chain loan lifecycle for post-delivery productivity-backed loans.
///
///         Pre-origination checks:
///           1. Asset must be ACTIVE in AssetRegistry.
///           2. Latest HealthAttestation score >= 60 and submitted within 48 hours.
///           3. Asset must not already have an active loan.
///
///         On repayment, the interest portion is split:
///           - 90% routed to sWattUSD.receiveYield() → staker NAV rises.
///           - 10% accumulated as protocol fees → withdrawn by ADMIN_ROLE.
///
///         Principal is minted to the borrower at origination via MINTER_ROLE on WattUSD.
///         Interest is collected as WATT from the repayer.
///
/// @dev UUPS upgradeable. Storage layout — only append new variables.
contract LendingPool is
    Initializable,
    AccessControlUpgradeable,
    PausableUpgradeable,
    UUPSUpgradeable,
    ILendingPool
{
    using SafeERC20 for IERC20;

    // ── Roles ─────────────────────────────────────────────────────────────────

    /// @notice Loan originators (backend curator wallets or approved institutions).
    bytes32 public constant CURATOR_ROLE = keccak256("CURATOR_ROLE");

    /// @notice Allowed to liquidate DEFAULTED loans.
    bytes32 public constant LIQUIDATOR_ROLE = keccak256("LIQUIDATOR_ROLE");

    /// @notice Protocol admin — withdraw fees, pause, update parameters.
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    /// @notice Allows pausing the pool.
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    /// @notice Allows authorizing contract upgrades.
    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");

    // ── Constants ─────────────────────────────────────────────────────────────

    /// @inheritdoc ILendingPool
    uint8 public constant MIN_HEALTH_SCORE = 60;

    /// @inheritdoc ILendingPool
    uint256 public constant ATTESTATION_MAX_AGE = 48 hours;

    /// @notice Protocol's share of interest income in basis points (10%).
    uint256 public constant PROTOCOL_FEE_BPS = 1000;

    // ── Storage ───────────────────────────────────────────────────────────────

    /// @notice AssetRegistry contract — checked for active collateral.
    IAssetRegistry public assetRegistry;

    /// @notice HealthAttestation contract — checked for fresh, healthy attestation.
    IHealthAttestation public healthAttestation;

    /// @notice WattUSD token — minted to borrower at origination.
    address public wattUSD;

    /// @notice sWattUSD vault — receives yield portion of repayments.
    address public sWattUSD;

    mapping(bytes32 => Loan) private _loans;
    mapping(bytes32 => bool) private _loanExists;
    mapping(address => bytes32[]) private _borrowerLoans;

    /// @dev Prevents double-encumbering the same asset collateral.
    mapping(bytes32 => bool) private _assetHasActiveLoan;

    /// @dev Accumulated protocol fees (WATT, 6 decimals) pending withdrawal.
    uint256 private _protocolFees;

    /// @dev Inline reentrancy guard: 1 = not entered, 2 = entered.
    uint8 private _reentrancyStatus;

    error ReentrantCall();

    // ── Constructor ───────────────────────────────────────────────────────────

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    // ── Initializer ───────────────────────────────────────────────────────────

    /// @param admin             Address granted all admin roles.
    /// @param assetRegistry_    AssetRegistry proxy address.
    /// @param healthAttestation_ HealthAttestation proxy address.
    /// @param wattUSD_          WattUSD proxy address.
    /// @param sWattUSD_         sWattUSD proxy address.
    function initialize(
        address admin,
        address assetRegistry_,
        address healthAttestation_,
        address wattUSD_,
        address sWattUSD_
    ) external initializer {
        if (admin == address(0) || assetRegistry_ == address(0) ||
            healthAttestation_ == address(0) || wattUSD_ == address(0) ||
            sWattUSD_ == address(0)) revert ZeroAddress();

        __AccessControl_init();
        __Pausable_init();
        _reentrancyStatus = 1;

        _grantRole(DEFAULT_ADMIN_ROLE, admin);
        _grantRole(ADMIN_ROLE, admin);
        _grantRole(PAUSER_ROLE, admin);
        _grantRole(UPGRADER_ROLE, admin);

        assetRegistry    = IAssetRegistry(assetRegistry_);
        healthAttestation = IHealthAttestation(healthAttestation_);
        wattUSD          = wattUSD_;
        sWattUSD         = sWattUSD_;
    }

    // ── Modifiers ─────────────────────────────────────────────────────────────

    modifier nonReentrant() {
        if (_reentrancyStatus == 2) revert ReentrantCall();
        _reentrancyStatus = 2;
        _;
        _reentrancyStatus = 1;
    }

    // ── ILendingPool — write ──────────────────────────────────────────────────

    /// @inheritdoc ILendingPool
    function originateLoan(
        bytes32 assetId,
        address borrower,
        uint256 principal,
        uint256 interestRate,
        uint256 termDays,
        uint8   engineType
    ) external onlyRole(CURATOR_ROLE) nonReentrant whenNotPaused returns (bytes32 loanId) {
        if (borrower == address(0)) revert ZeroAddress();
        if (principal == 0 || termDays == 0) revert ZeroAmount();

        // ── Pre-origination checks ────────────────────────────────────────────

        if (!assetRegistry.isActive(assetId)) revert AssetNotActive(assetId);

        if (_assetHasActiveLoan[assetId]) revert AssetAlreadyEncumbered(assetId);

        IHealthAttestation.Attestation memory att = healthAttestation.getLatestAttestation(assetId);
        if (att.timestamp == 0 || block.timestamp - att.timestamp > ATTESTATION_MAX_AGE) {
            revert AttestationStale(assetId, att.timestamp);
        }
        if (att.score < MIN_HEALTH_SCORE) {
            revert HealthScoreTooLow(assetId, att.score);
        }

        // ── Create loan record ────────────────────────────────────────────────

        loanId = keccak256(abi.encodePacked(assetId, borrower, block.timestamp, msg.sender));

        uint256 maturityAt = block.timestamp + termDays * 1 days;
        uint256 totalInterest = _calcInterest(principal, interestRate, termDays);
        uint256 outstanding = principal + totalInterest;

        _loans[loanId] = Loan({
            loanId:       loanId,
            assetId:      assetId,
            borrower:     borrower,
            curator:      msg.sender,
            principal:    principal,
            outstanding:  outstanding,
            interestRate: interestRate,
            status:       LoanStatus.ACTIVE,
            engineType:   engineType,
            originatedAt: block.timestamp,
            maturityAt:   maturityAt
        });
        _loanExists[loanId] = true;
        _borrowerLoans[borrower].push(loanId);
        _assetHasActiveLoan[assetId] = true;

        // ── Disburse: mint WATT to borrower ──────────────────────────────────

        IWattUSD(wattUSD).mint(borrower, principal);

        emit LoanOriginated(loanId, assetId, borrower, msg.sender, principal, interestRate, maturityAt);
    }

    /// @inheritdoc ILendingPool
    function repay(bytes32 loanId, uint256 amount) external nonReentrant whenNotPaused {
        if (!_loanExists[loanId]) revert LoanNotFound(loanId);
        if (amount == 0) revert ZeroAmount();

        Loan storage loan = _loans[loanId];
        if (loan.status != LoanStatus.ACTIVE && loan.status != LoanStatus.REPAYING) {
            revert LoanNotActive(loanId, loan.status);
        }
        if (amount > loan.outstanding) revert ExceedsOutstanding(amount, loan.outstanding);

        // ── Split repayment into principal + interest ─────────────────────────

        uint256 termDays = (loan.maturityAt - loan.originatedAt) / 1 days;
        uint256 totalInterest = _calcInterest(loan.principal, loan.interestRate, termDays);
        uint256 totalDebt = loan.principal + totalInterest;

        // Pro-rata: interestPaid = amount * totalInterest / totalDebt
        uint256 interestPaid  = totalInterest > 0 ? (amount * totalInterest / totalDebt) : 0;
        uint256 principalPaid = amount - interestPaid;

        // Protocol fee: 10% of interest
        uint256 protocolFee = interestPaid * PROTOCOL_FEE_BPS / 10_000;
        uint256 yieldAmount = interestPaid - protocolFee;

        // ── Pull WATT from repayer ────────────────────────────────────────────

        IERC20(wattUSD).safeTransferFrom(msg.sender, address(this), amount);

        // ── Route yield to sWattUSD ───────────────────────────────────────────

        if (yieldAmount > 0) {
            IERC20(wattUSD).forceApprove(sWattUSD, yieldAmount);
            IsWattUSD(sWattUSD).receiveYield(yieldAmount);
        }

        // Accumulate protocol fee (stays in contract until withdrawn)
        _protocolFees += protocolFee;

        // ── Update loan state ─────────────────────────────────────────────────

        loan.outstanding -= amount;
        if (loan.status == LoanStatus.ACTIVE) loan.status = LoanStatus.REPAYING;

        emit RepaymentReceived(
            loanId, msg.sender, amount, principalPaid, interestPaid, loan.outstanding
        );

        // Auto-settle when fully repaid
        if (loan.outstanding == 0) {
            _settle(loanId, loan);
        }
    }

    /// @inheritdoc ILendingPool
    function fullRepay(bytes32 loanId) external nonReentrant whenNotPaused {
        if (!_loanExists[loanId]) revert LoanNotFound(loanId);
        Loan storage loan = _loans[loanId];
        if (loan.status != LoanStatus.ACTIVE && loan.status != LoanStatus.REPAYING) {
            revert LoanNotActive(loanId, loan.status);
        }

        uint256 outstanding = loan.outstanding;
        if (outstanding == 0) revert ZeroAmount();

        // Compute split for the full outstanding balance
        uint256 termDays = (loan.maturityAt - loan.originatedAt) / 1 days;
        uint256 totalInterest = _calcInterest(loan.principal, loan.interestRate, termDays);
        uint256 totalDebt = loan.principal + totalInterest;

        uint256 interestPaid  = totalInterest > 0 ? (outstanding * totalInterest / totalDebt) : 0;
        uint256 protocolFee   = interestPaid * PROTOCOL_FEE_BPS / 10_000;
        uint256 yieldAmount   = interestPaid - protocolFee;
        uint256 principalPaid = outstanding - interestPaid;

        IERC20(wattUSD).safeTransferFrom(msg.sender, address(this), outstanding);

        if (yieldAmount > 0) {
            IERC20(wattUSD).forceApprove(sWattUSD, yieldAmount);
            IsWattUSD(sWattUSD).receiveYield(yieldAmount);
        }

        _protocolFees += protocolFee;

        loan.outstanding = 0;
        loan.status      = LoanStatus.REPAYING; // will be overwritten by _settle

        emit RepaymentReceived(
            loanId, msg.sender, outstanding, principalPaid, interestPaid, 0
        );

        _settle(loanId, loan);
    }

    /// @inheritdoc ILendingPool
    /// @dev Callable by anyone — no role restriction. Enforces maturity deadline.
    function flagDefaulted(bytes32 loanId) external {
        if (!_loanExists[loanId]) revert LoanNotFound(loanId);
        Loan storage loan = _loans[loanId];
        if (loan.status != LoanStatus.ACTIVE && loan.status != LoanStatus.REPAYING) {
            revert LoanNotActive(loanId, loan.status);
        }
        if (block.timestamp <= loan.maturityAt) revert LoanNotActive(loanId, loan.status);

        loan.status = LoanStatus.DEFAULTED;
        emit LoanDefaulted(loanId, loan.borrower);
    }

    /// @inheritdoc ILendingPool
    function liquidate(bytes32 loanId) external onlyRole(LIQUIDATOR_ROLE) nonReentrant {
        if (!_loanExists[loanId]) revert LoanNotFound(loanId);
        Loan storage loan = _loans[loanId];
        if (loan.status != LoanStatus.DEFAULTED) revert LoanNotActive(loanId, loan.status);

        bytes32 assetId = loan.assetId;
        loan.status = LoanStatus.LIQUIDATED;
        _assetHasActiveLoan[assetId] = false;

        assetRegistry.updateStatus(assetId, IAssetRegistry.AssetStatus.LIQUIDATED);

        emit LoanLiquidated(loanId, loan.borrower, msg.sender);
    }

    /// @inheritdoc ILendingPool
    function withdrawFees(address to) external onlyRole(ADMIN_ROLE) nonReentrant {
        if (to == address(0)) revert ZeroAddress();
        uint256 fees = _protocolFees;
        if (fees == 0) revert NoFeesAvailable();
        _protocolFees = 0;
        IERC20(wattUSD).safeTransfer(to, fees);
        emit FeesWithdrawn(to, fees);
    }

    // ── Admin ─────────────────────────────────────────────────────────────────

    function pause() external onlyRole(PAUSER_ROLE) {
        _pause();
    }

    function unpause() external onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    // ── ILendingPool — view ───────────────────────────────────────────────────

    /// @inheritdoc ILendingPool
    function getLoan(bytes32 loanId) external view returns (Loan memory) {
        if (!_loanExists[loanId]) revert LoanNotFound(loanId);
        return _loans[loanId];
    }

    /// @inheritdoc ILendingPool
    function getBorrowerLoans(address borrower) external view returns (bytes32[] memory) {
        return _borrowerLoans[borrower];
    }

    /// @inheritdoc ILendingPool
    function getProtocolFees() external view returns (uint256) {
        return _protocolFees;
    }

    // ── Internal ──────────────────────────────────────────────────────────────

    /// @dev Simple annual interest: principal × rate(bps) × days / (365 × 10000).
    function _calcInterest(
        uint256 principal,
        uint256 interestRate,
        uint256 termDays
    ) internal pure returns (uint256) {
        return principal * interestRate * termDays / (365 * 10_000);
    }

    /// @dev Marks loan as SETTLED, clears asset lock, resets asset status to ACTIVE.
    function _settle(bytes32 loanId, Loan storage loan) internal {
        loan.status = LoanStatus.SETTLED;
        _assetHasActiveLoan[loan.assetId] = false;
        // Reset asset to ACTIVE in case it was flagged during the loan term
        assetRegistry.updateStatus(loan.assetId, IAssetRegistry.AssetStatus.ACTIVE);
        emit LoanSettled(loanId, loan.borrower);
    }

    /// @dev Only UPGRADER_ROLE can authorize an upgrade.
    function _authorizeUpgrade(address newImplementation)
        internal
        override
        onlyRole(UPGRADER_ROLE)
    {}
}
