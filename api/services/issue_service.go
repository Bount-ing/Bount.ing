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
func (s *IssueService) FetchIssueByID(id uint) (*models.Issue, error) {
    var issue models.Issue
    result := s.db.First(&issue, id)
    return &issue, result.Error
}

// CreateIssue creates a new issue in the database
func (s *IssueService) CreateIssue(issue models.Issue) (*models.Issue, error) {
    result := s.db.Create(&issue)
    return &issue, result.Error
}

