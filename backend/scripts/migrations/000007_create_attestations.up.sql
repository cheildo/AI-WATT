-- 000007_create_attestations.up.sql
CREATE TABLE IF NOT EXISTS attestations (
    id           CHAR(36)         NOT NULL,
    asset_id     VARCHAR(66)      NOT NULL,
    health_score TINYINT UNSIGNED NOT NULL,
    health_hash  VARCHAR(66)      NOT NULL COMMENT 'keccak256 of telemetry snapshot',
    xdc_tx_hash  VARCHAR(66)      NULL     COMMENT 'on-chain tx; NULL until confirmed',
    attested_at  DATETIME(3)      NOT NULL DEFAULT CURRENT_TIMESTAMP(3),

    PRIMARY KEY (id),
    INDEX idx_attestations_asset_id (asset_id),
    INDEX idx_attestations_attested_at (attested_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
