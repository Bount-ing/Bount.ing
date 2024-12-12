package models

import (
	"gorm.io/gorm"

	"github.com/bount-ing/bount.ing/api/db"
)

type Issue struct {
	gorm.Model   // This embeds fields like ID, CreatedAt, UpdatedAt, DeletedAt
	GithubID     int
	URL          string `json:"issueUrl"`
	AvatarURL    string `json:"avatarUrl"`
	Title        string
	Description  string
	Status       string `json:"status"`
	RepositoryID *uint  `json:"repositoryId,omitempty"` // Pointer type to make it nullable
	ClosedAt     string
	Bounties     []Bounty `gorm:"foreignKey:IssueID"`
	Claims       []Claim  `gorm:"foreignKey:IssueID"`
}

func init() {
	db.DB.AutoMigrate(&Issue{})
}
