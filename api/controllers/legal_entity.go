package controllers

import (
	"github.com/bount-ing/bount.ing/api/db"
	"github.com/bount-ing/bount.ing/api/models"
	"gorm.io/gorm"
)

func CreateLegalEntity(userLegalData models.LegalEntity) error {
	if err := db.DB.Create(&userLegalData).Error; err != nil {
		return err
	}

	return nil
}

func GetLegalEntity(userID uint) (models.LegalEntity, error) {
	var userLegalData models.LegalEntity

	dbc := db.DB.Where("user_id = ?", userID).First(&userLegalData)
	if dbc.Error == gorm.ErrRecordNotFound {
		return userLegalData, nil
	} else if dbc.Error != nil {
		return userLegalData, dbc.Error
	}
	return userLegalData, dbc.Error
}

func UpdateLegalEntity(userID uint, userLegalData models.LegalEntity) error {
	var userLegalDataInDB models.LegalEntity

	dbc := db.DB.Where("user_id = ?", userID).First(&userLegalDataInDB)
	if dbc.Error == gorm.ErrRecordNotFound {
		//create new legal entity
		userLegalData.UserID = userID
		return CreateLegalEntity(userLegalData)
	} else if dbc.Error != nil {
		return dbc.Error
	}

	userLegalDataInDB.IsCompany = userLegalData.IsCompany
	userLegalDataInDB.DocumentType = userLegalData.DocumentType
	userLegalDataInDB.DocumentCountry = userLegalData.DocumentCountry
	userLegalDataInDB.DocumentNumber = userLegalData.DocumentNumber
	userLegalDataInDB.LegalName = userLegalData.LegalName
	userLegalDataInDB.LegalAddress = userLegalData.LegalAddress
	userLegalDataInDB.LegalCity = userLegalData.LegalCity
	userLegalDataInDB.LegalZip = userLegalData.LegalZip
	userLegalDataInDB.LegalState = userLegalData.LegalState
	userLegalDataInDB.LegalCountry = userLegalData.LegalCountry
	userLegalDataInDB.ConfirmedAt = userLegalData.ConfirmedAt

	saveResult := db.DB.Save(&userLegalDataInDB)
	if saveResult.Error != nil {
		return saveResult.Error
	}
	return nil
}
