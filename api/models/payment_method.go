package models

import (
	"github.com/bount-ing/bount.ing/api/db"
)

// PaymentMethod represents the saved payment method in the database
type PaymentMethod struct {
	ID       string `gorm:"primaryKey"`
	UserID   uint   `gorm:"index"`
	Brand    string
	Last4    string
	ExpMonth int
	ExpYear  int
}

func init() {
	db.DB.AutoMigrate(&PaymentMethod{})
}
