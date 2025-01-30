package controllers

import (
	"errors"
	"os"
	"time"

	"github.com/bount-ing/bount.ing/api/db"
	"github.com/bount-ing/bount.ing/api/models"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/paymentintent"
	"github.com/stripe/stripe-go/setupintent"
)

func CreateStripeSetupIntent(bounty *models.Bounty, payerStripeCustomerID string) (*stripe.SetupIntent, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	params := &stripe.SetupIntentParams{
		Customer:           stripe.String(payerStripeCustomerID),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		Usage:              stripe.String("off_session"), // Save for future payments
	}

	setupIntent, err := setupintent.New(params)
	if err != nil {
		return nil, errors.New("failed to create Stripe SetupIntent: " + err.Error())
	}

	return setupIntent, nil
}

func PayoutBounty(bounty models.Bounty, payerStripeCustomerID, stripeConnectedAccountID string) error {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	// Fetch customer to get their default payment method
	customer, err := customer.Get(payerStripeCustomerID, nil)
	if err != nil {
		return errors.New("failed to retrieve Stripe Customer: " + err.Error())
	}

	if customer.InvoiceSettings.DefaultPaymentMethod == nil {
		return errors.New("no default payment method found for customer")
	}

	paymentMethodID := customer.InvoiceSettings.DefaultPaymentMethod.ID

	// Ensure that the connected account ID is provided
	if stripeConnectedAccountID == "" {
		return errors.New("stripe connected account id is required for payouts")
	}

	// Step 1: Create PaymentIntent (charge the payer)
	params := &stripe.PaymentIntentParams{
		Amount:               stripe.Int64(int64(bounty.ClaimedAmount * 100)), // Convert to cents
		Currency:             stripe.String(bounty.Currency),
		Customer:             stripe.String(payerStripeCustomerID),
		PaymentMethod:        stripe.String(paymentMethodID), // Corrected to fetch from Customer
		Confirm:              stripe.Bool(true),
		ApplicationFeeAmount: stripe.Int64(int64(bounty.ClaimedAmount * 0.05 * 100)), // Example 5% fee
		TransferData: &stripe.PaymentIntentTransferDataParams{
			Destination: stripe.String(stripeConnectedAccountID), // Send funds to the recipient
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
