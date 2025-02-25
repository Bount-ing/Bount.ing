package controllers

import (
	"errors"
	"fmt"
	"log"

	"github.com/bount-ing/bount.ing/api/tools"
)

func ProcessPayment(claimerID uint, bountyID uint) error {
	//retrieve bounty amount
	bounty, err := GetBountyByID(bountyID)
	if err != nil {
		log.Print(err)
		return err
	}

	//retrieve claimer
	claimer, err := GetUserByID(claimerID)
	if err != nil {
		log.Print(err)
		return err
	}

	claimerEmail := claimer.Email

	bountyHunterIdentities, err := GetUserIdentities(claimerID)
	if err != nil {
		log.Print(err)
	}

	stripeBountyHunterAccountID := ""
	stripeBountyHunterCustomerID := ""
	stripeBountyHunterIDFound := false
	for _, identity := range bountyHunterIdentities {
		if identity.Host.Address == "https://stripe.com" {
			stripeBountyHunterAccountID = identity.UserExternalID
			stripeBountyHunterCustomerID = identity.UserExternalSecondaryID
			stripeBountyHunterIDFound = true
			log.Printf("Found stripe account for bounty hunter %d\n- Account ID: %s\n- Customer ID: %s", claimerID, stripeBountyHunterAccountID, stripeBountyHunterCustomerID)
		}
	}

	if !stripeBountyHunterIDFound {
		//set bounty status to "payment_pending"
		bounty.Status = "payment_pending"
		err = UpdateBounty(bounty)
		if err != nil {
			log.Print(err)
		}
		return errors.New("user does not have a stripe account")
	}

	sponsorIdentities, err := GetUserIdentities(bounty.SponsorID)
	stripeBountySponsorAccountID := ""
	stripeBountySponsorCustomerID := ""
	stripeBountySponsorIDFound := false
	for _, identity := range sponsorIdentities {
		if identity.Host.Address == "https://stripe.com" {
			stripeBountySponsorAccountID = identity.UserExternalID
			stripeBountySponsorCustomerID = identity.UserExternalSecondaryID
			stripeBountySponsorIDFound = true
			log.Printf("Found stripe account for bounty sponsor %d\n- Account ID: %s\n- Customer ID: %s", bounty.SponsorID, stripeBountySponsorAccountID, stripeBountySponsorCustomerID)
		}
	}

	if !stripeBountySponsorIDFound {
		bounty.Status = "payment_pending"
		err = UpdateBounty(bounty)
		if err != nil {
			log.Print(err)
		}
		return errors.New("user does not have a stripe account")
	}

	//func PayoutBounty(bounty *models.Bounty, payerStripeCustomerID, stripeConnectedAccountID string) error

	PayoutBounty(bounty, stripeBountySponsorCustomerID, stripeBountyHunterCustomerID, stripeBountyHunterAccountID)

	paymentMsg := fmt.Sprintf("You have been paid %.2f for Bounty %d", bounty.ClaimedAmount, bountyID)

	tools.SendEmail(claimerEmail, "Payment Confirmation", paymentMsg)

	return nil
}
