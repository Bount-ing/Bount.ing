package models

import (
	"gorm.io/gorm"
)

type Issue struct {
	gorm.Model          // This embeds fields like ID, CreatedAt, UpdatedAt, DeletedAt
	Number       int    `json:"number" gorm:"not null"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Status       string `json:"status"`
	RepositoryID uint   `json:"repository_id"`
	//Bounties     []Bounty `json:"bounties" gorm:"foreignKey:IssueID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	//Claims []Claim `json:"claims" gorm:"foreignKey:IssueID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
