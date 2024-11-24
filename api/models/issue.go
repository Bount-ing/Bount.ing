package models

import (
	"gorm.io/gorm"

	"github.com/bount-ing/bount.ing/api/db"
)

type Issue struct {
	gorm.Model   // This embeds fields like ID, CreatedAt, UpdatedAt, DeletedAt
	GithubID     int
	GithubURL    string
	Title        string
	Description  string
	Status       string
	RepositoryID uint
	ClosedAt     string
	Bounties     []Bounty
	Claims       []Claim
}

func init() {
	db.DB.AutoMigrate(&Issue{})
}
