package models

import (
	"gorm.io/gorm"

	"github.com/bount-ing/bount.ing/api/db"
)

type Organization struct {
	gorm.Model
	Name         string
	Members      []User       `gorm:"many2many:organization_members;"`
	Repositories []Repository `gorm:"foreignKey:OrganizationID"`
	Hosts        []Host       `gorm:"foreignKey:OrganizationID"`
}

func init() {
	db.DB.AutoMigrate(&Organization{})
}
