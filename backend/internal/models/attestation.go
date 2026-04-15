package models

import "time"

// Attestation records an on-chain health attestation for an asset.
type Attestation struct {
	ID          string    `gorm:"type:char(36);primaryKey"`
	AssetID     string    `gorm:"type:varchar(66);not null;index"` // bytes32 hex
	HealthScore uint8     `gorm:"type:tinyint unsigned;not null"`
	HealthHash  string    `gorm:"type:varchar(66);not null"` // keccak256 of telemetry snapshot
	XDCTxHash   string    `gorm:"type:varchar(66)"`          // nil until confirmed on-chain
	AttestedAt  time.Time `gorm:"not null"`
}

func (Attestation) TableName() string { return "attestations" }
