// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

import {IAssetRegistry} from "../interfaces/IAssetRegistry.sol";

/// @title AssetRegistry
/// @notice On-chain registry of physical assets approved as collateral in the AI WATT
///         protocol. Stores asset metadata, LTV ratios, and lifecycle status.
///
///         REGISTRAR_ROLE (backend signer) creates records.
///         LENDINGPOOL_ROLE (Phase 4) and ADMIN_ROLE update status during the loan lifecycle.
///         LendingPool calls isActive() and getAsset() before originating loans.
///
/// @dev UUPS upgradeable. Storage layout — only append new variables.
contract AssetRegistry is
    Initializable,
    AccessControlUpgradeable,
    UUPSUpgradeable,
    IAssetRegistry
{
    // ── Roles ─────────────────────────────────────────────────────────────────

    /// @notice Backend signer wallet — allowed to register new assets.
    bytes32 public constant REGISTRAR_ROLE = keccak256("REGISTRAR_ROLE");

    /// @notice LendingPool contract — allowed to update asset status during loan lifecycle.
    bytes32 public constant LENDINGPOOL_ROLE = keccak256("LENDINGPOOL_ROLE");

    /// @notice Protocol admin — can update LTV, status, and grant roles.
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    /// @notice Allows authorizing contract upgrades.
    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");

    // ── Constants ─────────────────────────────────────────────────────────────

    /// @notice Maximum allowed LTV in basis points (90%).
    uint16 public constant MAX_LTV = 9000;

    // ── Storage ───────────────────────────────────────────────────────────────

    mapping(bytes32 => Asset) private _assets;
    mapping(bytes32 => bool)  private _registered;

    // ── Constructor ───────────────────────────────────────────────────────────

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    // ── Initializer ───────────────────────────────────────────────────────────

    /// @param admin Address granted all admin roles. Should be a multisig in production.
    function initialize(address admin) external initializer {
        if (admin == address(0)) revert ZeroAddress();
        __AccessControl_init();
        _grantRole(DEFAULT_ADMIN_ROLE, admin);
        _grantRole(ADMIN_ROLE, admin);
        _grantRole(UPGRADER_ROLE, admin);
    }

    // ── IAssetRegistry — write ────────────────────────────────────────────────

    /// @inheritdoc IAssetRegistry
    function registerAsset(
        bytes32   assetId,
        AssetType assetType,
        address   borrower,
        uint16    ltv
    ) external onlyRole(REGISTRAR_ROLE) {
        if (_registered[assetId]) revert AssetAlreadyRegistered(assetId);
        if (borrower == address(0)) revert ZeroAddress();
        if (ltv == 0 || ltv > MAX_LTV) revert InvalidLTV(ltv);

        _assets[assetId] = Asset({
            assetId:      assetId,
            assetType:    assetType,
            borrower:     borrower,
            ltv:          ltv,
            status:       AssetStatus.PENDING,
            registeredAt: block.timestamp
        });
        _registered[assetId] = true;

        emit AssetRegistered(assetId, assetType, borrower, ltv);
    }

    /// @inheritdoc IAssetRegistry
    function updateLTV(bytes32 assetId, uint16 newLTV) external onlyRole(ADMIN_ROLE) {
        if (!_registered[assetId]) revert AssetNotFound(assetId);
        if (newLTV == 0 || newLTV > MAX_LTV) revert InvalidLTV(newLTV);
        uint16 old = _assets[assetId].ltv;
        _assets[assetId].ltv = newLTV;
        emit LTVUpdated(assetId, old, newLTV);
    }

    /// @inheritdoc IAssetRegistry
    /// @dev Callable by LENDINGPOOL_ROLE (automated loan lifecycle) or ADMIN_ROLE (manual override).
    function updateStatus(bytes32 assetId, AssetStatus newStatus) external {
        if (!_registered[assetId]) revert AssetNotFound(assetId);
        if (!hasRole(LENDINGPOOL_ROLE, msg.sender) && !hasRole(ADMIN_ROLE, msg.sender)) {
            revert AccessControlUnauthorizedAccount(msg.sender, LENDINGPOOL_ROLE);
        }
        AssetStatus old = _assets[assetId].status;
        _assets[assetId].status = newStatus;
        emit StatusChanged(assetId, old, newStatus);
    }

    // ── IAssetRegistry — view ─────────────────────────────────────────────────

    /// @inheritdoc IAssetRegistry
    function getAsset(bytes32 assetId) external view returns (Asset memory) {
        if (!_registered[assetId]) revert AssetNotFound(assetId);
        return _assets[assetId];
    }

    /// @inheritdoc IAssetRegistry
    function isActive(bytes32 assetId) external view returns (bool) {
        return _registered[assetId] && _assets[assetId].status == AssetStatus.ACTIVE;
    }

    // ── Internal ──────────────────────────────────────────────────────────────

    /// @dev Only UPGRADER_ROLE can authorize an upgrade.
    function _authorizeUpgrade(address newImplementation)
        internal
        override
        onlyRole(UPGRADER_ROLE)
    {}
}
