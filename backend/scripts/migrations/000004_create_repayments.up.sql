-- 000004_create_repayments.up.sql
CREATE TABLE IF NOT EXISTS repayments (
    id       CHAR(36)        NOT NULL,
    loan_id  VARCHAR(66)     NOT NULL COMMENT 'bytes32 loan_id from LendingPool',
    amount   BIGINT UNSIGNED NOT NULL COMMENT 'WATT repaid, 6 decimals',
    tx_hash  VARCHAR(66)     NULL,
    paid_at  DATETIME(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3),

    PRIMARY KEY (id),
    INDEX idx_repayments_loan_id (loan_id),
    INDEX idx_repayments_paid_at (paid_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
