package services

import (
	"errors"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/invoice"
)

// NOTE: Should be called when the issue is closed with a pull request.
//		 Maybe this should be Claim instead of Transaction
//		 Creates an invoice, and a transaction object set as pending. A web hook
//		 of Stripe should recive an update and change the state.
//		 Recives the users and the bounty / Claim
func CreateTransaction() error {
	return errors.New("NOT IMPLEMENTED")
}

// createInvoice Creates an Invoice and exectues it taking a fee of 8%
// Recives IDs of the sender and destination accounts, title of the issue and the amount of the bounty
func createInvoice(sender, dest, title string, amount int64) (*stripe.Invoice, error) {
	params := stripe.InvoiceParams{
		Customer:             stripe.String(sender),
		ApplicationFeeAmount: stripe.Int64(amount * 8 / 100),
		AutoAdvance:          stripe.Bool(true),
		Description:          stripe.String("[Bount.ing] Bounty claimed for issue: " + title),
		TransferData: &stripe.InvoiceTransferDataParams{
			Amount:      stripe.Int64(amount),
			Destination: stripe.String(dest),
		},
	}
	i, err := invoice.New(&params)
	if err != nil {
		return nil, err
	}
	i, err = invoice.FinalizeInvoice(i.ID, nil)
	if err != nil {
		return nil, err
	}
	i, err = invoice.Pay(i.ID, nil)
	if err != nil {
		return nil, err
	}
	return i, nil
}
