-- 000003_create_loans.up.sql
CREATE TABLE IF NOT EXISTS loans (
    id               CHAR(36)        NOT NULL,
    asset_id         CHAR(36)        NOT NULL,
    borrower_id      CHAR(36)        NOT NULL,
    curator_id       CHAR(36)        NULL,
    amount           DECIMAL(20,6)   NOT NULL COMMENT 'WATT amount',
    engine_type      TINYINT         NOT NULL COMMENT '1=pre-delivery 2=post-delivery 3=capital-reactivation',
    status           VARCHAR(20)     NOT NULL DEFAULT 'pending' COMMENT 'pending | active | repaying | settled | defaulted',
    on_chain_tx_hash VARCHAR(66)     NOT NULL DEFAULT '',
    created_at       DATETIME(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_at       DATETIME(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    deleted_at       DATETIME(3)     NULL,

    PRIMARY KEY (id),
    INDEX idx_loans_asset_id (asset_id),
    INDEX idx_loans_borrower_id (borrower_id),
    INDEX idx_loans_status (status),
    INDEX idx_loans_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
