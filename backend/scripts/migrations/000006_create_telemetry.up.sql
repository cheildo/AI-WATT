-- 000006_create_telemetry.up.sql
CREATE TABLE IF NOT EXISTS telemetry (
    id               CHAR(36)         NOT NULL,
    asset_id         VARCHAR(66)      NOT NULL,
    gpu_utilization  DECIMAL(5,2)     NOT NULL DEFAULT 0.00 COMMENT 'percent 0-100',
    temperature      DECIMAL(5,2)     NOT NULL DEFAULT 0.00 COMMENT 'celsius',
    error_rate       DECIMAL(10,6)    NOT NULL DEFAULT 0.000000 COMMENT 'fraction, e.g. 0.0001',
    uptime_pct       DECIMAL(6,3)     NOT NULL DEFAULT 0.000 COMMENT 'percent 0-100',
    raw_json         JSON             NULL,
    recorded_at      DATETIME(3)      NOT NULL DEFAULT CURRENT_TIMESTAMP(3),

    PRIMARY KEY (id),
    INDEX idx_telemetry_asset_recorded (asset_id, recorded_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
PARTITION BY RANGE (TO_DAYS(recorded_at)) (
    PARTITION p_2025 VALUES LESS THAN (TO_DAYS('2026-01-01')),
    PARTITION p_2026_q1 VALUES LESS THAN (TO_DAYS('2026-04-01')),
    PARTITION p_2026_q2 VALUES LESS THAN (TO_DAYS('2026-07-01')),
    PARTITION p_2026_q3 VALUES LESS THAN (TO_DAYS('2026-10-01')),
    PARTITION p_2026_q4 VALUES LESS THAN (TO_DAYS('2027-01-01')),
    PARTITION p_future VALUES LESS THAN MAXVALUE
);
