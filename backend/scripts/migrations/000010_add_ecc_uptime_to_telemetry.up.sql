ALTER TABLE telemetry
    ADD COLUMN ecc_errors  BIGINT        NOT NULL DEFAULT 0    AFTER fan_speed_rpm,
    ADD COLUMN uptime_pct  DECIMAL(5,2)  NOT NULL DEFAULT 100  AFTER ecc_errors;
