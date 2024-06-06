package models

import (
	"gorm.io/gorm"
)

type Issue struct {
	gorm.Model         // This embeds fields like ID, CreatedAt, UpdatedAt, DeletedAt
	GithubID    int    `json:"github_id" gorm:"not null;index"`
	GithubURL   string `json:"github_url" gorm:"not null"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	//RepositoryID uint     `json:"repository_id"`
	Bounties []Bounty `json:"bounties" gorm:"foreignKey:IssueID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Claims   []Claim  `json:"claims" gorm:"foreignKey:IssueID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
