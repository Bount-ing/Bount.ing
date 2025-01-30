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
	Repositories   []Repository `gorm:"foreignKey:HostID"`
	LogoUrl        string
	Identities     []Identity `gorm:"foreignKey:HostID"`
}

func HostsSeed() {
	// Check if the table is empty before seeding
	var count int64
	db.DB.Model(&Host{}).Count(&count)
	if count == 0 {
		// Create initial data
		host := Host{
			Model: gorm.Model{
				ID: 1,
			},
			Name:    "github",
			Address: "https://github.com",
			Port:    0, // NULL in DB means 0 in Go
			Type:    "",
			Version: "",
			LogoUrl: "https://github.githubassets.com/assets/GitHub-Mark-ea2971cee799.png",
		}

		db.DB.Create(&host)
	}
}

func init() {
	db.DB.AutoMigrate(&Host{})
	HostsSeed()
}
