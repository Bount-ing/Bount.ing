package services

import (
	"open-bounties-api/models"

	"github.com/gin-gonic/gin"
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
func (s *RepositoryService) FetchRepositoryById(id uint) (*models.Repository, error) {
	var repository models.Repository
	result := s.db.First(&repository, id)
	return &repository, result.Error
}

// CreateRepository creates a new repository in the database
func (s *RepositoryService) CreateRepository(c *gin.Context, repository models.Repository) (*models.Repository, error) {
	result := s.db.Create(&repository)
	return &repository, result.Error
}

func (s *RepositoryService) UpdateRepository(id uint, updatedData models.Repository) (*models.Repository, error) {
	var repository models.Repository
	result := s.db.First(&repository, id)
	if result.Error != nil {
		return nil, result.Error
	}

	saveResult := s.db.Save(&repository)
	if saveResult.Error != nil {
		return nil, saveResult.Error
	}
	return &repository, nil
}

func (s *RepositoryService) DeleteRepository(id uint) error {
	var repository models.Repository
	result := s.db.First(&repository, id)
	if result.Error != nil {
		return result.Error
	}
	deleteResult := s.db.Delete(&repository)
	return deleteResult.Error
}
