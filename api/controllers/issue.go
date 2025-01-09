package controllers

import (
	"log"
	"sort"

	"github.com/bount-ing/bount.ing/api/db"
	"github.com/bount-ing/bount.ing/api/models"
)

func CreateIssue(issue models.Issue) (models.Issue, error) {
	err := db.DB.Create(&issue)
	if err.Error != nil {
		log.Print(err.Error)
		return issue, err.Error
	}
	return issue, nil
}

func GetIssue(issueID string) (models.Issue, error) {
	var issue models.Issue
	err := db.DB.First(&issue, issueID)
	if err.Error != nil {
		log.Print(err.Error)
		return issue, err.Error
	}
	return issue, nil
}

func GetIssueByUrl(issueUrl string) (models.Issue, error) {
	var issue models.Issue
	log.Print("Issue URL: ", issueUrl)
	err := db.DB.Where("url = ?", issueUrl).First(&issue)
	if err.Error != nil {
		log.Print(err.Error)
		return issue, err.Error
	}
	return issue, nil
}

func GetIssues() ([]models.Issue, error) {
	var issues []models.Issue
	err := db.DB.Find(&issues)
	if err.Error != nil {
		log.Print(err.Error)
		return issues, err.Error
	}
	return issues, nil
}

func UpdateIssue(issue models.Issue) error {
	err := db.DB.Save(&issue)
	if err.Error != nil {
		log.Print(err.Error)
		return err.Error
	}
	return nil
}

func DeleteIssue(issueID string) error {
	var issue models.Issue
	err := db.DB.First(&issue, issueID)
	if err.Error != nil {
		log.Print(err.Error)
		return err.Error
	}
	return nil
}

// controllers/issue.go

func GetMyIssues(userID uint) ([]models.Issue, error) {
	var issues []models.Issue

	// Preload relationships, including specific PRVerification fields
	err := db.DB.Preload("Bounties").
		Preload("Bounties.Variables").
		Preload("Bounties.Claims", "status = ?", "pending"). // Only pending claims
		Preload("Bounties.Claims.OwnerPRVerification").      // Owner PR verification
		Preload("Bounties.Claims.AuthorPRVerification").     // Author PR verification
		Find(&issues).Error

	if err != nil {
		log.Printf("Error fetching issues for user %d: %v", userID, err)
		return nil, err
	}

	// Filter issues based on user's ownership of bounties
	var myIssues []models.Issue
	for _, issue := range issues {
		var myBounties []models.Bounty
		for _, bounty := range issue.Bounties {
			if bounty.OwnerID == userID {
				// Filter and organize claims directly
				bounty.Claims = filterAndOrganizeClaims(bounty.Claims)
				myBounties = append(myBounties, bounty)
			}
		}

		if len(myBounties) > 0 {
			issue.Bounties = myBounties
			myIssues = append(myIssues, issue)
		}
	}

	return myIssues, nil
}

func GetIssueBounties(issueID string) (models.Issue, error) {
	var issue models.Issue

	err := db.DB.Preload("Bounties").
		Preload("Bounties.Variables").
		Preload("Bounties.Claims", "status = ?", "pending").
		Preload("Bounties.Claims.PRVerification").
		First(&issue, issueID)

	if err.Error != nil {
		log.Print(err.Error)
		return issue, err.Error
	}

	// Organize claims for each bounty
	for i := range issue.Bounties {
		issue.Bounties[i].Claims = filterAndOrganizeClaims(issue.Bounties[i].Claims)
	}

	return issue, nil
}

// Helper function to filter and organize claims
func filterAndOrganizeClaims(claims []models.Claim) []models.Claim {
	// You might want to sort claims by date, filter by status, etc.
	// This is just a basic example
	if claims == nil {
		return []models.Claim{}
	}

	// Sort claims by creation date (newest first)
	sort.Slice(claims, func(i, j int) bool {
		return claims[i].CreatedAt.After(claims[j].CreatedAt)
	})

	return claims
}
