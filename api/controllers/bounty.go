package controllers

import (
	"log"

	"github.com/bount-ing/bount.ing/api/db"
	"github.com/bount-ing/bount.ing/api/models"
)

func CreateBounty(bounty models.Bounty) error {

	err := db.DB.Create(&bounty)

	if err.Error != nil {
		log.Print(err.Error)
		return err.Error
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
