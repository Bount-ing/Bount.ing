package models

import (
	"github.com/bount-ing/bount.ing/api/db"
	"gorm.io/gorm"
)

type Identity struct {
	gorm.Model
	ID uint

	Username       string
	Email          string
	UserExternalID string

	// Relationships
	UserID uint `gorm:"not null"` // Foreign key
	User   User `gorm:"references:ID"`
	HostID uint `gorm:"not null"` // Foreign key
	Host   Host `gorm:"foreignKey:HostID"`
}

func init() {
	db.DB.AutoMigrate(&Identity{})
}
