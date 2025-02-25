package controllers

import (
	"fmt"
	"log"
	"time"

	"github.com/bount-ing/bount.ing/api/db"
	"github.com/bount-ing/bount.ing/api/models"
	"github.com/bount-ing/bount.ing/api/tools"
)

func GetClaim(claimID string) (models.Claim, error) {
	var claim models.Claim

	err := db.DB.First(&claim, claimID)

	if err.Error != nil {
		log.Print(err.Error)
		return claim, err.Error
	}
	return claim, nil
}

func GetClaims() ([]models.Claim, error) {
	var claims []models.Claim

	err := db.DB.Find(&claims)

	if err.Error != nil {
		log.Print(err.Error)
		return claims, err.Error
	}

	return claims, nil
}

func UpdateClaim(claim models.Claim) error {
	err := db.DB.Save(&claim)

	if err.Error != nil {
		log.Print(err.Error)
		return err.Error
	}
	return nil
}

func DeleteClaim(claimID string) error {
	var claim models.Claim
	err := db.DB.First(&claim, claimID)
	if err.Error != nil {
		log.Print(err.Error)
		return err.Error
	}
	err = db.DB.Delete(&claim)
	if err.Error != nil {
		log.Print(err.Error)
		return err.Error
	}
	return nil
}

func ClaimBounty(claimerID uint, issueID uint, pullRequestURL string, claimDetails string, claimCheck models.ClaimCheck) error {
	// Start transaction
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Find claimer user
	claimer, err := GetUserByID(claimerID)
	if err != nil {
		tx.Rollback()
		return models.ErrUserNotFound
	}

	// Check legal data for the claimer
	userLegalData, err := GetLegalEntity(claimer.ID)
	if err != nil {
		log.Printf("Error fetching user legal data: %s", err)
		return models.ErrClaimWithoutLegalEntity
	} else if userLegalData.DocumentNumber == "" {
		log.Printf("User does not have a tax ID")
		return models.ErrClaimWithoutLegalEntity
	}

	// Get issue and host details
	issue, err := GetIssueByID(issueID)
	if err != nil {
		tx.Rollback()
		return models.ErrIssueNotFound
	}

	// Find all bounties for the issue
	var bounties []models.Bounty
	if err := tx.Where("issue_id = ?", issueID).Find(&bounties).Error; err != nil {
		tx.Rollback()
		return models.ErrBountyNotFound
	}

	// Save the provided claim check
	claimCheck.CheckerID = claimerID
	claimCheck.CheckerType = "CLAIMER"
	if err := tx.Create(&claimCheck).Error; err != nil {
		tx.Rollback()
		return models.ErrClaimCheckCreationFailed
	}

	var bountyDetails string

	// Create a claim for each bounty
	for _, bounty := range bounties {
		// Calculate the claimed amount for this bounty
		claimedAmount, err := GetCurrentBountyAmount(bounty.ID)
		if err != nil {
			tx.Rollback()
			return models.ErrBountyAmountCalculationFailed
		}

		// Save the sponsor claim check
		sponsorClaimCheck := models.ClaimCheck{
			CheckerID:   bounty.SponsorID, // Assuming Bounty has an SponsorID field
			CheckerType: "OWNER",
		}
		if err := tx.Create(&sponsorClaimCheck).Error; err != nil {
			tx.Rollback()
			return models.ErrSponsorClaimCheckCreationFailed
		}

		// Save the system claim check
		systemClaimCheck := models.ClaimCheck{
			CheckerID:   0,
			CheckerType: "SYSTEM",
		}
		if err := tx.Create(&systemClaimCheck).Error; err != nil {
			tx.Rollback()
			return models.ErrSystemClaimCheckCreationFailed
		}

		// Create the claim for the bounty
		claim := &models.Claim{
			ClaimerID:            claimerID,
			ClaimedAmount:        claimedAmount,
			BountyID:             bounty.ID,
			IssueID:              issueID,
			PullRequestURL:       pullRequestURL,
			ClaimDetails:         claimDetails,
			Status:               "pending",
			BountyClaimerCheckID: claimCheck.ID,
			BountySponsorCheckID: sponsorClaimCheck.ID,
			BountySystemCheckID:  systemClaimCheck.ID,
		}

		if err := tx.Create(claim).Error; err != nil {
			tx.Rollback()
			return models.ErrClaimCreationFailed
		}

		// Update claimed amount in bounty
		bounty.ClaimedAmount = claimedAmount
		if err := tx.Save(&bounty).Error; err != nil {
			tx.Rollback()
			return models.ErrBountyUpdateFailed
		}

		// Accumulate bounty details for email
		bountyDetails += fmt.Sprintf("<li><strong>Bounty ID:</strong> %d - <strong>Claimed Amount:</strong> %.2f</li>", bounty.ID, claimedAmount)

		// Send email (non-blocking) to bounty sponsor
		go func() {
			mailContent := fmt.Sprintf(
				`<p>Hi there,</p>
				<p>A new claim has been made for bounty %d.</p>
				<p>Details:</p>
				<ul>
					<li>Issue: %d</li>
					<li>Pull Request: %s</li>
				</ul>
				<p> Please review the claim and approve or reject it.</p>
				<p> Dashboard: <a href="https://bount.ing/profile">Review your Bounties</a></p>
				<p>Thank you for your contribution!</p>`,
				bounty.ID, issueID, pullRequestURL,
			)

			sponsor, err := GetUserByID(bounty.SponsorID)
			if err != nil {
				log.Printf("Failed to find bounty sponsor %d: %v", bounty.SponsorID, err)
				return
			}

			if err := tools.SendEmail(sponsor.Email, "Bount.ing - New Bounty Claim", mailContent); err != nil {
				log.Printf("Failed to send email to bounty sponsor %s: %v", sponsor.Email, err)
			}
		}()
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return models.ErrClaimCreationFailed
	}

	// Send email (non-blocking)
	go func() {
		claimedAt := time.Now().Format("2006-01-02 15:04:05")

		mailContent := fmt.Sprintf(
			`<p>Hi %s,</p>
			<p>You have successfully claimed the bounty for <strong><a href="%s">Issue #%d</a></strong>. The claim is currently pending review.</p>
			<hr>
			<h3>Claim Details:</h3>
			<ul>
				<li><strong>Issue:</strong> <a href="%s">%s</a></li>
				<li><strong>Pull Request:</strong> <a href="%s">%s</a></li>
				<li><strong>Claim Date:</strong> %s</li>
			</ul>
			<hr>
			<h3>Bounties Claimed:</h3>
			<ul>
				%s
			</ul>
			<p>Thank you for your contribution!</p>`,
			claimer.Username, issue.URL, issueID, issue.URL, issue.URL, pullRequestURL, pullRequestURL, claimedAt, bountyDetails,
		)

		if err := tools.SendEmail(claimer.Email, "Bount.ing - Bounty Claimed", mailContent); err != nil {
			log.Printf("Failed to send email to claimer %s: %v", claimer.Email, err)
		}
	}()

	return nil
}

func ApproveClaim(userID, sponsorID uint, claimID uint, sponsorCheck models.ClaimCheck, invoice bool) error {
	log.Printf("Approving claim %d by bounty sponsor %d", claimID, sponsorID)
	log.Printf("Sponsor check: %+v", sponsorCheck)

	sponsor, err := GetUserByID(sponsorID)
	if err != nil {
		return models.ErrUserNotFound
	}

	// Check legal Tax ID for the sponsor
	sponsorLegalData, err := GetLegalEntity(sponsor.ID)
	if invoice {
		if err != nil {
			log.Printf("Error fetching user legal data: %s", err)
			return models.ErrInvoiceWithoutLegalEntity
		} else if sponsorLegalData.DocumentNumber == "" {
			log.Printf("User does not have a tax ID")
			return models.ErrInvoiceWithoutLegalEntity
		}
	}

	// Begin transaction
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Fetch the claim
	var claim models.Claim
	if err := tx.First(&claim, claimID).Error; err != nil {
		tx.Rollback()
		return models.ErrClaimNotFound
	}

	// Verify the bounty sponsor is associated with the claim
	bounty, err := GetBountyByID(claim.BountyID)
	if err != nil || bounty.SponsorID != sponsorID || bounty.SponsorID != userID {
		tx.Rollback()
		return models.ErrApproveClaimNotOwned
	}

	// Fetch the original claimer check
	var claimerCheck models.ClaimCheck
	if err := tx.First(&claimerCheck, claim.BountyClaimerCheckID).Error; err != nil {
		tx.Rollback()
		return models.ErrClaimCheckNotFound
	}

	// Validate sponsor check against claimer check
	if !validateSponsorCheck(&claimerCheck, &sponsorCheck) {
		tx.Rollback()
		return models.ErrSponsorRejectedClaimCheck
	}

	//retrive preexisting sponsor check
	var preexistingSponsorCheck models.ClaimCheck
	if err := tx.First(&preexistingSponsorCheck, claim.BountySponsorCheckID).Error; err != nil {
		tx.Rollback()
		return models.ErrClaimCheckNotFound
	}

	// Set sponsor check details
	sponsorCheck.CheckerID = sponsorID
	sponsorCheck.CheckerType = "OWNER"

	// Save sponsor check
	if err := tx.Save(&sponsorCheck).Error; err != nil {
		tx.Rollback()
		return models.ErrClaimCheckCreationFailed
	}

	// Update claim with sponsor check and status
	claim.BountySponsorCheckID = sponsorCheck.ID
	claim.Status = "approved"
	if err := tx.Save(&claim).Error; err != nil {
		tx.Rollback()
		return models.ErrClaimCheckUpdateFailed
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return models.ErrClaimCheckUpdateFailed
	}

	//Retrieve bounty & update status
	bounty, err = GetBountyByID(claim.BountyID)
	if err != nil {
		log.Print(err)
	}
	bounty.Status = "closed"
	if err := db.DB.Save(&bounty).Error; err != nil {
		log.Print(err)
	}

	// Send email (non-blocking) to claimer
	go func() {

		mailContent := ""
		if invoice {
			mailContent = fmt.Sprintf(
				`<p>Hi there,</p>
				<p>Your claim for bounty %d has been approved by the bounty sponsor.</p>
				<p>Please note that the sponsor has requested an invoice for the payment.</p>
				<p>Mission Details:</p>
				<ul>
					<li>Issue: %d</li>
					<li>Pull Request: %s</li>
					<li>Claim Date: %s</li>
					<li>Claim Reward: %.2f</li>
				</ul>
				<p>Sponsor Details:</p>
				<ul>
					<li>Legal Name / Company: %s</li>
					<li>Address: %s</li>
					<li>City: %s</li>
					<li>Zip: %s</li>
					<li>State: %s</li>
					<li>Country: %s</li>
				</ul>
				<ul>
					<li>Document Type: %s</li>
					<li>Document Country: %s</li>
					<li>Document Number: %s</li>
				</ul>
				<p>Thank you for your contribution!</p>`,
				claim.BountyID, claim.IssueID, claim.PullRequestURL, claim.CreatedAt.Format("2006-01-02 15:04:05"), claim.ClaimedAmount,
				sponsorLegalData.LegalName, sponsorLegalData.LegalAddress, sponsorLegalData.LegalCity, sponsorLegalData.LegalZip, sponsorLegalData.LegalState, sponsorLegalData.LegalCountry,
				sponsorLegalData.DocumentType, sponsorLegalData.DocumentCountry, sponsorLegalData.DocumentNumber,
			)
		} else {
			mailContent = fmt.Sprintf(
				`<p>Hi there,</p>
			<p>Your claim for bounty %d has been approved by the bounty sponsor.</p>
			<p>Mission Details:</p>
			<ul>
				<li>Issue: %d</li>
				<li>Pull Request: %s</li>
				<li>Claim Date: %s</li>
				<li>Claim Reward: %.2f</li>
			</ul>
			<p>Thank you for your contribution!</p>`,
				claim.BountyID, claim.IssueID, claim.PullRequestURL, claim.CreatedAt.Format("2006-01-02 15:04:05"), claim.ClaimedAmount,
			)
		}
		//get claimer
		claimer, err := GetUserByID(claim.ClaimerID)
		if err != nil {
			log.Print(err)
			return
		}

		if err := tools.SendEmail(claimer.Email, "Bount.ing - Bounty Claim Approved", mailContent); err != nil {
			log.Printf("Failed to send email to claimer %s: %v", claimer.Email, err)
		}
	}()

	//process paiment
	go func() {
		// Process payment
		if err := ProcessPayment(claimerCheck.CheckerID, bounty.ID); err != nil {
			log.Printf("Failed to process payment for claimer %d: %v", claimerCheck.CheckerID, err)
			return
		}
		claim.Status = "paid"
		if err := db.DB.Save(&claim).Error; err != nil {
			log.Print(err)
		}

		// Generate PDF invoice for the platform fees
		if err := tools.GenerateInvoice(bounty.ID, claim.ID); err != nil {
			log.Printf("Failed to generate invoice for claimer %d: %v", claimerCheck.CheckerID, err)
		}

	}()
	return nil
}

func validateSponsorCheck(claimerCheck, sponsorCheck *models.ClaimCheck) bool {
	// Implement validation logic
	// Compare relevant fields between claimer and sponsor checks
	return sponsorCheck.PRNumber == claimerCheck.PRNumber &&
		sponsorCheck.RepoName == claimerCheck.RepoName &&
		sponsorCheck.RepoSponsor == claimerCheck.RepoSponsor
}

func GetClaimByCaimCheck(claimerCheckID uint) (models.Claim, error) {
	var claim models.Claim
	err := db.DB.Where("bounty_claimer_check_id = ?", claimerCheckID).First(&claim)
	if err.Error != nil {
		return claim, err.Error
	}
	return claim, nil
}
