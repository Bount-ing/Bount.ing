package models

import (
	"time"

	"gorm.io/gorm"

	"github.com/bount-ing/bount.ing/api/db"
)

type User struct {
	gorm.Model

	Email    string
	Password string
	Admin    bool

	Verified                bool
	VerifCode               string
	VerifCodeExpirationTime time.Time

	GithubID int
	Username string

	PublishedBounties []Bounty
	Claims            []Claim
	StipeAccountID    string
}

func init() {
	db.DB.AutoMigrate(&User{})
}
