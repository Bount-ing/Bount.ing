package controllers

import (
	"github.com/bount-ing/bount.ing/api/db"
	"github.com/bount-ing/bount.ing/api/models"
)

func GetUserIdentities(userID uint) ([]models.Identity, error) {
	var identities []models.Identity

	if err := db.DB.Where("user_id = ?", userID).Find(&identities).Error; err != nil {
		return identities, err
	}

	return identities, nil
}
func GetHostIdentitiesFromAddress(hostAddress string) ([]models.Identity, error) {
	var identities []models.Identity

	if err := db.DB.
		Joins("JOIN hosts ON hosts.id = identities.host_id").
		Where("hosts.address = ?", hostAddress).
		Find(&identities).Error; err != nil {
		return identities, err
	}

	return identities, nil
}

func CreateIdentity(identity *models.Identity) error {
	if err := db.DB.Create(identity).Error; err != nil {
		return err
	}

	return nil
}
