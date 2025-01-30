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

	stripeBountyHunterID := ""
	stripeBountyHunterIDFound := false
	for _, identity := range bountyHunterIdentities {
		if identity.Host.Address == "https://stripe.com" {
			stripeBountyHunterID = identity.UserExternalID
			stripeBountyHunterIDFound = true
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

	bountyOwnerIdentities, err := GetUserIdentities(bounty.OwnerID)
	stripeBountyOwnerID := ""
	stripeBountyOwnerIDFound := false
	for _, identity := range bountyOwnerIdentities {
		if identity.Host.Address == "https://stripe.com" {
			stripeBountyOwnerID = identity.UserExternalID
			stripeBountyOwnerIDFound = true
		}
	}

	if !stripeBountyOwnerIDFound {
		bounty.Status = "payment_pending"
		err = UpdateBounty(bounty)
		if err != nil {
			log.Print(err)
		}
		return errors.New("user does not have a stripe account")
	}

	//func PayoutBounty(bounty *models.Bounty, payerStripeCustomerID, stripeConnectedAccountID string) error

	PayoutBounty(bounty, stripeBountyHunterID, stripeBountyOwnerID)

	paymentMsg := fmt.Sprintf("You have been paid %.f for Bounty %d", bounty.ClaimedAmount, bountyID)

	tools.SendEmail(claimerEmail, "Payment Confirmation", paymentMsg)

	return nil
}
