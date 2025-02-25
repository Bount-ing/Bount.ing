package models

import (
	"time"

	"gorm.io/gorm"

	"github.com/bount-ing/bount.ing/api/db"
)

type LegalEntity struct {
	gorm.Model

	IsCompany bool

	DocumentType    string
	DocumentCountry string
	DocumentNumber  string

	LegalName    string
	LegalAddress string
	LegalCity    string
	LegalZip     string
	LegalState   string
	LegalCountry string

	ConfirmedAt time.Time

	UserID uint `json:"-"`
}

func LegalEntitySeed() {
	var bounting LegalEntity
	if db.DB.Where("document_number = ?", "B19779453").First(&bounting).Error != nil {
		db.DB.Create(&LegalEntity{
			IsCompany:       true,
			DocumentType:    "CIF",
			DocumentCountry: "ES",
			DocumentNumber:  "B19779453",
			LegalName:       "BOUNT ING S.L.",
			LegalAddress:    "C/ Lepant 270",
			LegalCity:       "Barcelona",
			LegalZip:        "08013",
			LegalState:      "Barcelona",
			LegalCountry:    "ES",
		})
	}
}

func init() {
	db.DB.AutoMigrate(&LegalEntity{})
}
