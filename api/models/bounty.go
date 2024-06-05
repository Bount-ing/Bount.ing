package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// Define BountyType as a string type
type BountyType string

// Define constants for BountyType
const (
	Crescendo   string = "crescendo"
	Flat        string = "flat"
	Decrescendo string = "decrescendo"
)

type Bounty struct {
	gorm.Model
	Amount          float64   `json:"amount" gorm:"not null"`
	BountyType      string    `json:"bounty_type" gorm:"default:'flat'"`
	Currency        string    `json:"currency" gorm:"default:'USD'"`
	IssueGithubID   int       `json:"issue_github_id" gorm:"not null"`
	IssueGithubURL  string    `json:"issue_github_url" gorm:"not null"`
	IssueImageURL   string    `json:"issue_image_url" gorm:"not null"`
	UserGithubLogin string    `json:"user_github_login" gorm:"not null"`
	CreatedAt       time.Time `json:"created_at"`
	StartAt         time.Time `json:"start_at"`
	EndAt           time.Time `json:"end_at"`
}

func ValidateBountyType(bt string) error {
	switch bt {
	case Crescendo, Flat, Decrescendo:
		return nil
	default:
		return errors.New("invalid bounty type")
	}
}
