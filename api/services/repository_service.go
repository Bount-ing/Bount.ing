package services

import (
    "open-bounties-api/models"
    "gorm.io/gorm"
)

type RepositoryService struct {
    db *gorm.DB
}

func NewRepositoryService(db *gorm.DB) *RepositoryService {
    return &RepositoryService{db: db}
}

// FetchAllRepositories returns all repositories from the database
func (s *RepositoryService) FetchAllRepositories() ([]models.Repository, error) {
    var repositories []models.Repository
    result := s.db.Find(&repositories)
    return repositories, result.Error
}

// FetchRepositoryByID retrieves an repository by its ID from the database
func (s *RepositoryService) FetchRepositoryByID(id uint) (*models.Repository, error) {
    var repository models.Repository
    result := s.db.First(&repository, id)
    return &repository, result.Error
}

// CreateRepository creates a new repository in the database
func (s *RepositoryService) CreateRepository(repository models.Repository) (*models.Repository, error) {
    result := s.db.Create(&repository)
    return &repository, result.Error
}

