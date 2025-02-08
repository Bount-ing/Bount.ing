package controllers

import (
	"errors"
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
		return fmt.Errorf("failed to find claimer %d: %v", claimerID, err)
	}

	// Get issue and host details
	issue, err := GetIssueByID(issueID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to find issue %d: %v", issue.ID, err)
	}

	// Find all bounties for the issue
	var bounties []models.Bounty
	if err := tx.Where("issue_id = ?", issueID).Find(&bounties).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to find bounties for issue %d: %v", issueID, err)
	}

	// Save the provided claim check
	claimCheck.CheckerID = claimerID
	claimCheck.CheckerType = "CLAIMER"
	if err := tx.Create(&claimCheck).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create claim check: %v", err)
	}

	// Create a claim for each bounty
	for _, bounty := range bounties {
		ownerClaimCheck := models.ClaimCheck{
			CheckerID:   bounty.OwnerID, // Assuming Bounty has an OwnerID field
			CheckerType: "OWNER",
		}
		if err := tx.Create(&ownerClaimCheck).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to create owner claim check: %v", err)
		}
		systemClaimCheck := models.ClaimCheck{
			CheckerID:   0,
			CheckerType: "SYSTEM",
		}
		if err := tx.Create(&systemClaimCheck).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to create system claim check: %v", err)
		}

		// Calculate the claimed amount
		claimedAmount, err := GetCurrentBountyAmount(bounty.ID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to calculate claimed amount for bounty %d: %v", bounty.ID, err)
		}

		claim := &models.Claim{
			ClaimerID:            claimerID,
			ClaimedAmount:        claimedAmount,
			BountyID:             bounty.ID,
			IssueID:              issueID,
			PullRequestURL:       pullRequestURL,
			ClaimDetails:         claimDetails,
			Status:               "pending",
			BountyClaimerCheckID: claimCheck.ID,
			BountyOwnerCheckID:   ownerClaimCheck.ID,
			BountySystemCheckID:  systemClaimCheck.ID,
		}

		if err := tx.Create(claim).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to create claim for bounty %d: %v", bounty.ID, err)
		}

		//update claimed amount
		bounty.ClaimedAmount = claimedAmount
		if err := tx.Save(&bounty).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update bounty %d: %v", bounty.ID, err)
		}
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	// Send email (non-blocking)
	go func() {
		claimedAt := time.Now().Format("2006-01-02 15:04:05")

		var bountyDetails string
		for _, bounty := range bounties {
			claimedAmount, err := GetCurrentBountyAmount(bounty.ID)
			if err != nil {
				log.Printf("Failed to calculate claimed amount for bounty %d: %v", bounty.ID, err)
				continue
			}
			bountyDetails += fmt.Sprintf("<li><strong>Bounty ID:</strong> %d - <strong>Claimed Amount:</strong> %.2f</li>", bounty.ID, claimedAmount)
		}

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

func ApproveClaim(userID, bountyOwnerID uint, claimID uint, ownerCheck models.ClaimCheck) error {
	log.Printf("Approving claim %d by bounty owner %d", claimID, bountyOwnerID)
	log.Printf("Owner check: %+v", ownerCheck)
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
		return fmt.Errorf("claim not found: %v", err)
	}

	// Verify the bounty owner is associated with the claim
	bounty, err := GetBountyByID(claim.BountyID)
	if err != nil || bounty.OwnerID != bountyOwnerID || bounty.OwnerID != userID {
		tx.Rollback()
		return errors.New("unauthorized: not bounty owner")
	}

	// Fetch the original claimer check
	var claimerCheck models.ClaimCheck
	if err := tx.First(&claimerCheck, claim.BountyClaimerCheckID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("claimer check not found: %v", err)
	}

	// Validate owner check against claimer check
	if !validateOwnerCheck(&claimerCheck, &ownerCheck) {
		tx.Rollback()
		return errors.New("owner check validation failed")
	}

	//retrive preexisting owner check
	var preexistingOwnerCheck models.ClaimCheck
	if err := tx.First(&preexistingOwnerCheck, claim.BountyOwnerCheckID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("preexisting owner check not found: %v", err)
	}

	// Set owner check details
	ownerCheck.CheckerID = bountyOwnerID
	ownerCheck.CheckerType = "OWNER"

	// Save owner check
	if err := tx.Save(&ownerCheck).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to save owner check: %v", err)
	}

	// Update claim with owner check and status
	claim.BountyOwnerCheckID = ownerCheck.ID
	claim.Status = "approved"
	if err := tx.Save(&claim).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update claim: %v", err)
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
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

	// Send email (non-blocking)
	go func() {
		mailContent := fmt.Sprintf(
			`<p>Hi there,</p>
			<p>Your claim for bounty %d has been approved by the bounty owner.</p>
			<p>Details:</p>
			<ul>
				<li>Issue: %d</li>
				<li>Pull Request: %s</li>
			</ul>
			<p>Thank you for your contribution!</p>`,
			bounty.ID, claim.IssueID, claim.PullRequestURL,
		)

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
		}
	}()

	return nil
}

func validateOwnerCheck(claimerCheck, ownerCheck *models.ClaimCheck) bool {
	// Implement validation logic
	// Compare relevant fields between claimer and owner checks
	return ownerCheck.PRNumber == claimerCheck.PRNumber &&
		ownerCheck.RepoName == claimerCheck.RepoName &&
		ownerCheck.RepoOwner == claimerCheck.RepoOwner
}

func GetClaimByCaimCheck(claimerCheckID uint) (models.Claim, error) {
	var claim models.Claim
	err := db.DB.Where("bounty_claimer_check_id = ?", claimerCheckID).First(&claim)
	if err.Error != nil {
		return claim, err.Error
	}
	return claim, nil
}
