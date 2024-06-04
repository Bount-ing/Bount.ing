package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string `json:"username" gorm:"unique;not null"`
	Email          string `json:"email" gorm:"unique;not null"`
	GithubID       int    `json:"github_id"`
	StipeAccountID string `json:"stripe_id"`
}
