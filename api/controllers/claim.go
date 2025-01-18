package controllers

import (
	"fmt"
	"log"

	"github.com/bount-ing/bount.ing/api/db"
	"github.com/bount-ing/bount.ing/api/models"
	"github.com/bount-ing/bount.ing/api/tools"
)

func CreateClaim(claimerID uint, issueID uint, pullRequestURL string, claimDetails string) error {
	// Start transaction
	tx := db.DB.Begin()

	//find claimer user
	claimer, err := GetUserByID(claimerID)
	if err != nil {
		tx.Rollback()
		log.Printf("Failed to find claimer %d: %v", claimerID, err)
		return err
	}

	// Find all bounties for the issue
	var bounties []models.Bounty
	if err := tx.Where("issue_id = ?", issueID).Find(&bounties).Error; err != nil {
		tx.Rollback()
		log.Printf("Failed to find bounties for issue %d: %v", issueID, err)
		return err
	}

	// Get Host ID from Issue
	issue, err := GetIssueByID(issueID)
	if err != nil {
		tx.Rollback()
		log.Printf("Failed to find issue %d: %v", issueID, err)
	}

	//get host id from issue
	hostID := issue.HostID

	//Get user Ext identities and get the one matching the host id
	extUser, err := GetExternalIdentityByUserIDAndHostID(claimerID, hostID)
	if err != nil {
		tx.Rollback()
		log.Printf("Failed to find external identity for user %d and host %d: %v", claimerID, hostID, err)
	}

	// Create a claim for each bounty
	for _, bounty := range bounties {
		// First create the claimer's check
		claimerCheck := &models.ClaimCheck{
			Found:            false, // Initial state
			Linked:           false,
			Accepted:         false,
			Closed:           false,
			AuthorUsername:   extUser.Username,
			AuthorExternalID: extUser.ID,
			RepoOwner:        "", // These will be populated when PR is verified
			RepoName:         "",
			PRNumber:         "",
			CheckerID:        claimerID,
			CheckerType:      "CLAIMER",
		}

		if err := tx.Create(claimerCheck).Error; err != nil {
			tx.Rollback()
			log.Printf("Failed to create claimer check: %v", err)
			return err
		}

		// Create the claim with the associated check
		claim := &models.Claim{
			ClaimerID:            claimerID,
			BountyID:             bounty.ID,
			IssueID:              issueID,
			PullRequestURL:       pullRequestURL,
			ClaimDetails:         claimDetails,
			Status:               "pending",
			BountyClaimerCheckID: claimerCheck.ID, // Link to the newly created check
			// Note: BountyOwnerCheckID and BountySystemCheckID will be set later
			// when those checks are performed
		}

		if err := tx.Create(claim).Error; err != nil {
			tx.Rollback()
			log.Printf("Failed to create claim for bounty %d: %v", bounty.ID, err)
			return err
		}

		// Update the claim reference in the check
		if err := tx.Model(claimerCheck).Update("claim_id", claim.ID).Error; err != nil {
			tx.Rollback()
			log.Printf("Failed to update claim reference in check: %v", err)
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return err
	}

	mailContent := fmt.Sprintf(
		`<p>Hi there,</p>
		<p>You have successfully claimed the bounty for issue %d. The claim is currently pending review.</p>
		<p>Details:</p>
		<ul>
			<li>Issue: %d</li>
			<li>Pull Request: %s</li>
		</ul>
		<p>Thank you for your contribution!</p>`,
		issueID, issueID, pullRequestURL,
	)

	// Send email to claimer
	if err := tools.SendEmail("Bount.ing - Bounty Claimed", mailContent, claimer.Email); err != nil {
		log.Printf("Failed to send email to claimer %s: %v", claimer.Email, err)
	}

	return nil
}

func CreateClaimCheck(verification *models.ClaimCheck) error {
	if err := db.DB.Create(verification).Error; err != nil {
		log.Printf("Failed to create Claim Check: %v", err)
		return err
	}
	return nil
}

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
