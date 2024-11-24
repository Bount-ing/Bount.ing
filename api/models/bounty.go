package models

import (
	"errors"
	"time"

	"github.com/bount-ing/bount.ing/api/db"
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
	Amount          float64
	BountyType      string
	Currency        string
	IssueGithubID   int
	IssueGithubURL  string
	IssueImageURL   string
	StartAt         time.Time
	EndAt           time.Time
	OwnerID         uint
	FinalizedAt     time.Time
	IssueID         uint
	StripeInvoiceID string
	Claims          []Claim
	Status          string
}

func ValidateBountyType(bt string) error {
	switch bt {
	case Crescendo, Flat, Decrescendo:
		return nil
	default:
		return errors.New("invalid bounty type")
	}
}

func init() {
	db.DB.AutoMigrate(&Bounty{})
}
