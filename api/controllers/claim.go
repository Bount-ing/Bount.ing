package controllers

import (
	"log"

	"github.com/bount-ing/bount.ing/api/db"
	"github.com/bount-ing/bount.ing/api/models"
)

func CreateClaim(claim models.Claim) error {
	err := db.DB.Create(&claim)

	if err.Error != nil {
		log.Print(err.Error)
		return err.Error
	}
	return nil
}

func GetClaim(claimID string) (models.Claim, error) {
	var claim models.Claim

	err := db.DB.First(&claim, claimID)

	if err.Error != nil {
		log.Print(err.Error)
		return claim, err.Error
	}
	return claim, nil
}

func GetClaims() ([]models.Claim, error) {
	var claims []models.Claim

	err := db.DB.Find(&claims)

	if err.Error != nil {
		log.Print(err.Error)
		return claims, err.Error
	}

	return claims, nil
}

func UpdateClaim(claim models.Claim) error {
	err := db.DB.Save(&claim)

	if err.Error != nil {
		log.Print(err.Error)
		return err.Error
	}
	return nil
}

func DeleteClaim(claimID string) error {
	var claim models.Claim
	err := db.DB.First(&claim, claimID)
	if err.Error != nil {
		log.Print(err.Error)
		return err.Error
	}
	err = db.DB.Delete(&claim)
	if err.Error != nil {
		log.Print(err.Error)
		return err.Error
	}
	return nil
}
