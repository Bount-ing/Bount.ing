package models

import (
	"github.com/bount-ing/bount.ing/api/db"
	"gorm.io/gorm"
)

// PRVerification represents the verification details of a pull request
type ClaimCheck struct {
	gorm.Model
	Found            bool   `json:"found"`
	Linked           bool   `json:"linked"`
	Accepted         bool   `json:"accepted"`
	Closed           bool   `json:"closed"`
	AuthorUsername   string `json:"authorUsername"`
	AuthorExternalID uint   `json:"authorExternalId"`
	RepoOwner        string `json:"repoOwner"`
	RepoName         string `json:"repoName"`
	PRNumber         string `json:"prNumber"`

	// Check Attribution
	CheckerID   uint   `json:"checkerId"`    // ID of the user who made the check (0 for system checks)
	CheckerType string `json:"checker_type"` // "SYSTEM", "OWNER", or "CLAIMER"

	// Relationship to parent Claim
}

func init() {
	db.DB.AutoMigrate(&ClaimCheck{})
}
