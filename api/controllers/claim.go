package controllers

import (
	"log"

	"github.com/bount-ing/bount.ing/api/db"
	"github.com/bount-ing/bount.ing/api/models"
)

func CreateClaim(claimerID uint, issueID uint, pullRequestURL string, claimDetails string, prVerificationID uint) error {
	// Start transaction
	tx := db.DB.Begin()

	// Find all bounties for the issue
	var bounties []models.Bounty
	if err := tx.Where("issue_id = ?", issueID).Find(&bounties).Error; err != nil {
		tx.Rollback()
		log.Printf("Failed to find bounties for issue %d: %v", issueID, err)
		return err
	}

	// Create a claim for each bounty
	for _, bounty := range bounties {
		claim := &models.Claim{
			ClaimerID:              claimerID,
			BountyID:               bounty.ID,
			IssueID:                issueID,
			PullRequestURL:         pullRequestURL,
			ClaimDetails:           claimDetails,
			Status:                 "pending",
			OwnerPRVerificationID:  prVerificationID,
			AuthorPRVerificationID: prVerificationID,
		}

		if err := tx.Create(claim).Error; err != nil {
			tx.Rollback()
			log.Printf("Failed to create claim for bounty %d: %v", bounty.ID, err)
			return err
		}
	}

	return tx.Commit().Error
}

func CreatePRVerification(verification *models.PRVerification) error {
	if err := db.DB.Create(verification).Error; err != nil {
		log.Printf("Failed to create PRVerification: %v", err)
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
