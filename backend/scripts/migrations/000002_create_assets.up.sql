-- 000002_create_assets.up.sql
CREATE TABLE IF NOT EXISTS assets (
    id              CHAR(36)        NOT NULL,
    asset_type      VARCHAR(50)     NOT NULL COMMENT 'gpu_cluster | robotics | energy',
    owner_id        CHAR(36)        NOT NULL,
    loan_id         CHAR(36)        NULL,
    health_score    DECIMAL(5,2)    NOT NULL DEFAULT 0.00,
    current_ltv     DECIMAL(5,4)    NOT NULL DEFAULT 0.0000 COMMENT 'e.g. 0.7000 = 70%',
    status          VARCHAR(20)     NOT NULL DEFAULT 'pending' COMMENT 'pending | active | flagged | liquidated',
    ocnft_token_id  VARCHAR(78)     NOT NULL DEFAULT '',
    metadata_uri    VARCHAR(512)    NOT NULL DEFAULT '',
    hmac_secret     VARCHAR(64)     NOT NULL DEFAULT '' COMMENT 'shared HMAC secret for Veriflow agent',
    created_at      DATETIME(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_at      DATETIME(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    deleted_at      DATETIME(3)     NULL,

    PRIMARY KEY (id),
    INDEX idx_assets_owner_id (owner_id),
    INDEX idx_assets_loan_id (loan_id),
    INDEX idx_assets_status (status),
    INDEX idx_assets_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
