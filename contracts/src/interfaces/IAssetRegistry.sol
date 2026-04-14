// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title IAssetRegistry
/// @notice Registry of physical assets (GPU clusters, robotics, AI energy infrastructure)
///         approved as collateral in the AI WATT protocol.
interface IAssetRegistry {
    // ── Enums ─────────────────────────────────────────────────────────────────

    /// @notice Class of physical asset being collateralised.
    enum AssetType { GPU_CLUSTER, ROBOTICS, AI_ENERGY }

    /// @notice Lifecycle status of a registered asset.
    enum AssetStatus { PENDING, ACTIVE, FLAGGED, LIQUIDATED }

    // ── Structs ───────────────────────────────────────────────────────────────

    /// @notice On-chain record for a registered asset.
    struct Asset {
        bytes32  assetId;       // keccak256(serial + borrower + timestamp) — generated off-chain
        AssetType assetType;
        address  borrower;
        uint16   ltv;           // loan-to-value in basis points (e.g. 7000 = 70%)
        AssetStatus status;
        uint256  registeredAt;  // block.timestamp at registration
    }

    // ── Events ────────────────────────────────────────────────────────────────

    /// @notice Emitted when a new asset is registered.
    event AssetRegistered(
        bytes32 indexed assetId,
        AssetType       assetType,
        address indexed borrower,
        uint16          ltv
    );

    /// @notice Emitted when an asset's LTV is updated.
    event LTVUpdated(bytes32 indexed assetId, uint16 oldLTV, uint16 newLTV);

    /// @notice Emitted when an asset's status changes.
    event StatusChanged(
        bytes32 indexed assetId,
        AssetStatus     oldStatus,
        AssetStatus     newStatus
    );

    // ── Errors ────────────────────────────────────────────────────────────────

    error AssetNotFound(bytes32 assetId);
    error AssetAlreadyRegistered(bytes32 assetId);

    /// @param ltv The invalid LTV value that was rejected.
    error InvalidLTV(uint16 ltv);

    error ZeroAddress();

    // ── Write ─────────────────────────────────────────────────────────────────

    /// @notice Register a new asset. Callable only by REGISTRAR_ROLE.
    /// @param assetId   Unique identifier (keccak256 of serial + borrower + timestamp).
    /// @param assetType Class of physical asset.
    /// @param borrower  Wallet address of the borrower who owns the asset.
    /// @param ltv       Loan-to-value ratio in basis points (1–9000).
    function registerAsset(
        bytes32   assetId,
        AssetType assetType,
        address   borrower,
        uint16    ltv
    ) external;

    /// @notice Update an asset's LTV. Callable by ADMIN_ROLE.
    function updateLTV(bytes32 assetId, uint16 newLTV) external;

    /// @notice Update an asset's lifecycle status. Callable by LENDINGPOOL_ROLE or ADMIN_ROLE.
    function updateStatus(bytes32 assetId, AssetStatus newStatus) external;

    // ── View ──────────────────────────────────────────────────────────────────

    /// @notice Returns the full asset record. Reverts if not registered.
    function getAsset(bytes32 assetId) external view returns (Asset memory);

    /// @notice Returns true iff the asset is registered and has ACTIVE status.
    function isActive(bytes32 assetId) external view returns (bool);
}
