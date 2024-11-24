package models

import (
	"github.com/bount-ing/bount.ing/api/db"
	"gorm.io/gorm"
)

type HostRepository struct {
	gorm.Model
	HostID       uint
	RepositoryID uint
}

func init() {
	db.DB.AutoMigrate(&HostRepository{})
}
