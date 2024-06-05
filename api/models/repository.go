package models

import (
	"gorm.io/gorm"
)

type Repository struct {
	gorm.Model
	Name           string  `json:"name" gorm:"not null"`
	OrganizationID uint    `json:"organization_id"`
	Issues         []Issue `json:"issues" gorm:"foreignKey:RepositoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
