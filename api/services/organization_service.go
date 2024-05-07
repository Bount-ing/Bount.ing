package services

import (
	"open-bounties-api/models"

	"gorm.io/gorm"
)

type OrganizationService struct {
	db *gorm.DB
}

func NewOrganizationService(db *gorm.DB) *OrganizationService {
	return &OrganizationService{db: db}
}

// FetchAllOrganizations returns all organizations from the database
func (s *OrganizationService) FetchAllOrganizations() ([]models.Organization, error) {
	var organizations []models.Organization
	result := s.db.Find(&organizations)
	return organizations, result.Error
}

// FetchOrganizationByID retrieves an organization by its ID from the database
func (s *OrganizationService) FetchOrganizationById(id uint) (*models.Organization, error) {
	var organization models.Organization
	result := s.db.First(&organization, id)
	return &organization, result.Error
}

// CreateOrganization creates a new organization in the database
func (s *OrganizationService) CreateOrganization(organization models.Organization) (*models.Organization, error) {
	result := s.db.Create(&organization)
	return &organization, result.Error
}

func (s *OrganizationService) UpdateOrganization(id uint, updatedData models.Organization) (*models.Organization, error) {
	var organization models.Organization
	result := s.db.First(&organization, id)
	if result.Error != nil {
		return nil, result.Error
	}

	saveResult := s.db.Save(&organization)
	if saveResult.Error != nil {
		return nil, saveResult.Error
	}
	return &organization, nil
}

func (s *OrganizationService) DeleteOrganization(id uint) error {
	var organization models.Organization
	result := s.db.First(&organization, id)
	if result.Error != nil {
		return result.Error
	}
	deleteResult := s.db.Delete(&organization)
	return deleteResult.Error
}
