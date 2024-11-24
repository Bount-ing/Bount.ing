package controllers

import (
	"log"

	"github.com/bount-ing/bount.ing/api/db"
	"github.com/bount-ing/bount.ing/api/models"
)

func CreateIssue(issue models.Issue) error {
	err := db.DB.Create(&issue)
	if err.Error != nil {
		log.Print(err.Error)
		return err.Error
	}
	return nil
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
