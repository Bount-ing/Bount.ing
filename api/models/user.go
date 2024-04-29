package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username  string    `json:"username" gorm:"unique;not null"`
    Email     string    `json:"email" gorm:"unique;not null"`
    Password  string    `json:"-" gorm:"not null"` // Password is not returned in JSON responses
    Bounties  []Bounty  `json:"bounties"` // User can have multiple bounties
}

