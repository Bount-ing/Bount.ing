package controllers

import (
	"errors"
	"log"
	"time"

	"github.com/bount-ing/bount.ing/api/db"
	"github.com/bount-ing/bount.ing/api/models"
)

func CreateBounty(bounty *models.Bounty) error {
	bounty.Status = "open"
	sponsorStripeAccountID := ""
	sponsorStripeCustomerID := ""
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

	log.Printf("Creating Bounty - Sponsor ID: %d", bounty.SponsorID)
	// Fetch users identities where the host is stripe
	identities, err := GetUserIdentities(bounty.SponsorID)
	if err != nil {
		log.Printf("Error fetching user identities: %s", err)
		return err
	}

	for _, identity := range identities {
		log.Printf("Identity: %+v", identity)
		if identity.Host.Address == "https://stripe.com" {
			sponsorStripeAccountID = identity.UserExternalID
			sponsorStripeCustomerID = identity.UserExternalSecondaryID
			stripeAccountFound = true
		}
	}

	if !stripeAccountFound {
		log.Printf("User does not have a stripe account")
		return errors.New("user does not have a stripe account")
	}

	// Step 1: Create a SetupIntent in Stripe
	setupIntent, err := CreateStripeSetupIntent(bounty, bounty.SponsorID, sponsorStripeCustomerID, sponsorStripeAccountID)
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

	if bounty.SponsorID != userID {
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
			return models.ErrDateRangeStartAfterEnd
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
	Crescendo   string = "increase"
	Flat        string = "flat"
	Decrescendo string = "decrease"
)

func GetCurrentBountyAmount(bountyID uint) (float64, error) {
	var bounty models.Bounty

	// Fetch the bounty details and preload the related variables
	err := db.DB.Preload("Variables").First(&bounty, bountyID).Error
	if err != nil {
		log.Print(err)
		return 0, err
	}

	// Initialize the total bounty amount with the fixed amount
	totalAmount := bounty.Amount

	// Get the current time (ensure it matches the timezone of StartAt/EndAt)
	currentTime := time.Now()

	// Loop through all variable amounts for this bounty
	for _, variable := range bounty.Variables {
		// Log the current variable and time information
		log.Printf("Processing variable: Amount=%v, Direction=%v, StartAt=%v, EndAt=%v", variable.Amount, variable.Direction, variable.StartAt, variable.EndAt)

		// Check if current time is within the variable's timeframe
		if currentTime.After(variable.StartAt) && currentTime.Before(variable.EndAt) {
			log.Printf("Variable is within timeframe: CurrentTime=%v, StartAt=%v, EndAt=%v", currentTime, variable.StartAt, variable.EndAt)

			// Calculate the variable bounty depending on the direction
			switch variable.Direction {
			case Crescendo:
				// Calculate Crescendo (increasing)
				duration := variable.EndAt.Sub(variable.StartAt)
				elapsed := currentTime.Sub(variable.StartAt)

				log.Printf("Crescendo calculation: elapsed=%v seconds, duration=%v seconds", elapsed.Seconds(), duration.Seconds())

				if duration.Seconds() > 0 {
					increaseAmount := float64(variable.Amount) * (elapsed.Seconds() / duration.Seconds())
					log.Printf("Crescendo increaseAmount=%v", increaseAmount)
					totalAmount += increaseAmount
				}
			case Flat:
				// Flat means no change, so we just add the full amount
				log.Printf("Flat: Adding full amount=%v", float64(variable.Amount))
				totalAmount += float64(variable.Amount)
			case Decrescendo:
				// Calculate Decrescendo (decreasing)
				duration := variable.EndAt.Sub(variable.StartAt)
				elapsed := currentTime.Sub(variable.StartAt)

				log.Printf("Decrescendo calculation: elapsed=%v seconds, duration=%v seconds", elapsed.Seconds(), duration.Seconds())

				if duration.Seconds() > 0 {
					decreaseAmount := float64(variable.Amount) * (1 - (elapsed.Seconds() / duration.Seconds()))
					log.Printf("Decrescendo decreaseAmount=%v", decreaseAmount)
					totalAmount += decreaseAmount
				}
			}
		} else {
			log.Printf("Variable is not within timeframe, skipping: CurrentTime=%v, StartAt=%v, EndAt=%v", currentTime, variable.StartAt, variable.EndAt)
		}
	}

	// Final log of the total amount
	log.Printf("Final total bounty amount: %v", totalAmount)

	return totalAmount, nil
}
