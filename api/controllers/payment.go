package controllers

import (
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
	bountyAmount := bounty.Amount

	paymentMsg := fmt.Sprintf("You have been paid %d for Bounty %d", bountyAmount, bountyID)

	tools.SendEmail(claimerEmail, "Payment Confirmation", paymentMsg)

	return nil
}
