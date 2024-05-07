package services

import (
	"open-bounties-api/models"

	"gorm.io/gorm"
)

type BountyService struct {
	db *gorm.DB
}

func NewBountyService(db *gorm.DB) *BountyService {
	return &BountyService{db: db}
}

// FetchAllBounties returns all bounties from the database
func (s *BountyService) FetchAllBounties() ([]models.Bounty, error) {
	var bounties []models.Bounty
	result := s.db.Find(&bounties)
	return bounties, result.Error
}

// FetchBountyByID retrieves an bounty by its ID from the database
func (s *BountyService) FetchBountyById(id uint) (*models.Bounty, error) {
	var bounty models.Bounty
	result := s.db.First(&bounty, id)
	return &bounty, result.Error
}

// CreateBounty creates a new bounty in the database
func (s *BountyService) CreateBounty(bounty models.Bounty) (*models.Bounty, error) {
	result := s.db.Create(&bounty)
	return &bounty, result.Error
}

func (s *BountyService) UpdateBounty(id uint, updatedData models.Bounty) (*models.Bounty, error) {
	var bounty models.Bounty
	result := s.db.First(&bounty, id)
	if result.Error != nil {
		return nil, result.Error
	}

	saveResult := s.db.Save(&bounty)
	if saveResult.Error != nil {
		return nil, saveResult.Error
	}
	return &bounty, nil
}

func (s *BountyService) DeleteBounty(id uint) error {
	var bounty models.Bounty
	result := s.db.First(&bounty, id)
	if result.Error != nil {
		return result.Error
	}
	deleteResult := s.db.Delete(&bounty)
	return deleteResult.Error
}
