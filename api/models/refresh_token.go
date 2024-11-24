package models

import (
	"time"

	"github.com/bount-ing/bount.ing/api/db"
	"gorm.io/gorm"
)

type RefreshToken struct {
	gorm.Model
	Value      string
	ValidUntil time.Time
	UserID     uint
}

func init() {
	db.DB.AutoMigrate(&RefreshToken{})
}
