// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {ERC20Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import {ERC20PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20PausableUpgradeable.sol";
import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

import {IWattUSD} from "../interfaces/IWattUSD.sol";

/// @title WattUSD
/// @notice Synthetic USD token pegged 1:1 to USDC/USDT.
///         Minted exclusively by MintEngine. Burns on redemption.
///         6 decimals — matches the underlying stablecoins.
///
/// @dev UUPS upgradeable proxy pattern. All upgrades require UPGRADER_ROLE
///      (assigned to the Timelock after governance is deployed in Phase 6).
///      Storage layout must never be changed — only append new variables.
contract WattUSD is
    Initializable,
    ERC20Upgradeable,
    ERC20PausableUpgradeable,
    AccessControlUpgradeable,
    UUPSUpgradeable,
    IWattUSD
{
    // ── Roles ─────────────────────────────────────────────────────────────────

    /// @notice Allows minting and burning. Granted exclusively to MintEngine.
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

    /// @notice Allows pausing/unpausing transfers.
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    /// @notice Allows authorizing contract upgrades.
    ///         Should be transferred to Timelock once governance is live.
    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");

    // ── Constructor ───────────────────────────────────────────────────────────

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    // ── Initializer ───────────────────────────────────────────────────────────

    /// @notice Initializes the proxy. Called once by the deploy script.
    /// @param admin Address granted DEFAULT_ADMIN_ROLE, PAUSER_ROLE, and UPGRADER_ROLE.
    ///              Should be a multisig or Timelock in production.
    function initialize(address admin) external initializer {
        if (admin == address(0)) revert ZeroAddress();

        __ERC20_init("WattUSD", "WATT");
        __ERC20Pausable_init();
        __AccessControl_init();

        _grantRole(DEFAULT_ADMIN_ROLE, admin);
        _grantRole(PAUSER_ROLE, admin);
        _grantRole(UPGRADER_ROLE, admin);
        // MINTER_ROLE is NOT granted here — granted to MintEngine post-deploy
    }

    // ── IWattUSD ──────────────────────────────────────────────────────────────

    /// @inheritdoc IWattUSD
    function mint(address to, uint256 amount) external onlyRole(MINTER_ROLE) {
        if (amount == 0) revert ZeroAmount();
        _mint(to, amount);
        emit WattMinted(to, amount);
    }

    /// @inheritdoc IWattUSD
    function burn(address from, uint256 amount) external onlyRole(MINTER_ROLE) {
        if (amount == 0) revert ZeroAmount();
        _burn(from, amount);
        emit WattBurned(from, amount);
    }

    /// @inheritdoc IWattUSD
    function pause() external onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /// @inheritdoc IWattUSD
    function unpause() external onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    // ── ERC-20 overrides ──────────────────────────────────────────────────────

    /// @notice Returns 6 decimals to match USDC and USDT.
    function decimals() public pure override(ERC20Upgradeable, IWattUSD) returns (uint8) {
        return 6;
    }

    // ── Internal ──────────────────────────────────────────────────────────────

    /// @dev Required by OZ v5 to resolve diamond inheritance between
    ///      ERC20Upgradeable and ERC20PausableUpgradeable.
    function _update(
        address from,
        address to,
        uint256 value
    ) internal override(ERC20Upgradeable, ERC20PausableUpgradeable) {
        super._update(from, to, value);
    }

    /// @dev Only UPGRADER_ROLE (Timelock in production) can authorize an upgrade.
    function _authorizeUpgrade(address newImplementation)
        internal
        override
        onlyRole(UPGRADER_ROLE)
    {}
}
