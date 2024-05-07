package models

import (
    "gorm.io/gorm"
)

type Bounty struct {
    gorm.Model
    Amount      float64 `json:"amount" gorm:"not null"`
    Currency    string  `json:"currency" gorm:"default:'USD'"`
    IssueID     uint    `json:"issue_id"`
    Issue       Issue   `gorm:"foreignKey:IssueID"`
    UserID      uint    `json:"user_id"`
    User        User    `gorm:"foreignKey:UserID"`
    Status      string  `json:"status" gorm:"default:'open'"` // e.g., open, claimed, cancelled
}
