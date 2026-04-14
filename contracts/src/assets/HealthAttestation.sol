// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

import {IHealthAttestation} from "../interfaces/IHealthAttestation.sol";

/// @title HealthAttestation
/// @notice Immutable on-chain log of hardware health attestations produced by the
///         Veriflow AttestationWriter (Phase 11). Each attestation anchors a
///         keccak256 hash of the 24-hour telemetry snapshot for an asset.
///
///         LendingPool (Phase 4) reads the latest attestation before originating
///         a loan — it rejects if the score is below 60 or the data is stale (> 48h).
///
///         A 12-hour minimum interval between successive attestations prevents
///         spamming and aligns with the daily Veriflow cron schedule.
///
/// @dev UUPS upgradeable. History arrays are append-only. Storage layout — only append.
contract HealthAttestation is
    Initializable,
    AccessControlUpgradeable,
    UUPSUpgradeable,
    IHealthAttestation
{
    // ── Roles ─────────────────────────────────────────────────────────────────

    /// @notice Veriflow AttestationWriter backend wallet — the only address allowed
    ///         to submit attestations.
    bytes32 public constant VERIFLOW_SIGNER = keccak256("VERIFLOW_SIGNER");

    /// @notice Protocol admin.
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    /// @notice Allows authorizing contract upgrades.
    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");

    // ── Constants ─────────────────────────────────────────────────────────────

    /// @notice Minimum seconds between two successive attestations for the same asset.
    uint256 public constant COOLDOWN = 12 hours;

    // ── Storage ───────────────────────────────────────────────────────────────

    /// @dev assetId → most recent attestation
    mapping(bytes32 => Attestation) private _latest;

    /// @dev assetId → full history (append-only)
    mapping(bytes32 => Attestation[]) private _history;

    /// @dev assetId → whether any attestation has been submitted
    mapping(bytes32 => bool) private _hasAttestation;

    // ── Constructor ───────────────────────────────────────────────────────────

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    // ── Initializer ───────────────────────────────────────────────────────────

    /// @param admin Address granted all admin roles.
    function initialize(address admin) external initializer {
        if (admin == address(0)) revert ZeroAddress();
        __AccessControl_init();
        _grantRole(DEFAULT_ADMIN_ROLE, admin);
        _grantRole(ADMIN_ROLE, admin);
        _grantRole(UPGRADER_ROLE, admin);
    }

    // ── IHealthAttestation — write ────────────────────────────────────────────

    /// @inheritdoc IHealthAttestation
    function submitAttestation(
        bytes32 assetId,
        bytes32 healthHash,
        uint8   score
    ) external onlyRole(VERIFLOW_SIGNER) {
        if (score > 100) revert InvalidScore(score);

        if (_hasAttestation[assetId]) {
            uint256 nextAllowed = _latest[assetId].timestamp + COOLDOWN;
            if (block.timestamp < nextAllowed) {
                revert AttestationTooSoon(assetId, nextAllowed);
            }
        }

        Attestation memory attestation = Attestation({
            assetId:    assetId,
            healthHash: healthHash,
            score:      score,
            timestamp:  block.timestamp
        });

        _latest[assetId] = attestation;
        _history[assetId].push(attestation);
        _hasAttestation[assetId] = true;

        emit AttestationSubmitted(assetId, healthHash, score, block.timestamp);
    }

    // ── IHealthAttestation — view ─────────────────────────────────────────────

    /// @inheritdoc IHealthAttestation
    function getLatestAttestation(bytes32 assetId) external view returns (Attestation memory) {
        return _latest[assetId];
    }

    /// @inheritdoc IHealthAttestation
    /// @dev Returns newest-first. Pass limit=0 to return all entries.
    function getAttestationHistory(bytes32 assetId, uint256 limit)
        external
        view
        returns (Attestation[] memory)
    {
        Attestation[] storage history = _history[assetId];
        uint256 len = history.length;
        if (len == 0) return new Attestation[](0);

        uint256 count = (limit == 0 || limit > len) ? len : limit;
        Attestation[] memory result = new Attestation[](count);

        // Newest first
        for (uint256 i = 0; i < count; i++) {
            result[i] = history[len - 1 - i];
        }
        return result;
    }

    /// @inheritdoc IHealthAttestation
    function hasAttestation(bytes32 assetId) external view returns (bool) {
        return _hasAttestation[assetId];
    }

    // ── Internal ──────────────────────────────────────────────────────────────

    /// @dev Only UPGRADER_ROLE can authorize an upgrade.
    function _authorizeUpgrade(address newImplementation)
        internal
        override
        onlyRole(UPGRADER_ROLE)
    {}
}
