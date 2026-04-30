-- 000010_add_ecc_uptime_to_telemetry.up.sql
-- uptime_pct was added to the base table in 006; only ecc_errors is new here.
ALTER TABLE telemetry
    ADD COLUMN ecc_errors BIGINT NOT NULL DEFAULT 0 AFTER fan_speed_rpm;
