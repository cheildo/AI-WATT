ALTER TABLE assets
    ADD COLUMN hmac_secret VARCHAR(64) NOT NULL DEFAULT '' AFTER metadata_uri;
