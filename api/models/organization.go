package models

import (
	"gorm.io/gorm"

	"github.com/bount-ing/bount.ing/api/db"
)

type Organization struct {
	gorm.Model
	Name string `json:"name" gorm:"unique;not null"`
}

func init() {
	db.DB.AutoMigrate(&Organization{})
}
