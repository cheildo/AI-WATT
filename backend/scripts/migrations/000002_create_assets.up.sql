-- 000002_create_assets.up.sql
CREATE TABLE IF NOT EXISTS assets (
    id               CHAR(36)         NOT NULL,
    asset_id         VARCHAR(66)      NOT NULL COMMENT 'bytes32 hex from AssetRegistry',
    asset_type       VARCHAR(50)      NOT NULL COMMENT 'gpu_cluster | robotics | ai_energy',
    borrower_wallet  VARCHAR(42)      NOT NULL,
    health_score     TINYINT UNSIGNED NOT NULL DEFAULT 0,
    ltv              SMALLINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'basis points, e.g. 7000 = 70%',
    status           VARCHAR(20)      NOT NULL DEFAULT 'pending',
    loan_id          CHAR(36)         NULL,
    location         VARCHAR(255)     NULL,
    hmac_secret      VARCHAR(255)     NULL COMMENT 'shared HMAC secret for Veriflow agent',
    registered_at    DATETIME(3)      NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_at       DATETIME(3)      NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),

    PRIMARY KEY (id),
    UNIQUE KEY uq_assets_asset_id (asset_id),
    INDEX idx_assets_borrower_wallet (borrower_wallet),
    INDEX idx_assets_status (status),
    INDEX idx_assets_loan_id (loan_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
