package tools

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bount-ing/bount.ing/api/db"
	"github.com/bount-ing/bount.ing/api/models"
	"github.com/thdelmas/pdf-invoicer/pdfinvoicer"

	"gopkg.in/mail.v2"
)

type EmailInvoiceData struct {
	To            string
	InvoiceNumber string
	Amount        float64
	CustomerName  string
}

type InvoiceData struct {
	InvoiceNumber   string
	CustomerName    string
	CustomerAddress string
	CustomerNIF     string
	Amount          float64
	VAT             float64
}

func GenerateInvoice(bountyID, claimID uint) error {

	// Bount.ing Data
	companyAddress, err := pdfinvoicer.NewAddress(
		"C/ Lepant",
		"N. 270",
		"",
		"",
		"",
		"08013",
		"Barcelona",
		"Barcelona",
		"España")
	if err != nil {
		log.Println(err)
	}

	company, err := pdfinvoicer.NewIssuer(
		"BOUNT ING S.L.",
		companyAddress,
		"B19779453",
		"",
		"contact@bount.ing",
		"+34 614 10 70 09")

	if err != nil {
		log.Println(err)
	}

	// Get required data from database
	var (
		bounty             models.Bounty
		claim              models.Claim
		claimer            models.User
		claimerLegalEntity models.LegalEntity
	)

	if err := db.DB.First(&bounty, bountyID).Error; err != nil {
		return models.ErrBountyNotFound
	}

	if err := db.DB.First(&claim, claimID).Error; err != nil {
		return models.ErrClaimNotFound
	}

	if err := db.DB.First(&claimer, claim.ClaimerID).Error; err != nil {
		return models.ErrUserNotFound
	}

	if err := db.DB.Where("user_id = ?", claimer.ID).First(&claimerLegalEntity).Error; err != nil {
		return models.ErrLegalEntityNotFound
	}

	//Create customer
	clientAddress, err := pdfinvoicer.NewAddress(
		claimerLegalEntity.LegalAddress,
		"",
		"",
		"",
		"",
		claimerLegalEntity.LegalZip,
		claimerLegalEntity.LegalCity,
		claimerLegalEntity.LegalState,
		claimerLegalEntity.LegalCountry)
	if err != nil {
		log.Println(err)
	}

	client, err := pdfinvoicer.NewClient(
		claimerLegalEntity.LegalName,
		clientAddress,
		claimerLegalEntity.DocumentNumber)
	if err != nil {
		log.Println(err)
	}

	currentYear := time.Now().Year()
	currentMonth := time.Now().Month()

	invoiceNumber := GetNextInvoiceNumber(currentYear, currentMonth)
	invoiceDate := time.Now()

	opDate := claim.CreatedAt
	dueDate := invoiceDate.AddDate(0, 0, 30)

	invoice, err := pdfinvoicer.NewInvoice(
		invoiceNumber,
		invoiceDate,
		opDate,
		dueDate,
		company,
		client,
		[]pdfinvoicer.Item{
			{
				Description: "Commission Fees 4.2%\nIssue #1 Resolution",
				Quantity:    claim.ClaimedAmount,
				UnitPrice:   0.042,
				VATRate:     0.21,
			},
		},
		"This Invoice has already been paid through Stripe. No further action is required.",
		"",
	)
	if err != nil {
		log.Println(err)
	}
	total := claim.ClaimedAmount * 0.042 * 1.21

	outputFile := fmt.Sprintf("invoice_%s.pdf", invoiceNumber)
	invoice.GeneratePDF(outputFile)

	// Send email with invoice
	emailData := EmailInvoiceData{
		To:            claimer.Email,
		InvoiceNumber: invoiceNumber,
		Amount:        total,
		CustomerName:  claimerLegalEntity.LegalName,
	}

	if err := SendInvoiceEmail(emailData, outputFile); err != nil {
		return fmt.Errorf("failed to send invoice email: %v", err)
	}

	return nil
}

func SendInvoiceEmail(data EmailInvoiceData, pdfPath string) error {
	m := mail.NewMessage()
	m.SetHeader("From", os.Getenv("NOREPLY_MAIL_ADDRESS"))
	m.SetHeader("To", data.To)
	m.SetHeader("Subject", fmt.Sprintf("Invoice #%s from Bount.ing", data.InvoiceNumber))

	// Create the email content
	content := fmt.Sprintf(`
		<h2>Invoice #%s</h2>
		<p>Dear %s,</p>
		<p>Please find attached your invoice for %.2f€.</p>
		<p>Thank you for using Bount.ing!</p>
		<p>If you have any questions, please don't hesitate to contact us.</p>
	`, data.InvoiceNumber, data.CustomerName, data.Amount)

	m.SetBody("text/html", fmt.Sprintf(MailTemplate, content, time.Now().Year()))

	// Attach the PDF file
	if _, err := os.Stat(pdfPath); err == nil {
		m.Attach(pdfPath,
			mail.Rename(fmt.Sprintf("invoice_%s.pdf", data.InvoiceNumber)), // Set attachment name
			mail.SetHeader(map[string][]string{
				"Content-Type": {"application/pdf"},
			}),
		)
	} else {
		return fmt.Errorf("PDF file not found at path: %s", pdfPath)
	}

	log.Printf("Sending invoice email to %s from %s", data.To, os.Getenv("NOREPLY_MAIL_ADDRESS"))

	d := mail.NewDialer("smtp.gmail.com",
		465,
		os.Getenv("NOREPLY_MAIL_ADDRESS"),
		os.Getenv("NOREPLY_MAIL_PASSWD"),
	)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	// Optionally clean up the PDF file after sending
	if err := os.Remove(pdfPath); err != nil {
		log.Printf("Warning: Could not delete temporary PDF file: %v", err)
	}

	return nil
}

func GetNextInvoiceNumber(year int, month time.Month) string {
	// Get the last invoice number for the given year and month
	var lastInvoice models.Invoice
	if err := db.DB.Where("EXTRACT(YEAR FROM created_at) = ? AND EXTRACT(MONTH FROM created_at) = ?", year, month).Last(&lastInvoice).Error; err != nil {
		log.Printf("Warning: Could not get last invoice number: %v", err)
	}

	// Extract number from F/YYYY/MM/Number
	var lastInvoiceNumber string
	if lastInvoice.InvoiceNumber != "" {
		lastInvoiceNumber = strings.Split(lastInvoice.InvoiceNumber, "/")[3]
	} else {
		lastInvoiceNumber = "0"
	}

	prevInvoiceNumber, err := strconv.Atoi(lastInvoiceNumber)
	if err != nil {
		log.Printf("Warning: Could not convert last invoice number to integer: %v", err)
	}

	// Increment the last invoice number
	nextInvoiceNumber := fmt.Sprintf("F/%d/%d/%04d", year, month, 1+prevInvoiceNumber)

	return nextInvoiceNumber
}
