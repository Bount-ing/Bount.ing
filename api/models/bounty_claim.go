package models

import (
	"github.com/bount-ing/bount.ing/api/db"
)

type BountyClaim struct {
	BountyID uint
	ClaimID  uint
}

func init() {
	db.DB.AutoMigrate(&BountyClaim{})
}
