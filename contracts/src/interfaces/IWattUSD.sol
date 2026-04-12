// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title IWattUSD
/// @notice Interface for the WattUSD synthetic dollar token.
/// @dev ERC-20 pegged 1:1 to USD. Mint/burn callable only by MintEngine (MINTER_ROLE).
interface IWattUSD {
    // ── Events ────────────────────────────────────────────────────────────────

    /// @notice Emitted when tokens are minted.
    event WattMinted(address indexed to, uint256 amount);

    /// @notice Emitted when tokens are burned.
    event WattBurned(address indexed from, uint256 amount);

    // ── Errors ────────────────────────────────────────────────────────────────

    /// @notice Thrown when caller does not hold MINTER_ROLE.
    error NotMinter();

    /// @notice Thrown when a zero address is passed where one is not allowed.
    error ZeroAddress();

    /// @notice Thrown when a zero amount is passed where one is not allowed.
    error ZeroAmount();

    // ── Mutative ─────────────────────────────────────────────────────────────

    /// @notice Mints `amount` WATT to `to`. Callable only by MINTER_ROLE.
    /// @param to    Recipient address.
    /// @param amount Amount in token units (6 decimals).
    function mint(address to, uint256 amount) external;

    /// @notice Burns `amount` WATT from `from`. Callable only by MINTER_ROLE.
    /// @param from  Address to burn from.
    /// @param amount Amount in token units (6 decimals).
    function burn(address from, uint256 amount) external;

    /// @notice Pauses all token transfers. Callable only by PAUSER_ROLE.
    function pause() external;

    /// @notice Unpauses token transfers. Callable only by PAUSER_ROLE.
    function unpause() external;

    // ── View ──────────────────────────────────────────────────────────────────

    /// @notice Returns the number of decimals (6 — matches USDC/USDT).
    function decimals() external view returns (uint8);
}
