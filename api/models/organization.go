package models

import (
	"gorm.io/gorm"

	"github.com/bount-ing/bount.ing/api/db"
)

type Organization struct {
	gorm.Model
	Name         string
	Members      []User
	Repositories []Repository
	Hosts        []Host
}

func init() {
	db.DB.AutoMigrate(&Organization{})
}
