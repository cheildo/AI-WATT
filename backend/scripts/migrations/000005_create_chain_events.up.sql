-- 000005_create_chain_events.up.sql
CREATE TABLE IF NOT EXISTS chain_events (
    id               CHAR(36)     NOT NULL,
    event_type       VARCHAR(100) NOT NULL,
    contract_address VARCHAR(42)  NOT NULL,
    tx_hash          VARCHAR(66)  NOT NULL,
    block_number     BIGINT UNSIGNED NOT NULL,
    log_index        INT UNSIGNED NOT NULL,
    args_json        JSON         NULL,
    created_at       DATETIME(3)  NOT NULL DEFAULT CURRENT_TIMESTAMP(3),

    PRIMARY KEY (id),
    UNIQUE KEY uq_chain_events_tx_log (tx_hash, log_index),
    INDEX idx_chain_events_event_type (event_type),
    INDEX idx_chain_events_contract (contract_address),
    INDEX idx_chain_events_block_number (block_number)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
