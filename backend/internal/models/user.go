package models

import (
	"time"

	"gorm.io/gorm"
)

// User roles
const (
	RoleDepositor = "depositor"
	RoleBorrower  = "borrower"
	RoleCurator   = "curator"
	RoleAdmin     = "admin"
)

// KYC statuses
const (
	KYCPending  = "pending"
	KYCApproved = "approved"
	KYCRejected = "rejected"
)

// User is the GORM model for the users table.
type User struct {
	ID            string         `gorm:"type:char(36);primaryKey"`
	WalletAddress string         `gorm:"type:varchar(42);uniqueIndex;not null"`
	Email         string         `gorm:"type:varchar(255);uniqueIndex"`
	PasswordHash  string         `gorm:"type:varchar(255)"`
	Role          string         `gorm:"type:varchar(20);not null;default:'depositor'"`
	KYCStatus     string         `gorm:"type:varchar(20);not null;default:'pending'"`
	KYCProviderID string         `gorm:"type:varchar(255)"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

// TableName overrides the GORM default table name.
func (User) TableName() string {
	return "users"
}
