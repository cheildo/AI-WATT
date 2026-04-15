// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title IWEVQueue
/// @notice sWATT redemption queue. Large redemptions that exceed the sWattUSD threshold
///         are routed here. Users enter the standard (~30-day) queue or pay a 0.5% priority
///         fee for the expedited (~3-day) queue. A backend keeper calls processBatch to
///         fulfill requests on schedule.
interface IWEVQueue {
    // ── Enums ─────────────────────────────────────────────────────────────────

    enum RequestStatus { QUEUED, PROCESSING, FULFILLED, CANCELLED }

    // ── Structs ───────────────────────────────────────────────────────────────

    struct RedemptionRequest {
        bytes32       requestId;
        address       user;
        uint256       sWattAmount;  // sWATT shares held by WEVQueue pending fulfillment
        uint256       priorityFee;  // WATT fee paid (0 for standard; >0 for priority)
        uint256       requestedAt;
        RequestStatus status;
    }

    // ── Events ────────────────────────────────────────────────────────────────

    /// @notice Emitted when a redemption request is created.
    event RedemptionRequested(
        bytes32 indexed requestId,
        address indexed user,
        uint256         sWattAmount,
        bool            isPriority
    );

    /// @notice Emitted when a request is fulfilled during processBatch.
    event RedemptionFulfilled(
        bytes32 indexed requestId,
        address indexed user,
        uint256         wattAmount
    );

    /// @notice Emitted when a user cancels their own QUEUED request.
    event RedemptionCancelled(bytes32 indexed requestId, address indexed user);

    /// @notice Emitted after processBatch completes.
    event BatchProcessed(uint256 count);

    /// @notice Emitted when the admin withdraws accumulated priority fees.
    event FeesWithdrawn(address indexed to, uint256 amount);

    // ── Errors ────────────────────────────────────────────────────────────────

    error ZeroAddress();
    error ZeroAmount();

    /// @notice requestId does not exist.
    error RequestNotFound(bytes32 requestId);

    /// @notice Request is not in QUEUED status.
    error NotQueued(bytes32 requestId, RequestStatus status);

    /// @notice Only the request owner may cancel.
    error NotRequestOwner(bytes32 requestId);

    /// @notice processBatch was called with an empty array.
    error NothingToProcess();

    /// @notice Priority fee is below the required minimum (0.5% of sWattAmount).
    error InsufficientPriorityFee(uint256 provided, uint256 required);

    /// @notice No accumulated priority fees to withdraw.
    error NoFeesAvailable();

    // ── Write ─────────────────────────────────────────────────────────────────

    /// @notice Enter the standard redemption queue (~30-day wait).
    ///         Pulls `sWattAmount` sWATT from the caller.
    /// @return requestId keccak256 identifier of the new request.
    function requestRedeem(uint256 sWattAmount) external returns (bytes32 requestId);

    /// @notice Enter the priority redemption queue (~3-day wait).
    ///         Pulls `sWattAmount` sWATT and `priorityFee` WATT from the caller.
    ///         `priorityFee` must be >= 0.5% of `sWattAmount`.
    /// @return requestId keccak256 identifier of the new request.
    function requestPriorityRedeem(uint256 sWattAmount, uint256 priorityFee)
        external returns (bytes32 requestId);

    /// @notice Cancel a QUEUED request. Returns sWATT (and priority fee WATT) to the caller.
    ///         Only the original requester may cancel.
    function cancelRequest(bytes32 requestId) external;

    /// @notice Fulfill a batch of QUEUED requests. Callable only by PROCESSOR_ROLE.
    ///         For each request: redeems sWATT via sWattUSD → sends WATT to the user,
    ///         accumulates any priority fee into protocol fees.
    function processBatch(bytes32[] calldata requestIds) external;

    /// @notice Withdraw accumulated priority fees to `to`. Callable only by ADMIN_ROLE.
    function withdrawFees(address to) external;

    // ── View ──────────────────────────────────────────────────────────────────

    /// @notice Returns the full redemption request. Reverts if not found.
    function getRequest(bytes32 requestId) external view returns (RedemptionRequest memory);

    /// @notice Returns all request IDs created by `user`.
    function getUserRequests(address user) external view returns (bytes32[] memory);

    /// @notice Returns the number of currently QUEUED (active) requests.
    function getQueueDepth() external view returns (uint256);

    /// @notice Returns the estimated next standard-queue processing timestamp.
    function nextProcessingTimestamp() external view returns (uint256);

    /// @notice Returns accumulated priority fees not yet withdrawn (WATT, 6 decimals).
    function getProtocolFees() external view returns (uint256);

    // ── Constants ─────────────────────────────────────────────────────────────

    /// @notice Priority queue fee in basis points (50 = 0.5%).
    function PRIORITY_FEE_BPS() external view returns (uint256);

    /// @notice Estimated wait for standard queue (30 days in seconds).
    function STANDARD_WAIT() external view returns (uint256);

    /// @notice Estimated wait for priority queue (3 days in seconds).
    function PRIORITY_WAIT() external view returns (uint256);
}
