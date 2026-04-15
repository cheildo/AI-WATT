-- 000003_create_loans.up.sql
CREATE TABLE IF NOT EXISTS loans (
    id            CHAR(36)          NOT NULL,
    loan_id       VARCHAR(66)       NOT NULL COMMENT 'bytes32 hex from LendingPool',
    asset_id      VARCHAR(66)       NOT NULL,
    borrower_id   CHAR(36)          NOT NULL,
    curator_id    CHAR(36)          NULL,
    engine_type   TINYINT UNSIGNED  NOT NULL COMMENT '1=PO 2=productivity 3=treasury',
    principal     BIGINT UNSIGNED   NOT NULL COMMENT 'WATT amount, 6 decimals',
    outstanding   BIGINT UNSIGNED   NOT NULL COMMENT 'remaining WATT owed',
    interest_rate SMALLINT UNSIGNED NOT NULL COMMENT 'annual rate in basis points',
    status        VARCHAR(20)       NOT NULL DEFAULT 'pending',
    originated_at DATETIME(3)       NULL,
    maturity_at   DATETIME(3)       NULL,
    settled_at    DATETIME(3)       NULL,
    created_at    DATETIME(3)       NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_at    DATETIME(3)       NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),

    PRIMARY KEY (id),
    UNIQUE KEY uq_loans_loan_id (loan_id),
    INDEX idx_loans_asset_id (asset_id),
    INDEX idx_loans_borrower_id (borrower_id),
    INDEX idx_loans_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
