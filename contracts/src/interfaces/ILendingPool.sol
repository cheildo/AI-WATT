// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title ILendingPool
/// @notice On-chain loan lifecycle for Engine 2: post-delivery productivity-backed loans.
///         Reads AssetRegistry to verify active collateral and HealthAttestation to ensure
///         the asset is healthy before disbursing WATT to the borrower.
///         Repayments are split: interest → sWattUSD.receiveYield(), principal → pool.
interface ILendingPool {
    // ── Enums ─────────────────────────────────────────────────────────────────

    /// @notice Lifecycle status of a loan.
    enum LoanStatus { PENDING, ACTIVE, REPAYING, SETTLED, DEFAULTED, LIQUIDATED }

    // ── Structs ───────────────────────────────────────────────────────────────

    /// @notice On-chain loan record.
    struct Loan {
        bytes32    loanId;
        bytes32    assetId;
        address    borrower;
        address    curator;
        uint256    principal;      // WATT (6 decimals)
        uint256    outstanding;    // remaining WATT owed (principal + accrued interest)
        uint256    interestRate;   // annual rate in basis points (e.g. 1200 = 12%)
        LoanStatus status;
        uint8      engineType;     // 1 = PO financing, 2 = productivity, 3 = treasury
        uint256    originatedAt;
        uint256    maturityAt;
    }

    // ── Events ────────────────────────────────────────────────────────────────

    event LoanOriginated(
        bytes32 indexed loanId,
        bytes32 indexed assetId,
        address indexed borrower,
        address         curator,
        uint256         principal,
        uint256         interestRate,
        uint256         maturityAt
    );

    event RepaymentReceived(
        bytes32 indexed loanId,
        address indexed payer,
        uint256         amount,
        uint256         principalPaid,
        uint256         interestPaid,
        uint256         outstandingAfter
    );

    event LoanSettled(bytes32 indexed loanId, address indexed borrower);

    event LoanDefaulted(bytes32 indexed loanId, address indexed borrower);

    event LoanLiquidated(bytes32 indexed loanId, address indexed borrower, address liquidator);

    event FeesWithdrawn(address indexed to, uint256 amount);

    // ── Errors ────────────────────────────────────────────────────────────────

    error ZeroAddress();
    error ZeroAmount();

    /// @notice Asset is not registered and ACTIVE in AssetRegistry.
    error AssetNotActive(bytes32 assetId);

    /// @notice Latest attestation timestamp is older than 48 hours.
    error AttestationStale(bytes32 assetId, uint256 lastTimestamp);

    /// @notice Latest health score is below the minimum threshold (60).
    error HealthScoreTooLow(bytes32 assetId, uint8 score);

    /// @notice Loan does not exist.
    error LoanNotFound(bytes32 loanId);

    /// @notice Operation requires loan to be in ACTIVE or REPAYING status.
    error LoanNotActive(bytes32 loanId, LoanStatus status);

    /// @notice Asset already has an active loan — cannot double-encumber collateral.
    error AssetAlreadyEncumbered(bytes32 assetId);

    /// @notice Repayment amount exceeds outstanding balance.
    error ExceedsOutstanding(uint256 amount, uint256 outstanding);

    /// @notice No protocol fees available to withdraw.
    error NoFeesAvailable();

    // ── Write ─────────────────────────────────────────────────────────────────

    /// @notice Originate a new loan. Callable only by CURATOR_ROLE.
    ///         Verifies asset is ACTIVE and health score >= 60 with fresh attestation (< 48h).
    ///         Mints `principal` WATT to borrower.
    /// @param assetId      AssetRegistry identifier of the collateral.
    /// @param borrower     Wallet receiving the WATT disbursement.
    /// @param principal    WATT amount to disburse (6 decimals).
    /// @param interestRate Annual interest rate in basis points.
    /// @param termDays     Loan duration in days.
    /// @param engineType   Engine classification (1, 2, or 3).
    /// @return loanId      keccak256 identifier for the new loan.
    function originateLoan(
        bytes32 assetId,
        address borrower,
        uint256 principal,
        uint256 interestRate,
        uint256 termDays,
        uint8   engineType
    ) external returns (bytes32 loanId);

    /// @notice Make a partial or full repayment on a loan.
    ///         Interest portion is routed to sWattUSD.receiveYield().
    ///         Automatically settles the loan if outstanding reaches zero.
    /// @param loanId Loan identifier.
    /// @param amount WATT amount to repay (must not exceed outstanding).
    function repay(bytes32 loanId, uint256 amount) external;

    /// @notice Repay the full outstanding balance in one transaction.
    function fullRepay(bytes32 loanId) external;

    /// @notice Flag an overdue loan as DEFAULTED. Callable by anyone once past maturity.
    function flagDefaulted(bytes32 loanId) external;

    /// @notice Liquidate a DEFAULTED loan. Callable only by LIQUIDATOR_ROLE.
    ///         Updates AssetRegistry status to LIQUIDATED.
    function liquidate(bytes32 loanId) external;

    /// @notice Withdraw accumulated protocol fees. Callable only by ADMIN_ROLE.
    function withdrawFees(address to) external;

    // ── View ──────────────────────────────────────────────────────────────────

    /// @notice Returns the full loan record. Reverts if loanId does not exist.
    function getLoan(bytes32 loanId) external view returns (Loan memory);

    /// @notice Returns all loan IDs originated for a borrower address.
    function getBorrowerLoans(address borrower) external view returns (bytes32[] memory);

    /// @notice Returns total accumulated protocol fees not yet withdrawn (WATT, 6 decimals).
    function getProtocolFees() external view returns (uint256);

    /// @notice Minimum health score required to originate a loan.
    function MIN_HEALTH_SCORE() external view returns (uint8);

    /// @notice Maximum age of a valid attestation in seconds (48 hours).
    function ATTESTATION_MAX_AGE() external view returns (uint256);
}
