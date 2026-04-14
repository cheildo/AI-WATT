// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title IOCNFT
/// @notice Ownership Certificate NFT. One ERC-721 token is minted per approved
///         physical asset deal. Token IDs map 1-to-1 with assetIds from AssetRegistry.
///
///         Tokens are soulbound by default — only MINTER_ROLE can transfer them
///         (e.g. on settlement to release collateral to the borrower).
interface IOCNFT {
    // ── Events ────────────────────────────────────────────────────────────────

    /// @notice Emitted when an OC-NFT is minted for an asset.
    event OCNFTMinted(
        uint256 indexed tokenId,
        address indexed to,
        bytes32 indexed assetId
    );

    /// @notice Emitted when an OC-NFT is burned (asset settled or liquidated).
    event OCNFTBurned(uint256 indexed tokenId, bytes32 indexed assetId);

    // ── Errors ────────────────────────────────────────────────────────────────

    /// @notice Thrown when attempting to mint a second OC-NFT for the same assetId.
    error AlreadyMinted(bytes32 assetId);

    /// @notice Thrown when querying a tokenId that does not exist.
    error TokenNotFound(uint256 tokenId);

    /// @notice Thrown when a non-MINTER_ROLE address attempts a transfer.
    error TransferRestricted();

    error ZeroAddress();

    // ── Write ─────────────────────────────────────────────────────────────────

    /// @notice Mint an OC-NFT for a registered asset. Callable only by MINTER_ROLE.
    /// @param to          Recipient (typically the borrower).
    /// @param assetId     AssetRegistry identifier.
    /// @param metadataURI IPFS URI of the asset metadata JSON.
    /// @return tokenId    The newly minted token ID.
    function mintOCNFT(
        address to,
        bytes32 assetId,
        string calldata metadataURI
    ) external returns (uint256 tokenId);

    /// @notice Burn an OC-NFT. Callable only by MINTER_ROLE.
    function burnOCNFT(uint256 tokenId) external;

    // ── View ──────────────────────────────────────────────────────────────────

    /// @notice Returns the assetId linked to a token. Reverts if tokenId does not exist.
    function getAssetId(uint256 tokenId) external view returns (bytes32);

    /// @notice Returns the tokenId linked to an assetId. Returns 0 if no token minted yet.
    function getTokenId(bytes32 assetId) external view returns (uint256);
}
