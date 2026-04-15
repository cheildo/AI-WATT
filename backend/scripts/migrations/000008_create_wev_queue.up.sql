-- 000008_create_wev_queue.up.sql
CREATE TABLE IF NOT EXISTS wev_queue (
    id           CHAR(36)        NOT NULL,
    request_id   VARCHAR(66)     NOT NULL COMMENT 'bytes32 requestId from WEVQueue contract',
    user_id      CHAR(36)        NOT NULL,
    swatt_amount BIGINT UNSIGNED NOT NULL COMMENT 'sWATT shares, 6 decimals',
    priority_fee BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'WATT fee paid for priority',
    status       VARCHAR(20)     NOT NULL DEFAULT 'queued' COMMENT 'queued | processing | fulfilled | cancelled',
    requested_at DATETIME(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    processed_at DATETIME(3)     NULL,

    PRIMARY KEY (id),
    UNIQUE KEY uq_wev_queue_request_id (request_id),
    INDEX idx_wev_queue_user_id (user_id),
    INDEX idx_wev_queue_status (status),
    INDEX idx_wev_queue_requested_at (requested_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
