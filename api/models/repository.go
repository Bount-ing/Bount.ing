package models

import (
	"gorm.io/gorm"
)

type Repository struct {
	gorm.Model
	GithubID             int     `json:"github_id" gorm:"not null;index"`
	GithubURL            string  `json:"github_url" gorm:"not null"`
	GithubWebhookEnabled bool    `json:"github_webhook_enabled" gorm:"not null"`
	Name                 string  `json:"name" gorm:"not null"`
	Issues               []Issue `json:"issues" gorm:"foreignKey:RepositoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
