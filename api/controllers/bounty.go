package controllers

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/bount-ing/bount.ing/api/db"
	"github.com/bount-ing/bount.ing/api/models"
)

func CreateBounty(bounty *models.Bounty) error {
	bounty.Status = "open"
	ownerStripeAccountID := ""
	ownerStripeCustomerID := ""
	stripeAccountFound := false

	// Check if the bounty issue url is empty
	if bounty.IssueURL == "" && bounty.IssueID != 0 {
		//retreive the issue url from the issue
		issue, err := GetIssueByID(bounty.IssueID)
		if err != nil {
			log.Printf("Error fetching issue: %s", err)
			return err
		}
		bounty.IssueURL = issue.URL
	} else if bounty.IssueURL != "" && bounty.IssueID == 0 {
		//retreive the issue id from the issue url
		issue, err := GetIssueByURL(bounty.IssueURL)
		if err != nil {
			log.Printf("Error fetching issue: %s", err)
			return err
		}
		bounty.IssueID = issue.ID
	}

	log.Printf("Creating Bounty - Owner ID: %d", bounty.OwnerID)
	// Fetch users identities where the host is stripe
	identities, err := GetUserIdentities(bounty.OwnerID)
	if err != nil {
		log.Printf("Error fetching user identities: %s", err)
		return err
	}

	for _, identity := range identities {
		log.Printf("Identity: %+v", identity)
		if identity.Host.Address == "https://stripe.com" {
			ownerStripeAccountID = identity.UserExternalID
			ownerStripeCustomerID = identity.UserExternalSecondaryID
			stripeAccountFound = true
		}
	}

	if !stripeAccountFound {
		log.Printf("User does not have a stripe account")
		return errors.New("user does not have a stripe account")
	}

	// Step 1: Create a SetupIntent in Stripe
	setupIntent, err := CreateStripeSetupIntent(bounty, bounty.OwnerID, ownerStripeCustomerID, ownerStripeAccountID)
	if err != nil {
		log.Printf("Error creating Stripe SetupIntent: %s", err)
		return err
	}

	// Step 2: Store the SetupIntent ID in the bounty
	bounty.StripeInvoiceID = setupIntent.ID

	// Step 3: Save the bounty in the database
	if err := db.DB.Create(&bounty).Error; err != nil {
		log.Printf("Error creating bounty: %s", err)
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

func DeleteBounty(bountyID, userID uint) error {
	var bounty models.Bounty

	err := db.DB.First(&bounty, bountyID)

	if err.Error != nil {
		log.Print(err.Error)
		return err.Error
	}

	if bounty.OwnerID != userID {
		return errors.New("user does not own bounty")
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

	err := db.DB.
		Preload("Bounties", "status = ?", "open").
		Preload("Bounties.Variables").
		Find(&issues)

	if err.Error != nil {
		log.Print(err.Error)
		return issues, err.Error
	}

	return issues, nil
}

func GetBountyByID(bountyID uint) (models.Bounty, error) {
	var bounty models.Bounty

	err := db.DB.First(&bounty, bountyID)

	if err.Error != nil {
		log.Print(err.Error)
		return bounty, err.Error
	}

	return bounty, nil
}

const (
	Crescendo   string = "crescendo"
	Flat        string = "flat"
	Decrescendo string = "decrescendo"
)

func GetCurrentBountyAmount(bountyID uint) (float64, error) {
	var bounty models.Bounty

	// Fetch the bounty details
	err := db.DB.First(&bounty, bountyID).Error
	if err != nil {
		log.Print(err)
		return 0, err
	}

	// Initialize the total bounty amount with the fixed amount
	totalAmount := bounty.Amount

	// Get the current time
	currentTime := time.Now()

	// Loop through all variable amounts for this bounty
	for _, variable := range bounty.Variables {
		// Check if current time is within the variable's timeframe
		if currentTime.After(variable.StartAt) && currentTime.Before(variable.EndAt) {
			// Calculate the variable bounty depending on the direction
			switch variable.Direction {
			case Crescendo:
				// Calculate Crescendo (increasing)
				// Assuming some linear increase for simplicity
				duration := variable.EndAt.Sub(variable.StartAt)
				elapsed := currentTime.Sub(variable.StartAt)
				totalAmount += float64(variable.Amount) * (elapsed / duration).Seconds()
			case Flat:
				// Flat means no change, so we just add the full amount
				totalAmount += float64(variable.Amount)
			case Decrescendo:
				// Calculate Decrescendo (decreasing)
				// Assuming a linear decrease
				duration := variable.EndAt.Sub(variable.StartAt)
				elapsed := currentTime.Sub(variable.StartAt)
				totalAmount += float64(variable.Amount) * (1 - (elapsed / duration).Seconds())
			}
		}
	}

	return totalAmount, nil
}
