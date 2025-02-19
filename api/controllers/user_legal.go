package controllers

import (
	"github.com/bount-ing/bount.ing/api/db"
	"github.com/bount-ing/bount.ing/api/models"
	"gorm.io/gorm"
)

func CreateUserLegalData(userLegalData models.UserLegalData) error {
	if err := db.DB.Create(&userLegalData).Error; err != nil {
		return err
	}

	return nil
}

func GetUserLegalData(userID uint) (models.UserLegalData, error) {
	var userLegalData models.UserLegalData

	dbc := db.DB.Where("user_id = ?", userID).First(&userLegalData)
	if dbc.Error == gorm.ErrRecordNotFound {
		return userLegalData, nil
	} else if dbc.Error != nil {
		return userLegalData, dbc.Error
	}
	return userLegalData, dbc.Error
}

func UpdateUserLegalData(userID uint, userLegalData models.UserLegalData) error {
	var userLegalDataInDB models.UserLegalData

	dbc := db.DB.Where("user_id = ?", userID).First(&userLegalDataInDB)
	if dbc.Error == gorm.ErrRecordNotFound {
		return dbc.Error
	}

	userLegalDataInDB.TaxID = userLegalData.TaxID
	userLegalDataInDB.DocumentType = userLegalData.DocumentType
	userLegalDataInDB.DocumentCountry = userLegalData.DocumentCountry
	userLegalDataInDB.IsCompany = userLegalData.IsCompany
	userLegalDataInDB.CompanyName = userLegalData.CompanyName
	userLegalDataInDB.CompanyAddress = userLegalData.CompanyAddress
	userLegalDataInDB.CompanyCity = userLegalData.CompanyCity
	userLegalDataInDB.CompanyState = userLegalData.CompanyState
	userLegalDataInDB.CompanyZip = userLegalData.CompanyZip
	userLegalDataInDB.CompanyCountry = userLegalData.CompanyCountry

	saveResult := db.DB.Save(&userLegalDataInDB)
	if saveResult.Error != nil {
		return saveResult.Error
	}
	return nil
}
