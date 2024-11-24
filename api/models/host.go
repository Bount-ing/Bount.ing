package models

import (
	"gorm.io/gorm"

	"github.com/bount-ing/bount.ing/api/db"
)

type Host struct {
	gorm.Model
	Name         string
	HostType     string
	HostConfig   string
	Repositories []HostRepository
}

func init() {
	db.DB.AutoMigrate(&Host{})
}
