package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `json:"username" gorm:"unique;not null"`
	Email        string `json:"email" gorm:"unique;not null"`
	Password     string `json:"password" gorm:"not null"` // Password is not returned in JSON responses
	PasswordHash string `json:"-" gorm:"not null"`
	GithubID     int    `json:"github_id"`
}
