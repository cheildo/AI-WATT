package models

import (
	"time"

	"gorm.io/gorm"
)

// Asset types
const (
	AssetTypeGPUCluster = "gpu_cluster"
	AssetTypeRobotics   = "robotics"
	AssetTypeEnergy     = "energy"
)

// Asset statuses
const (
	AssetStatusPending    = "pending"
	AssetStatusActive     = "active"
	AssetStatusFlagged    = "flagged"
	AssetStatusLiquidated = "liquidated"
)

// Asset is the GORM model for the assets table.
type Asset struct {
	ID           string         `gorm:"type:char(36);primaryKey"`
	AssetType    string         `gorm:"type:varchar(50);not null"`
	LoanID       string         `gorm:"type:char(36);index"`
	OwnerID      string         `gorm:"type:char(36);not null;index"`
	HealthScore  float64        `gorm:"type:decimal(5,2);default:0"`
	CurrentLTV   float64        `gorm:"type:decimal(5,4);default:0"`
	Status       string         `gorm:"type:varchar(20);not null;default:'pending'"`
	OCNFTTokenID string         `gorm:"type:varchar(78)"`
	MetadataURI  string         `gorm:"type:varchar(512)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (Asset) TableName() string {
	return "assets"
}
