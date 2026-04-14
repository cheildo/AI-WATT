// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title IHealthAttestation
/// @notice Stores daily cryptographic proofs of hardware health produced by the
///         Veriflow agent. Each attestation is a keccak256 hash of the 24-hour
///         telemetry snapshot and a normalised health score (0–100).
///
///         LendingPool reads the latest attestation before originating loans
///         to ensure the collateral asset is healthy and the data is fresh.
interface IHealthAttestation {
    // ── Structs ───────────────────────────────────────────────────────────────

    /// @notice A single on-chain health attestation record.
    struct Attestation {
        bytes32 assetId;
        bytes32 healthHash;  // keccak256(abi.encodePacked(assetId, score, avgGpuUtil, avgTemp, timestamp))
        uint8   score;       // 0–100 normalised health score
        uint256 timestamp;   // block.timestamp at submission
    }

    // ── Events ────────────────────────────────────────────────────────────────

    /// @notice Emitted on every successful attestation submission.
    event AttestationSubmitted(
        bytes32 indexed assetId,
        bytes32         healthHash,
        uint8           score,
        uint256         timestamp
    );

    // ── Errors ────────────────────────────────────────────────────────────────

    error ZeroAddress();

    /// @notice Thrown when score > 100.
    error InvalidScore(uint8 score);

    /// @notice Thrown when a second attestation is submitted within the cooldown window.
    /// @param nextAllowed Earliest timestamp at which the next attestation is accepted.
    error AttestationTooSoon(bytes32 assetId, uint256 nextAllowed);

    // ── Write ─────────────────────────────────────────────────────────────────

    /// @notice Submit a health attestation for an asset. Callable only by VERIFLOW_SIGNER.
    ///         Enforces a 12-hour minimum interval between successive attestations.
    /// @param assetId    AssetRegistry identifier.
    /// @param healthHash keccak256 hash of the 24h telemetry snapshot.
    /// @param score      Normalised health score 0–100.
    function submitAttestation(bytes32 assetId, bytes32 healthHash, uint8 score) external;

    // ── View ──────────────────────────────────────────────────────────────────

    /// @notice Returns the most recent attestation for an asset.
    ///         Returns a zero-value struct if no attestation has been submitted yet.
    function getLatestAttestation(bytes32 assetId) external view returns (Attestation memory);

    /// @notice Returns the attestation history for an asset, newest first.
    /// @param limit Maximum number of records to return. Pass 0 to return all.
    function getAttestationHistory(bytes32 assetId, uint256 limit)
        external
        view
        returns (Attestation[] memory);

    /// @notice Returns true if at least one attestation exists for the assetId.
    function hasAttestation(bytes32 assetId) external view returns (bool);
}
