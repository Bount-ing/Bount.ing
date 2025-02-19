package models

import (
	"gorm.io/gorm"

	"github.com/bount-ing/bount.ing/api/db"
)

type UserLegalData struct {
	gorm.Model

	TaxID           string `json:"tax_id"`
	DocumentType    string `json:"document_type"`
	DocumentCountry string `json:"document_country"`

	IsCompany      bool   `json:"is_company"`
	CompanyName    string `json:"company_name"`
	CompanyAddress string `json:"company_address"`
	CompanyCity    string `json:"company_city"`
	CompanyState   string `json:"company_state"`
	CompanyZip     string `json:"company_zip"`
	CompanyCountry string `json:"company_country"`
}

func init() {
	db.DB.AutoMigrate(&User{})
}
