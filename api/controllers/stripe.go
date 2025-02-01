package controllers

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bount-ing/bount.ing/api/db"
	"github.com/bount-ing/bount.ing/api/models"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/paymentintent"
	"github.com/stripe/stripe-go/setupintent"
)

func CreateStripeSetupIntent(bounty *models.Bounty, bountyOwnerStripeID string, connectedAccountID string) (*stripe.SetupIntent, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	// Create SetupIntent params with all required fields
	params := &stripe.SetupIntentParams{
		Customer:           stripe.String(bountyOwnerStripeID),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		Usage:              stripe.String("off_session"), // Save for future payments
	}

	// Set the Stripe Connect account context
	params.SetStripeAccount(connectedAccountID)

	log.Printf("Creating SetupIntent for customer %s in connected account %s",
		bountyOwnerStripeID, connectedAccountID)

	setupIntent, err := setupintent.New(params)
	if err != nil {
		return nil, fmt.Errorf("failed to create Stripe SetupIntent in account %s: %v",
			connectedAccountID, err)
	}

	log.Printf("Successfully created SetupIntent %s for customer %s in account %s",
		setupIntent.ID, bountyOwnerStripeID, connectedAccountID)

	return setupIntent, nil
}

func PayoutBounty(bounty models.Bounty, bountyOwnerStripeID, bountyHunterStripeID string) error {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	// Fetch customer to get their default payment method
	customer, err := customer.Get(bountyOwnerStripeID, nil)
	if err != nil {
		return errors.New("failed to retrieve Stripe Customer: " + err.Error())
	}

	if customer.InvoiceSettings.DefaultPaymentMethod == nil {
		return errors.New("no default payment method found for customer")
	}

	paymentMethodID := customer.InvoiceSettings.DefaultPaymentMethod.ID

	// Ensure that the connected account ID is provided
	if bountyHunterStripeID == "" {
		return errors.New("stripe connected account id is required for payouts")
	}

	// Step 1: Create PaymentIntent (charge the payer)
	params := &stripe.PaymentIntentParams{
		Amount:               stripe.Int64(int64(bounty.ClaimedAmount * 100)), // Convert to cents
		Currency:             stripe.String(bounty.Currency),
		Customer:             stripe.String(bountyOwnerStripeID),
		PaymentMethod:        stripe.String(paymentMethodID), // Corrected to fetch from Customer
		Confirm:              stripe.Bool(true),
		ApplicationFeeAmount: stripe.Int64(int64(bounty.ClaimedAmount * 0.05 * 100)), // Example 5% fee
		TransferData: &stripe.PaymentIntentTransferDataParams{
			Destination: stripe.String(bountyHunterStripeID), // Send funds to the recipient
		},
	}

	_, err = paymentintent.New(params)
	if err != nil {
		return errors.New("failed to create PaymentIntent: " + err.Error())
	}

	// Step 2: Mark bounty as finalized
	bounty.Status = "paid"
	bounty.FinalizedAt = time.Now()

	if err := db.DB.Save(&bounty).Error; err != nil {
		return err
	}

	return nil
}

func GetOrCreateCustomerID(accessToken, stripeAccountID, email string) (string, error) {
	// List customers with email filter in the connected account
	params := &stripe.CustomerListParams{
		Email: stripe.String(email),
	}
	params.SetStripeAccount(stripeAccountID) // Essential: search in connected account
	log.Printf("List customers Params: %+v", params)

	i := customer.List(params)
	for i.Next() {
		c := i.Customer()
		log.Printf("Found Customer: %+v", c)
		return c.ID, nil
	}

	// No customer found, create a new one in the connected account
	log.Printf("No customer found with email %s in account %s, creating new one...",
		email, stripeAccountID)

	customerParams := &stripe.CustomerParams{
		Email: stripe.String(email),
	}
	customerParams.SetStripeAccount(stripeAccountID)

	// Add additional metadata to track the customer
	customerParams.AddMetadata("created_by", "platform")
	customerParams.AddMetadata("connected_account", stripeAccountID)

	newCustomer, err := customer.New(customerParams)
	if err != nil {
		return "", fmt.Errorf("failed to create customer in connected account %s: %v",
			stripeAccountID, err)
	}

	log.Printf("Created new customer ID: %s with email: %s in account: %s",
		newCustomer.ID, email, stripeAccountID)

	// Verify the customer was created by trying to retrieve it
	verifyParams := &stripe.CustomerParams{}
	verifyParams.SetStripeAccount(stripeAccountID)
	_, err = customer.Get(newCustomer.ID, verifyParams)
	if err != nil {
		return "", fmt.Errorf("failed to verify new customer in connected account: %v", err)
	}

	return newCustomer.ID, nil
}
