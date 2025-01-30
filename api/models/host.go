package models

import (
	"gorm.io/gorm"

	"github.com/bount-ing/bount.ing/api/db"
)

type Host struct {
	gorm.Model
	Name               string
	Address            string
	Port               int
	Type               string
	Version            string
	OrganizationID     uint
	Repositories       []Repository `gorm:"foreignKey:HostID"`
	LogoUrl            string
	Identities         []Identity `gorm:"foreignKey:HostID"`
	IsPaymentsProvider bool
	IsTasksProvider    bool
}

func HostsSeed() {
	// Check if the "github" host exists
	var github Host
	if db.DB.Where("name = ?", "github").First(&github).Error != nil {
		// If the host doesn't exist, create it
		host := Host{
			Name:               "github",
			Address:            "https://github.com",
			Port:               0, // NULL in DB means 0 in Go
			Type:               "",
			Version:            "",
			LogoUrl:            "https://github.githubassets.com/assets/GitHub-Mark-ea2971cee799.png",
			IsPaymentsProvider: false,
			IsTasksProvider:    true,
		}

		db.DB.Create(&host)
	}

	// Check if the "stripe" host exists
	var stripe Host
	if db.DB.Where("name = ?", "stripe").First(&stripe).Error != nil {
		// If the host doesn't exist, create it
		host := Host{
			Name:               "stripe",
			Address:            "https://stripe.com",
			Port:               0, // NULL in DB means 0 in Go
			Type:               "",
			Version:            "",
			LogoUrl:            "https://stripe.com/img/v3/home/twitter.png",
			IsPaymentsProvider: true,
			IsTasksProvider:    false,
		}

		db.DB.Create(&host)
	}
}

func init() {
	db.DB.AutoMigrate(&Host{})
	HostsSeed()
}
