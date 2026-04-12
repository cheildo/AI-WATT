// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import {IWattUSD} from "../interfaces/IWattUSD.sol";
import {IMintEngine} from "../interfaces/IMintEngine.sol";

/// @title MintEngine
/// @notice Entry point for WATT minting and redemption.
///
///         Deposit flow:
///           user deposits `amount` USDC → fee deducted → net WATT minted 1:1
///
///         Redemption flow:
///           user burns `wattAmount` WATT → fee deducted → net USDC returned
///
///         Fee: 0.1% (10 BPS) on both sides, sent to treasury.
///         Collateral: held in this contract; Engine 3 sweeps idle capital to T-bills (Phase 11).
///
/// @dev UUPS upgradeable. Storage layout — only append new variables after existing ones.
contract MintEngine is
    Initializable,
    AccessControlUpgradeable,
    PausableUpgradeable,
    UUPSUpgradeable,
    IMintEngine
{
    using SafeERC20 for IERC20;

    // ── Roles ─────────────────────────────────────────────────────────────────

    /// @notice Can update accepted stablecoins, treasury address, and protocol params.
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    /// @notice Allows pausing the engine.
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    /// @notice Allows authorizing contract upgrades.
    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");

    // ── Constants ─────────────────────────────────────────────────────────────

    /// @inheritdoc IMintEngine
    uint256 public constant FEE_BPS = 10; // 0.1%

    uint256 private constant BPS_DENOMINATOR = 10_000;

    // ── Storage ───────────────────────────────────────────────────────────────

    /// @notice WattUSD token contract — mint/burn called here.
    IWattUSD public wattUSD;

    /// @notice Address that receives protocol fees.
    address public treasury;

    /// @notice Stablecoins accepted for minting (USDC, USDT).
    mapping(address => bool) public acceptedStablecoins;

    // ── Constructor ───────────────────────────────────────────────────────────

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    // ── Initializer ───────────────────────────────────────────────────────────

    /// @notice Initializes the proxy. Called once by the deploy script.
    /// @param admin      Address granted ADMIN_ROLE, PAUSER_ROLE, UPGRADER_ROLE.
    /// @param _wattUSD   Deployed WattUSD proxy address.
    /// @param _treasury  Address that receives mint/redeem fees.
    function initialize(
        address admin,
        address _wattUSD,
        address _treasury
    ) external initializer {
        if (admin == address(0) || _wattUSD == address(0) || _treasury == address(0)) {
            revert ZeroAddress();
        }

        __AccessControl_init();
        __Pausable_init();

        _grantRole(DEFAULT_ADMIN_ROLE, admin);
        _grantRole(ADMIN_ROLE, admin);
        _grantRole(PAUSER_ROLE, admin);
        _grantRole(UPGRADER_ROLE, admin);

        wattUSD = IWattUSD(_wattUSD);
        treasury = _treasury;
    }

    // ── IMintEngine — core ────────────────────────────────────────────────────

    /// @inheritdoc IMintEngine
    /// @dev Caller must have approved this contract to spend `amount` of `stablecoin`.
    function mint(address stablecoin, uint256 amount) external whenNotPaused {
        if (amount == 0) revert ZeroAmount();
        if (!acceptedStablecoins[stablecoin]) revert StablecoinNotAccepted(stablecoin);

        uint256 fee = (amount * FEE_BPS) / BPS_DENOMINATOR;
        uint256 wattToMint = amount - fee;

        // Pull stablecoin from caller — full gross amount stays in contract as collateral
        IERC20(stablecoin).safeTransferFrom(msg.sender, address(this), amount);

        // Forward fee to treasury
        if (fee > 0) {
            IERC20(stablecoin).safeTransfer(treasury, fee);
        }

        // Mint net WATT to caller (1:1 with net stablecoin held as collateral)
        wattUSD.mint(msg.sender, wattToMint);

        emit Minted(msg.sender, stablecoin, amount, wattToMint, fee);
    }

    /// @inheritdoc IMintEngine
    /// @dev Caller must hold at least `wattAmount` WATT.
    ///      Collateral for `stablecoin` must be sufficient in this contract.
    function redeem(address stablecoin, uint256 wattAmount) external whenNotPaused {
        if (wattAmount == 0) revert ZeroAmount();
        if (!acceptedStablecoins[stablecoin]) revert StablecoinNotAccepted(stablecoin);

        uint256 fee = (wattAmount * FEE_BPS) / BPS_DENOMINATOR;
        uint256 stablecoinToReturn = wattAmount - fee;

        uint256 available = IERC20(stablecoin).balanceOf(address(this));
        if (available < stablecoinToReturn) {
            revert InsufficientCollateral(stablecoin, stablecoinToReturn, available);
        }

        // Burn WATT from caller first (checks-effects-interactions)
        wattUSD.burn(msg.sender, wattAmount);

        // Return net stablecoin to caller
        IERC20(stablecoin).safeTransfer(msg.sender, stablecoinToReturn);

        // Forward fee to treasury
        if (fee > 0) {
            IERC20(stablecoin).safeTransfer(treasury, fee);
        }

        emit Redeemed(msg.sender, stablecoin, wattAmount, stablecoinToReturn, fee);
    }

    // ── IMintEngine — admin ───────────────────────────────────────────────────

    /// @inheritdoc IMintEngine
    function setAcceptedStablecoin(address stablecoin, bool accepted)
        external
        onlyRole(ADMIN_ROLE)
    {
        if (stablecoin == address(0)) revert ZeroAddress();
        acceptedStablecoins[stablecoin] = accepted;
        emit StablecoinUpdated(stablecoin, accepted);
    }

    /// @inheritdoc IMintEngine
    function setTreasury(address newTreasury) external onlyRole(ADMIN_ROLE) {
        if (newTreasury == address(0)) revert ZeroAddress();
        address old = treasury;
        treasury = newTreasury;
        emit TreasuryUpdated(old, newTreasury);
    }

    /// @notice Pause minting and redemption. Callable by PAUSER_ROLE.
    function pause() external onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /// @notice Unpause minting and redemption. Callable by PAUSER_ROLE.
    function unpause() external onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    // ── IMintEngine — view ────────────────────────────────────────────────────

    /// @inheritdoc IMintEngine
    function collateralBalance(address stablecoin) external view returns (uint256) {
        return IERC20(stablecoin).balanceOf(address(this));
    }

    /// @inheritdoc IMintEngine
    function isAcceptedStablecoin(address stablecoin) external view returns (bool) {
        return acceptedStablecoins[stablecoin];
    }

    // ── Internal ──────────────────────────────────────────────────────────────

    /// @dev Only UPGRADER_ROLE can authorize an upgrade.
    function _authorizeUpgrade(address newImplementation)
        internal
        override
        onlyRole(UPGRADER_ROLE)
    {}
}
