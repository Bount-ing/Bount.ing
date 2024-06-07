package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	GithubID          int      `json:"github_id"`
	Username          string   `json:"username" gorm:"unique;not null"`
	Email             string   `json:"email" gorm:"unique;not null"`
	PublishedBounties []Bounty `json:"published_bounties" gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
