// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title IMintEngine
/// @notice Interface for the MintEngine — entry point for WATT minting and redemption.
interface IMintEngine {
    // ── Events ────────────────────────────────────────────────────────────────

    /// @notice Emitted on successful WATT mint.
    /// @param depositor      Address that deposited stablecoin.
    /// @param stablecoin     Stablecoin address used (USDC or USDT).
    /// @param depositAmount  Gross stablecoin amount deposited.
    /// @param wattMinted     Net WATT minted to depositor (after fee).
    /// @param fee            Fee amount in stablecoin units sent to treasury.
    event Minted(
        address indexed depositor,
        address indexed stablecoin,
        uint256 depositAmount,
        uint256 wattMinted,
        uint256 fee
    );

    /// @notice Emitted on successful WATT redemption.
    /// @param redeemer           Address redeeming WATT.
    /// @param stablecoin         Stablecoin returned to redeemer.
    /// @param wattBurned         Amount of WATT burned.
    /// @param stablecoinReturned Net stablecoin returned (after fee).
    /// @param fee                Fee amount in stablecoin units sent to treasury.
    event Redeemed(
        address indexed redeemer,
        address indexed stablecoin,
        uint256 wattBurned,
        uint256 stablecoinReturned,
        uint256 fee
    );

    /// @notice Emitted when a stablecoin is added or removed from the accepted list.
    event StablecoinUpdated(address indexed stablecoin, bool accepted);

    /// @notice Emitted when the treasury address is updated.
    event TreasuryUpdated(address indexed oldTreasury, address indexed newTreasury);

    // ── Errors ────────────────────────────────────────────────────────────────

    error ZeroAmount();
    error ZeroAddress();
    error StablecoinNotAccepted(address stablecoin);
    error InsufficientCollateral(address stablecoin, uint256 requested, uint256 available);

    // ── Mutative ─────────────────────────────────────────────────────────────

    /// @notice Deposits `amount` of `stablecoin`, mints WATT 1:1 minus 0.1% fee.
    /// @param stablecoin  USDC or USDT address.
    /// @param amount      Gross deposit amount (in stablecoin decimals).
    function mint(address stablecoin, uint256 amount) external;

    /// @notice Burns `wattAmount` WATT and returns the equivalent stablecoin minus 0.1% fee.
    /// @param stablecoin  Stablecoin to receive on redemption.
    /// @param wattAmount  Amount of WATT to redeem.
    function redeem(address stablecoin, uint256 wattAmount) external;

    // ── Admin ─────────────────────────────────────────────────────────────────

    /// @notice Adds or removes a stablecoin from the accepted list. Callable by ADMIN_ROLE.
    function setAcceptedStablecoin(address stablecoin, bool accepted) external;

    /// @notice Updates the treasury address. Callable by ADMIN_ROLE.
    function setTreasury(address newTreasury) external;

    // ── View ──────────────────────────────────────────────────────────────────

    /// @notice Returns the stablecoin collateral balance held in this contract.
    function collateralBalance(address stablecoin) external view returns (uint256);

    /// @notice Returns whether `stablecoin` is accepted for minting.
    function isAcceptedStablecoin(address stablecoin) external view returns (bool);

    /// @notice Mint/redeem fee in basis points (10 = 0.1%).
    function FEE_BPS() external view returns (uint256);
}
