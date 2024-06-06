package services

import (
	"open-bounties-api/models"

	"gorm.io/gorm"
)

type IssueService struct {
	db *gorm.DB
}

func NewIssueService(db *gorm.DB) *IssueService {
	return &IssueService{db: db}
}

// FetchAllIssues returns all issues from the database
func (s *IssueService) FetchAllIssues() ([]models.Issue, error) {
	var issues []models.Issue
	result := s.db.Find(&issues)
	return issues, result.Error
}

// FetchIssueByID retrieves an issue by its ID from the database
func (s *IssueService) FetchIssueById(id uint) (*models.Issue, error) {
	var issue models.Issue
	result := s.db.First(&issue, id)
	return &issue, result.Error
}

func (s *IssueService) CreateIssue(issue models.Issue) (*models.Issue, error) {
	result := s.db.Create(&issue)
	return &issue, result.Error
}

func (s *IssueService) UpdateIssue(id uint, updatedData models.Issue) (*models.Issue, error) {
	var issue models.Issue
	result := s.db.First(&issue, id)
	if result.Error != nil {
		return nil, result.Error
	}

	saveResult := s.db.Save(&issue)
	if saveResult.Error != nil {
		return nil, saveResult.Error
	}
	return &issue, nil
}

func (s *IssueService) DeleteIssue(id uint) error {
	var issue models.Issue
	result := s.db.First(&issue, id)
	if result.Error != nil {
		return result.Error
	}
	deleteResult := s.db.Delete(&issue)
	return deleteResult.Error
}
