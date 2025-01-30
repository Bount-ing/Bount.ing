package models

import (
	"errors"
	"time"

	"github.com/bount-ing/bount.ing/api/db"
	"github.com/gin-gonic/gin"
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

type BountyVariable struct {
	gorm.Model
	Amount    int       `json:"amount" binding:"required"`
	StartAt   time.Time `json:"startAt" binding:"required"`
	EndAt     time.Time `json:"endAt" binding:"required"`
	Direction string    `json:"direction" binding:"required"`
	BountyID  uint      `json:"bounty_id"`
}

type Bounty struct {
	gorm.Model
	Amount          float64          `json:"amount" binding:"required"`
	ClaimedAmount   float64          `json:"claimedAmount"`
	Currency        string           `json:"currency" binding:"required"`
	IssueURL        string           `json:"issueUrl"`
	IssueImageURL   string           `json:"issue_image_url"`
	StartAt         time.Time        `json:"startAt" binding:"required"`
	EndAt           time.Time        `json:"endAt" binding:"required"`
	OwnerID         uint             `json:"owner_id"`
	FinalizedAt     time.Time        `json:"finalized_at"`
	IssueID         uint             `json:"issue_id"`
	StripeInvoiceID string           `json:"stripe_invoice_id"`
	Claims          []Claim          `gorm:"foreignKey:BountyID" json:"claims,omitempty"`
	Variables       []BountyVariable `gorm:"foreignKey:BountyID" json:"variables,omitempty"`
	Status          string           `json:"status"`

	// Claimer Status Timestamps
	ClaimerPRFound     time.Time `json:"claimerPRFound"`     // Timestamp when PR is found by the claimer
	ClaimerPRLinked    time.Time `json:"claimerPRLinked"`    // Timestamp when PR is linked to the issue by the claimer
	ClaimerPRAccepted  time.Time `json:"claimerPRAccepted"`  // Timestamp when the PR is accepted by the owner
	ClaimerIssueClosed time.Time `json:"claimerIssueClosed"` // Timestamp when the issue is closed by the claimer

	// Owner Status Timestamps
	OwnerPRFound     time.Time `json:"ownerPRFound"`     // Timestamp when PR is found by the owner
	OwnerPRLinked    time.Time `json:"ownerPRLinked"`    // Timestamp when PR is linked to the issue by the owner
	OwnerPRAccepted  time.Time `json:"ownerPRAccepted"`  // Timestamp when the PR is accepted by the owner
	OwnerIssueClosed time.Time `json:"ownerIssueClosed"` // Timestamp when the issue is closed by the owner
}

func ValidateBountyType(bt string) error {
	switch bt {
	case Crescendo, Flat, Decrescendo:
		return nil
	default:
		return errors.New("invalid bounty type")
	}
}
func (b *Bounty) Bind(c *gin.Context) error {
	if err := c.ShouldBindJSON(b); err != nil {
		return err
	}
	return nil
}

func init() {
	db.DB.AutoMigrate(&Bounty{})
	db.DB.AutoMigrate(&BountyVariable{})
}
