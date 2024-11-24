package models

import (
	"gorm.io/gorm"

	"github.com/bount-ing/bount.ing/api/db"
)

type Repository struct {
	gorm.Model
	GithubID             int
	GithubURL            string
	GithubWebhookEnabled bool
	Name                 string
	Issues               []Issue
}

func init() {
	db.DB.AutoMigrate(&Repository{})
}
