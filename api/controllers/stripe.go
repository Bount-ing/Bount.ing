package controllers

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/bount-ing/bount.ing/api/db"
	"github.com/bount-ing/bount.ing/api/models"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/paymentintent"
	"github.com/stripe/stripe-go/paymentmethod"
	"github.com/stripe/stripe-go/setupintent"
)

func CreateStripeSetupIntent(bounty *models.Bounty, sponsorID uint, sponsorStripeID string, connectedAccountID string) (*stripe.SetupIntent, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	cust, err := customer.Get(sponsorStripeID, nil)
	if err != nil {
		return nil, err
	}

	paymentMethods, err := GetPaymentMethods(sponsorID)
	if err != nil {
		return nil, err
	}

	// check len of paymentMethods
	if len(paymentMethods) == 0 {
		log.Printf("No payment methods found for customer %s", sponsorStripeID)
		return nil, errors.New("no payment methods found for customer")
	}

	//list payment methods
	for _, pm := range paymentMethods {
		log.Printf("Payment Method: %+v", pm)
	}

	// If customer already has a default payment method, check if it's still valid
	if cust.InvoiceSettings.DefaultPaymentMethod == nil {
		log.Printf("Customer %s does not have a default payment method, setting one now...", sponsorStripeID)

		if len(paymentMethods) > 0 {
			pmID := paymentMethods[0].ID // Pick the first available payment method

			// Attach the payment method to the customer
			_, err := paymentmethod.Attach(pmID, &stripe.PaymentMethodAttachParams{
				Customer: stripe.String(sponsorStripeID),
			})
			if err != nil {
				return nil, err
			}
			log.Printf("Successfully attached payment method %s to customer %s", pmID, sponsorStripeID)

			// Set the attached payment method as the default
			_, err = customer.Update(sponsorStripeID, &stripe.CustomerParams{
				InvoiceSettings: &stripe.CustomerInvoiceSettingsParams{
					DefaultPaymentMethod: stripe.String(pmID),
				},
			})
			if err != nil {
				return nil, err
			}
			log.Printf("Set default payment method %s for customer %s", pmID, sponsorStripeID)
		} else {
			return nil, models.ErrNoPaymentMethods
		}

	}

	log.Printf("Creating SetupIntent for customer %s", sponsorStripeID)

	// Create SetupIntent params with all required fields
	params := &stripe.SetupIntentParams{
		Customer:           stripe.String(sponsorStripeID),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		Usage:              stripe.String("off_session"), // Save for future payments
	}

	log.Printf("Creating SetupIntent for customer %s in connected account %s",
		sponsorStripeID, connectedAccountID)

	setupIntent, err := setupintent.New(params)
	if err != nil {
		return nil, err
	}

	log.Printf("Successfully created SetupIntent %s for customer %s in account %s",
		setupIntent.ID, sponsorStripeID, connectedAccountID)

	return setupIntent, nil
}

func PayoutBounty(bounty models.Bounty, sponsorStripeCustomerID, bountyHunterStripeCustomerID, bountyHunterStripeAccountID string) error {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	// Validate inputs
	if bountyHunterStripeAccountID == "" {
		return errors.New("stripe connected account id is required for payouts")
	}

	customerParams := &stripe.CustomerParams{}

	customer, err := customer.Get(sponsorStripeCustomerID, customerParams)
	if err != nil {
		log.Printf("Failed to retrieve Stripe Customer: %v", err)
		return err
	}

	if customer.InvoiceSettings.DefaultPaymentMethod == nil {
		log.Printf("No default payment method found for customer %s", sponsorStripeCustomerID)
		return errors.New("no default payment method found for customer")
	}

	paymentMethodID := customer.InvoiceSettings.DefaultPaymentMethod.ID

	// Create PaymentIntent
	params := &stripe.PaymentIntentParams{
		Amount:               stripe.Int64(int64(bounty.ClaimedAmount * 100)),
		Currency:             stripe.String(bounty.Currency),
		Customer:             stripe.String(sponsorStripeCustomerID),
		PaymentMethod:        stripe.String(paymentMethodID),
		Confirm:              stripe.Bool(true),
		ApplicationFeeAmount: stripe.Int64(int64(bounty.ClaimedAmount * 100)),
		TransferData: &stripe.PaymentIntentTransferDataParams{
			Destination: stripe.String(bountyHunterStripeAccountID),
		},
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		OffSession:         stripe.Bool(true),
	}

	paymentIntent, err := paymentintent.New(params)
	if err != nil {
		log.Printf("Failed to create PaymentIntent: %v", err)
		return err
	}

	log.Printf("Created PaymentIntent: %s with status: %s", paymentIntent.ID, paymentIntent.Status)

	// Check payment status
	if paymentIntent.Status != stripe.PaymentIntentStatusSucceeded {
		return models.ErrPaymentIntent
	}

	// Update bounty status
	bounty.Status = "paid"
	bounty.FinalizedAt = time.Now()

	if err := db.DB.Save(&bounty).Error; err != nil {
		log.Printf("Failed to update bounty status: %v", err)
		return models.ErrBountyUpdateFailed
	}

	log.Printf("Successfully processed payment for bounty %d", bounty.ID)
	return nil
}

func GetOrCreateCustomerID(accessToken, stripeAccountID, email string) (string, error) {
	// List customers with email filter in the connected account
	params := &stripe.CustomerListParams{
		Email: stripe.String(email),
	}
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

	// Add additional metadata to track the customer
	customerParams.AddMetadata("created_by", "platform")
	customerParams.AddMetadata("connected_account", stripeAccountID)

	newCustomer, err := customer.New(customerParams)
	if err != nil {
		return "", models.ErrStripeFailedToCreateCustomer
	}

	log.Printf("Created new customer ID: %s with email: %s in account: %s",
		newCustomer.ID, email, stripeAccountID)

	// Verify the customer was created by trying to retrieve it
	verifyParams := &stripe.CustomerParams{}
	_, err = customer.Get(newCustomer.ID, verifyParams)
	if err != nil {
		return "", models.ErrStripeFailedToVerifyCustomer
	}

	return newCustomer.ID, nil
}
