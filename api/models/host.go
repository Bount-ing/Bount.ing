package models

import (
	"gorm.io/gorm"

	"github.com/bount-ing/bount.ing/api/db"
)

type Host struct {
	gorm.Model
	Name           string
	Address        string
	Port           int
	Type           string
	Version        string
	OrganizationID uint
	Repositories   []HostRepository
}

func init() {
	db.DB.AutoMigrate(&Host{})
}
