package models

import (
	"gorm.io/gorm"
)

// Define BountyType as a string type
type Claim struct {
	gorm.Model
	//IssueID  uint `json:"issue_id" gorm:"not null"`
	BountyID uint `json:"bounty_id" gorm:"not null"`
}
