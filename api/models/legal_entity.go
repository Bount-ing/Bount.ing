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

func init() {
	db.DB.AutoMigrate(&LegalEntity{})
}
