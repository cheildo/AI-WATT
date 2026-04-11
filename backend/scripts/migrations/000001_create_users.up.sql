-- 000001_create_users.up.sql
-- Creates the users table for AI WATT protocol participants.

CREATE TABLE IF NOT EXISTS users (
    id              CHAR(36)        NOT NULL,
    wallet_address  VARCHAR(42)     NOT NULL,
    email           VARCHAR(255)    NULL,
    password_hash   VARCHAR(255)    NULL,
    role            VARCHAR(20)     NOT NULL DEFAULT 'depositor',
    kyc_status      VARCHAR(20)     NOT NULL DEFAULT 'pending',
    kyc_provider_id VARCHAR(255)    NULL,
    created_at      DATETIME(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_at      DATETIME(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    deleted_at      DATETIME(3)     NULL,

    PRIMARY KEY (id),
    UNIQUE KEY uq_users_wallet_address (wallet_address),
    UNIQUE KEY uq_users_email (email),
    INDEX idx_users_role (role),
    INDEX idx_users_kyc_status (kyc_status),
    INDEX idx_users_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
