-- 000006_create_telemetry.up.sql
CREATE TABLE IF NOT EXISTS telemetry (
    id                 BIGINT UNSIGNED  NOT NULL AUTO_INCREMENT,
    asset_id           CHAR(36)         NOT NULL,
    gpu_utilization    DECIMAL(5,2)     NOT NULL DEFAULT 0.00    COMMENT 'percent 0-100',
    gpu_temperature    DECIMAL(5,2)     NOT NULL DEFAULT 0.00    COMMENT 'celsius',
    gpu_memory_used_mb BIGINT           NOT NULL DEFAULT 0       COMMENT 'megabytes',
    gpu_error_rate     DECIMAL(10,6)    NOT NULL DEFAULT 0.000000 COMMENT 'fraction e.g. 0.0001',
    power_draw_watts   DECIMAL(8,2)     NOT NULL DEFAULT 0.00    COMMENT 'watts',
    fan_speed_rpm      INT              NOT NULL DEFAULT 0,
    uptime_pct         DECIMAL(5,2)     NOT NULL DEFAULT 100.00  COMMENT 'percent 0-100',
    hmac_signature     VARCHAR(64)      NOT NULL DEFAULT '',
    raw_json           JSON             NULL,
    recorded_at        DATETIME(3)      NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    created_at         DATETIME(3)      NOT NULL DEFAULT CURRENT_TIMESTAMP(3),

    -- Composite PK required by MySQL RANGE partitioning: all partition columns
    -- must be part of every unique/primary key.
    PRIMARY KEY (id, recorded_at),
    INDEX idx_telemetry_asset_recorded (asset_id, recorded_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
PARTITION BY RANGE (TO_DAYS(recorded_at)) (
    PARTITION p_2025    VALUES LESS THAN (TO_DAYS('2026-01-01')),
    PARTITION p_2026_q1 VALUES LESS THAN (TO_DAYS('2026-04-01')),
    PARTITION p_2026_q2 VALUES LESS THAN (TO_DAYS('2026-07-01')),
    PARTITION p_2026_q3 VALUES LESS THAN (TO_DAYS('2026-10-01')),
    PARTITION p_2026_q4 VALUES LESS THAN (TO_DAYS('2027-01-01')),
    PARTITION p_future  VALUES LESS THAN MAXVALUE
);
