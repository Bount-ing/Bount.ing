package controllers

import (
	"errors"
	"fmt"
	"log"

	"github.com/bount-ing/bount.ing/api/db"
	"github.com/bount-ing/bount.ing/api/models"
)

func CreateBounty(bounty *models.Bounty) error {
	// Save the bounty and associated relationships to the database
	if err := db.DB.Create(&bounty).Error; err != nil {
		return err
	}
	return nil
}

func GetBounty(bountyID string) (models.Bounty, error) {
	var bounty models.Bounty

	err := db.DB.First(&bounty, bountyID)

	if err.Error != nil {
		log.Print(err.Error)
		return bounty, err.Error
	}

	return bounty, nil
}

func GetBounties() ([]models.Bounty, error) {
	var bounties []models.Bounty

	err := db.DB.Find(&bounties)

	if err.Error != nil {
		log.Print(err.Error)
		return bounties, err.Error
	}

	return bounties, nil
}

func UpdateBounty(bounty models.Bounty) error {
	err := db.DB.Save(&bounty)

	if err.Error != nil {
		log.Print(err.Error)
		return err.Error
	}

	return nil
}

func DeleteBounty(bountyID string) error {
	var bounty models.Bounty

	err := db.DB.First(&bounty, bountyID)

	if err.Error != nil {
		log.Print(err.Error)
		return err.Error
	}

	err = db.DB.Delete(&bounty)

	if err.Error != nil {
		log.Print(err.Error)
		return err.Error
	}

	return nil
}

func FinalizeBounty(bountyID string) error {
	var bounty models.Bounty

	err := db.DB.First(&bounty, bountyID)

	if err.Error != nil {
		log.Print(err.Error)
		return err.Error
	}

	bounty.FinalizedAt = db.DB.NowFunc()

	err = db.DB.Save(&bounty)

	if err.Error != nil {
		log.Print(err.Error)
		return err.Error
	}

	return nil
}

func GetAllUnconfirmedBounties() ([]models.Bounty, error) {
	var bounties []models.Bounty

	err := db.DB.Where("finalized = ?", false).Find(&bounties)

	if err.Error != nil {
		log.Print(err.Error)
		return bounties, err.Error
	}

	return bounties, nil
}

func ValidateBountyData(bounty models.Bounty, variables []models.BountyVariable) error {
	// Ensure start date is not after end date for bounty
	if bounty.StartAt.After(bounty.EndAt) {
		return errors.New("bounty start date cannot be after end date")
	}

	for _, variable := range variables {
		// Check if variable amount is greater than zero
		if variable.Amount <= 0 {
			return errors.New("variable amount must be greater than zero")
		}

		// Compare start and end dates for each variable directly without parsing
		if variable.StartAt.After(variable.EndAt) {
			return fmt.Errorf("variable start date %s cannot be after end date %s", variable.StartAt, variable.EndAt)
		}
	}

	return nil
}
func GetPublicBountiesByIssue() ([]models.Issue, error) {
	// Fetch all issues with associated bounties and their variables
	var issues []models.Issue

	err := db.DB.Preload("Bounties").Preload("Bounties.Variables").Find(&issues)

	if err.Error != nil {
		log.Print(err.Error)
		return issues, err.Error
	}

	return issues, nil
}
