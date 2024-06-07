package models

import (
	"gorm.io/gorm"
)

type BountyClaim struct {
	BountyID uint
	ClaimID  uint
}

type Claim struct {
	gorm.Model
	OwnerID  uint   `json:"owner_id" gorm:"not null"`
	IssueID  uint   `json:"issue_id" gorm:"not null"`
	BountyID uint   `json:"bounty_id" gorm:"many2many:bounty_claims;"`
	Status   string `json:"status" gorm:"default:'pending'"`
}
