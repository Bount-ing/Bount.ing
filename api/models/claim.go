package models

import (
	"github.com/bount-ing/bount.ing/api/db"
	"gorm.io/gorm"
)

type Claim struct {
	gorm.Model
	OwnerID        uint
	IssueID        uint
	BountyID       uint
	PullRequestURL string
	Status         string
}

func init() {
	db.DB.AutoMigrate(&Claim{})
}
