// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {ERC721Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";
import {ERC721URIStorageUpgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol";
import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

import {IOCNFT} from "../interfaces/IOCNFT.sol";

/// @title OCNFT
/// @notice Ownership Certificate NFT (ERC-721). One token is minted per approved
///         physical asset deal. Token IDs map 1-to-1 with AssetRegistry assetIds.
///
///         Tokens are soulbound by default: transfers are blocked unless the caller
///         holds MINTER_ROLE. MINTER_ROLE may transfer on settlement (releasing the
///         collateral certificate to the borrower after loan repayment).
///
///         Metadata URIs point to IPFS JSON uploaded by the backend (Pinata) at mint time.
///
/// @dev UUPS upgradeable. Token ID counter starts at 1 — 0 is reserved as "no token".
contract OCNFT is
    Initializable,
    ERC721URIStorageUpgradeable,
    AccessControlUpgradeable,
    UUPSUpgradeable,
    IOCNFT
{
    // ── Roles ─────────────────────────────────────────────────────────────────

    /// @notice Backend signer — may mint, burn, and transfer OC-NFTs.
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

    /// @notice Protocol admin.
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    /// @notice Allows authorizing contract upgrades.
    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");

    // ── Storage ───────────────────────────────────────────────────────────────

    /// @dev Auto-incrementing token ID counter. Starts at 1 after initialize().
    uint256 private _nextTokenId;

    /// @dev tokenId → assetId
    mapping(uint256 => bytes32) private _tokenToAsset;

    /// @dev assetId → tokenId (0 = no token minted)
    mapping(bytes32 => uint256) private _assetToToken;

    // ── Constructor ───────────────────────────────────────────────────────────

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    // ── Initializer ───────────────────────────────────────────────────────────

    /// @param admin Address granted all admin roles.
    function initialize(address admin) external initializer {
        if (admin == address(0)) revert ZeroAddress();
        __ERC721_init("AI WATT Ownership Certificate", "OC-NFT");
        __ERC721URIStorage_init();
        __AccessControl_init();
        _grantRole(DEFAULT_ADMIN_ROLE, admin);
        _grantRole(ADMIN_ROLE, admin);
        _grantRole(UPGRADER_ROLE, admin);
        _nextTokenId = 1;
    }

    // ── IOCNFT — write ────────────────────────────────────────────────────────

    /// @inheritdoc IOCNFT
    function mintOCNFT(
        address to,
        bytes32 assetId,
        string calldata metadataURI
    ) external onlyRole(MINTER_ROLE) returns (uint256 tokenId) {
        if (to == address(0)) revert ZeroAddress();
        if (_assetToToken[assetId] != 0) revert AlreadyMinted(assetId);

        tokenId = _nextTokenId++;
        _tokenToAsset[tokenId] = assetId;
        _assetToToken[assetId] = tokenId;

        _safeMint(to, tokenId);
        _setTokenURI(tokenId, metadataURI);

        emit OCNFTMinted(tokenId, to, assetId);
    }

    /// @inheritdoc IOCNFT
    function burnOCNFT(uint256 tokenId) external onlyRole(MINTER_ROLE) {
        bytes32 assetId = _tokenToAsset[tokenId];
        if (assetId == bytes32(0)) revert TokenNotFound(tokenId);

        delete _assetToToken[assetId];
        delete _tokenToAsset[tokenId];

        _burn(tokenId);

        emit OCNFTBurned(tokenId, assetId);
    }

    // ── IOCNFT — view ─────────────────────────────────────────────────────────

    /// @inheritdoc IOCNFT
    function getAssetId(uint256 tokenId) external view returns (bytes32) {
        bytes32 assetId = _tokenToAsset[tokenId];
        if (assetId == bytes32(0)) revert TokenNotFound(tokenId);
        return assetId;
    }

    /// @inheritdoc IOCNFT
    /// @dev Returns 0 if no OC-NFT has been minted for this assetId.
    function getTokenId(bytes32 assetId) external view returns (uint256) {
        return _assetToToken[assetId];
    }

    // ── ERC-721 overrides ─────────────────────────────────────────────────────

    /// @dev Soulbound guard: blocks transfers unless caller is MINTER_ROLE.
    ///      Minting (from == address(0)) and burning (to == address(0)) are always allowed.
    function _update(address to, uint256 tokenId, address auth)
        internal
        override(ERC721Upgradeable)
        returns (address)
    {
        address from = _ownerOf(tokenId);
        if (from != address(0) && to != address(0)) {
            // Transfer (not mint, not burn) — only MINTER_ROLE allowed
            if (!hasRole(MINTER_ROLE, _msgSender())) revert TransferRestricted();
            // Pass address(0) as auth to skip OZ's ownership/approval check —
            // MINTER_ROLE authorization is already verified above.
            return super._update(to, tokenId, address(0));
        }
        return super._update(to, tokenId, auth);
    }

    /// @dev Resolves diamond: ERC721URIStorageUpgradeable + AccessControlUpgradeable both
    ///      implement supportsInterface.
    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC721URIStorageUpgradeable, AccessControlUpgradeable)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }

    // ── Internal ──────────────────────────────────────────────────────────────

    /// @dev Only UPGRADER_ROLE can authorize an upgrade.
    function _authorizeUpgrade(address newImplementation)
        internal
        override
        onlyRole(UPGRADER_ROLE)
    {}
}
