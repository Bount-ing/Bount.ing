package models

import (
	"github.com/bount-ing/bount.ing/api/db"
	"gorm.io/gorm"
)

// PRVerification represents the verification details of a pull request
type PRVerification struct {
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
}

// Claim represents the claim with two PRVerification objects, one for the owner and one for the author
type Claim struct {
	gorm.Model
	ClaimerID      uint   `json:"claimerId"`
	BountyID       uint   `json:"bountyId"`
	IssueID        uint   `json:"issueId"`
	PullRequestURL string `json:"prUrl"`
	ClaimDetails   string `json:"claimDetails"`
	Status         string `json:"status"`

	// Foreign Keys for PRVerification
	OwnerPRVerificationID  uint `json:"ownerPRVerificationID"`
	AuthorPRVerificationID uint `json:"authorPRVerificationID"`

	// Relations
	OwnerPRVerification  PRVerification `gorm:"foreignKey:OwnerPRVerificationID"`
	AuthorPRVerification PRVerification `gorm:"foreignKey:AuthorPRVerificationID"`
}

func init() {
	db.DB.AutoMigrate(&Claim{}, &PRVerification{})
}
