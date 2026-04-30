package models

import "time"

// Telemetry stores raw Veriflow agent readings.
// The table is RANGE-partitioned by recorded_at in MySQL; the composite
// PRIMARY KEY (id, recorded_at) satisfies MySQL's partition key requirement.
type Telemetry struct {
	ID              uint64    `gorm:"primaryKey;autoIncrement"`
	AssetID         string    `gorm:"type:char(36);not null;index:idx_asset_recorded,priority:1"`
	GPUUtilization  float64   `gorm:"type:decimal(5,2);not null;default:0"`
	GPUTemperature  float64   `gorm:"type:decimal(5,2);not null;default:0"`
	GPUMemoryUsedMB int64     `gorm:"type:bigint;not null;default:0"`
	GPUErrorRate    float64   `gorm:"type:decimal(10,6);not null;default:0"`
	PowerDrawWatts  float64   `gorm:"type:decimal(8,2);not null;default:0"`
	FanSpeedRPM     int       `gorm:"type:int;not null;default:0"`
	ECCErrors       int64     `gorm:"type:bigint;not null;default:0"`
	UptimePct       float64   `gorm:"type:decimal(5,2);not null;default:100"`
	HMACSignature   string    `gorm:"type:varchar(64);not null;default:''"`
	RawJSON         *string   `gorm:"type:json"`
	RecordedAt      time.Time `gorm:"not null;index:idx_asset_recorded,priority:2"`
	CreatedAt       time.Time `gorm:"not null"`
}

func (Telemetry) TableName() string {
	return "telemetry"
}
