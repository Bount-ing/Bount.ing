package models

import (
	"time"

	"gorm.io/gorm"

	"github.com/bount-ing/bount.ing/api/db"
)

type User struct {
	gorm.Model
	Email                   string    `json:"email"`
	Password                string    `json:"-"` // "-" to never send password in JSON
	Admin                   bool      `json:"admin"`
	Verified                bool      `json:"verified"`
	VerifCode               string    `json:"-"`
	VerifCodeExpirationTime time.Time `json:"-"`
	StipeAccountID          string    `json:"stripe_account_id"`

	// Relationships
	ExternalIdentities []ExternalIdentity `gorm:"foreignKey:UserID;references:ID" json:"external_identities,omitempty"`
	PublishedBounties  []Bounty           `gorm:"foreignKey:OwnerID" json:"published_bounties,omitempty"`
	Claims             []Claim            `gorm:"foreignKey:ClaimerID" json:"claims,omitempty"`
	RefreshTokens      []RefreshToken     `gorm:"foreignKey:UserID" json:"refresh_tokens,omitempty"`

	// We'll keep Username at the User level as it might be used as a display name
	// across all platforms
	Username string `json:"username"`
}

func init() {
	db.DB.AutoMigrate(&User{})
}
