package models

import (
	"github.com/bount-ing/bount.ing/api/db"
	"gorm.io/gorm"
)

type Claim struct {
	gorm.Model
	OwnerID  uint   `json:"owner_id" gorm:"not null"`
	IssueID  uint   `json:"issue_id" gorm:"not null"`
	BountyID uint   `json:"bounty_id" gorm:"many2many:bounty_claims;"`
	Status   string `json:"status" gorm:"default:'pending'"`
}

func init() {
	db.DB.AutoMigrate(&Claim{})
}
