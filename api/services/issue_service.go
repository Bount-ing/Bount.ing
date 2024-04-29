package services

import (
    "open-bounties-api/models"
    "errors"
    "gorm.io/gorm"
)

type IssueService struct {
    db *gorm.DB // Assuming you are using GORM for database operations
}

func NewIssueService(db *gorm.DB) *IssueService {
    return &IssueService{db: db}
}

func (s *IssueService) GetIssueByID(id uint) (*models.Issue, error) {
    var issue models.Issue
    result := s.db.First(&issue, id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &issue, nil
}

func (s *IssueService) CreateIssue(issue models.Issue) (*models.Issue, error) {
    if issue.Title == "" {
        return nil, errors.New("issue title cannot be empty")
    }

    result := s.db.Create(&issue)
    if result.Error != nil {
        return nil, result.Error
    }

    return &issue, nil
}

