package models

import (
	"gorm.io/gorm"

	"github.com/bount-ing/bount.ing/api/db"
)

// Invoice Data must be immutable for history and legal reason
type Invoice struct {
	gorm.Model

	InvoiceNumber string
	InvoiceDate   string

	// Emitter
	EmitterName            string
	EmitterTaxID           string
	EmitterVIESCountryCode string
	EmitterVIESValid       bool

	// Address
	EmitterStreet string
	EmitterNumber string
	EmitterFloor  string
	EmitterDoor   string

	EmitterCity string
	EmitterZip  string

	EmitterState   string
	EmitterCountry string

	// Receiver
	ReceiverName            string
	ReceiverTaxID           string
	ReceiverVIESCountryCode string
	ReceiverVIESValid       bool

	// Address
	ReceiverStreet string
	ReceiverNumber string
	ReceiverFloor  string
	ReceiverDoor   string

	ReceiverCity string
	ReceiverZip  string

	ReceiverState   string
	ReceiverCountry string
}

func init() {
	db.DB.AutoMigrate(&Invoice{})
}
