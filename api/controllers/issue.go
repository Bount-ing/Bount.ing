package controllers

import (
	"log"

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
func GetMyIssues(userID uint) ([]models.Issue, error) {
	// Fetch all issues with associated bounties and their variables
	var issues []models.Issue

	err := db.DB.Preload("Bounties").Preload("Bounties.Variables").Find(&issues)

	if err.Error != nil {
		log.Print(err.Error)
		return issues, err.Error
	}

	//for each issue check if the owner id is equal of bounty owner id
	var myIssues []models.Issue
	for _, issue := range issues {
		myBounties := []models.Bounty{}
		for _, bounty := range issue.Bounties {
			if bounty.OwnerID == userID {
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
	err := db.DB.Preload("Bounties").First(&issue, issueID)
	if err.Error != nil {
		log.Print(err.Error)
		return issue, err.Error
	}
	return issue, nil
}
