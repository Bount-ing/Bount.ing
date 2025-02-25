package models

import (
	"github.com/bount-ing/bount.ing/api/db"
	"gorm.io/gorm"
)

// Claim represents the claim with two PRVerification objects, one for the sponsor and one for the author
type Claim struct {
	gorm.Model
	ClaimerID      uint    `json:"claimerId"`
	BountyID       uint    `json:"bountyId"`
	IssueID        uint    `json:"issueId"`
	PullRequestURL string  `json:"prUrl"`
	ClaimDetails   string  `json:"claimDetails"`
	Status         string  `json:"status"`
	ClaimedAmount  float64 `json:"claimedAmount"`

	// Relationships
	Bounty             Bounty     `gorm:"foreignKey:BountyID"`
	BountyClaimerCheck ClaimCheck `gorm:"foreignKey:BountyClaimerCheckID"`
	BountySponsorCheck ClaimCheck `gorm:"foreignKey:BountySponsorCheckID"`
	BountySystemCheck  ClaimCheck `gorm:"foreignKey:BountySystemCheckID"`

	// Foreign keys for ClaimChecks
	BountyClaimerCheckID uint `json:"bountyClaimerCheckID"`
	BountySponsorCheckID uint `json:"sponsorCheckID"`
	BountySystemCheckID  uint `json:"bountySystemCheckID"`
}

func init() {
	db.DB.AutoMigrate(&Claim{})
}
