package models

import (
	"time"

	"gorm.io/gorm"

	"github.com/bount-ing/bount.ing/api/db"
)

type UserProfileUpdatePayload struct {
	Username string `json:"username"`
	FullName string `json:"fullName"`
}

type UserProfileReadPayload struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Admin    bool   `json:"admin"`

	StripeConnected bool `json:"stripe_account_connected"`
	GithubConnected bool `json:"github_connected"`
}

type User struct {
	gorm.Model

	// System
	Email                   string    `gorm:"uniqueIndex" json:"email"`
	Password                string    `json:"-"` // "-" to never send password in JSON
	Admin                   bool      `json:"admin"`
	Verified                bool      `json:"verified"`
	VerifCode               string    `json:"-"`
	VerifCodeExpirationTime time.Time `json:"-"`
	LastLogin               time.Time

	// Profile Settings
	Username string        `json:"username"`
	FullName string        `json:"fullName"`
	Legal    UserLegalData `gorm:"foreignKey:ID" json:"legal,omitempty"`

	// Relationships
	Identities        []Identity     `gorm:"foreignKey:UserID;references:ID" json:"identities,omitempty"`
	PublishedBounties []Bounty       `gorm:"foreignKey:OwnerID" json:"published_bounties,omitempty"`
	Claims            []Claim        `gorm:"foreignKey:ClaimerID" json:"claims,omitempty"`
	RefreshTokens     []RefreshToken `gorm:"foreignKey:UserID" json:"refresh_tokens,omitempty"`
}

func init() {
	db.DB.AutoMigrate(&User{})
}
