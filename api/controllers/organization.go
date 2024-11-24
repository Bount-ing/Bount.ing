package controllers

import (
	"log"

	"github.com/bount-ing/bount.ing/api/db"
	"github.com/bount-ing/bount.ing/api/models"
)

func CreateOrganization(organization models.Organization) error {
	err := db.DB.Create(&organization)
	if err.Error != nil {
		log.Print(err.Error)
		return err.Error
	}
	return nil
}

func GetOrganization(organizationID string) (models.Organization, error) {
	var organization models.Organization
	err := db.DB.First(&organization, organizationID)
	if err.Error != nil {
		log.Print(err.Error)
		return organization, err.Error
	}
	return organization, nil
}

func GetOrganizations() ([]models.Organization, error) {
	var organizations []models.Organization
	err := db.DB.Find(&organizations)
	if err.Error != nil {
		log.Print(err.Error)
		return organizations, err.Error
	}
	return organizations, nil
}

func UpdateOrganization(organization models.Organization) error {
	err := db.DB.Save(&organization)
	if err.Error != nil {
		log.Print(err.Error)
		return err.Error
	}
	return nil
}

func DeleteOrganization(organizationID string) error {
	var organization models.Organization
	err := db.DB.First(&organization, organizationID)
	if err.Error != nil {
		log.Print(err.Error)
		return err.Error
	}
	return nil
}
