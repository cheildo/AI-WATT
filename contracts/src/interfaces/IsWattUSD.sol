// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title IsWattUSD
/// @notice Interface for the sWattUSD tokenized yield vault (ERC-4626).
/// @dev Underlying asset is WATT (WattUSD). NAV per share rises as
///      GPU loan repayments are distributed into the vault by LendingPool.
interface IsWattUSD {
    // ── Events ────────────────────────────────────────────────────────────────

    /// @notice Emitted when yield is pushed into the vault by LendingPool / TreasuryService.
    event YieldReceived(uint256 amount, uint256 newTotalAssets);

    /// @notice Emitted when the WEV queue contract address is updated.
    event WEVQueueUpdated(address indexed oldQueue, address indexed newQueue);

    /// @notice Emitted when the large-redemption threshold is updated.
    event WEVThresholdUpdated(uint256 oldThreshold, uint256 newThreshold);

    // ── Errors ────────────────────────────────────────────────────────────────

    error ZeroAmount();
    error ZeroAddress();

    /// @notice Thrown when a redemption exceeds the threshold and must use the WEV queue.
    /// @param amount    Requested redemption amount.
    /// @param threshold Current WEV threshold.
    /// @param wevQueue  Address of the WEVQueue contract to use instead.
    error LargeRedemptionUseWEVQueue(uint256 amount, uint256 threshold, address wevQueue);

    // ── Yield distribution ────────────────────────────────────────────────────

    /// @notice Transfers `amount` WATT from caller into the vault as yield.
    ///         Callable only by YIELD_DISTRIBUTOR_ROLE (LendingPool, TreasuryService).
    ///         Increases totalAssets → NAV per share rises for all sWATT holders.
    /// @param amount Amount of WATT to inject (6 decimals).
    function receiveYield(uint256 amount) external;

    // ── Admin ─────────────────────────────────────────────────────────────────

    /// @notice Sets the WEVQueue contract address. Pass address(0) to disable the guard.
    function setWEVQueue(address newWEVQueue) external;

    /// @notice Sets the redemption threshold above which users must use WEVQueue.
    function setWEVThreshold(uint256 newThreshold) external;

    /// @notice Pauses deposits, mints, withdrawals, and redemptions.
    function pause() external;

    /// @notice Unpauses the vault.
    function unpause() external;

    // ── View ──────────────────────────────────────────────────────────────────

    /// @notice Address of the WEVQueue contract (address(0) if not yet set).
    function wevQueue() external view returns (address);

    /// @notice Redemption size above which the WEVQueue must be used.
    function wevThreshold() external view returns (uint256);

    /// @notice Returns the current NAV: how many WATT one full sWATT share is worth.
    ///         Equivalent to convertToAssets(10 ** decimals()).
    function navPerShare() external view returns (uint256);
}
