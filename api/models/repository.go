package models

import (
    "gorm.io/gorm"
)

type Repository struct {
    gorm.Model
    Name           string    `json:"name" gorm:"not null"`
    OrganizationID uint      `json:"organization_id"`
    Organization   Organization `gorm:"foreignKey:OrganizationID"`
    Issues         []Issue   `json:"issues"`
}

