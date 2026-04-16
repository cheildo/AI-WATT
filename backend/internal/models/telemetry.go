package models

import "time"

// Telemetry stores raw Veriflow agent readings.
// Partitioned by (asset_id, recorded_at date) in MySQL.
type Telemetry struct {
	ID              uint64    `gorm:"primaryKey;autoIncrement"`
	AssetID         string    `gorm:"type:char(36);not null;index:idx_asset_date"`
	GPUUtilization  float64   `gorm:"type:decimal(5,2)"`
	GPUTemperature  float64   `gorm:"type:decimal(5,2)"`
	GPUMemoryUsedMB int64     `gorm:"type:bigint"`
	GPUErrorRate    float64   `gorm:"type:decimal(10,6)"`
	PowerDrawWatts  float64   `gorm:"type:decimal(8,2)"`
	FanSpeedRPM     int       `gorm:"type:int"`
	ECCErrors       int64     `gorm:"type:bigint;default:0"`
	UptimePct       float64   `gorm:"type:decimal(5,2);default:100"`
	HMACSignature   string    `gorm:"type:varchar(64);not null"`
	RecordedAt      time.Time `gorm:"not null;index:idx_asset_date"`
	CreatedAt       time.Time
}

func (Telemetry) TableName() string {
	return "telemetry"
}
