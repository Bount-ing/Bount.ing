package models

import (
	"gorm.io/gorm"
)

type Organization struct {
	gorm.Model
	Name         string       `json:"name" gorm:"unique;not null"`
	Repositories []Repository `json:"repositories" gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
