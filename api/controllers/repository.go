package controllers

import (
	"github.com/bount-ing/bount.ing/api/db"
	"github.com/bount-ing/bount.ing/api/models"
)

func CreateRepository(repository models.Repository) error {
	err := db.DB.Create(&repository)
	if err != nil {
		return err.Error
	}
	return nil
}

func GetRepository(repositoryID string) (models.Repository, error) {
	var repository models.Repository
	err := db.DB.First(&repository, repositoryID)
	if err != nil {
		return repository, err.Error
	}
	return repository, nil
}

func GetRepositories() ([]models.Repository, error) {
	var repositories []models.Repository
	err := db.DB.Find(&repositories)
	if err != nil {
		return repositories, err.Error
	}
	return repositories, nil
}

func UpdateRepository(repository models.Repository) error {
	err := db.DB.Save(&repository)
	if err != nil {
		return err.Error
	}
	return nil
}

func DeleteRepository(repositoryID string) error {
	var repository models.Repository
	err := db.DB.First(&repository, repositoryID)
	if err != nil {
		return err.Error
	}
	return nil
}
