package controllers

import (
	"log"

	"github.com/bount-ing/bount.ing/api/db"
	"github.com/bount-ing/bount.ing/api/models"
	"github.com/stripe/stripe-go"
)

func SavePaymentMethod(userID uint, pm *stripe.PaymentMethod) error {

	newPaymentMethod := models.PaymentMethod{
		ID:       pm.ID,
		UserID:   userID,
		Brand:    string(pm.Card.Brand),
		Last4:    pm.Card.Last4,
		ExpMonth: int(pm.Card.ExpMonth),
		ExpYear:  int(pm.Card.ExpYear),
	}

	err := db.DB.Create(&newPaymentMethod)
	if err.Error != nil {
		log.Print(err.Error)
		return err.Error
	}

	return nil
}

func GetPaymentMethods(userID uint) ([]models.PaymentMethod, error) {
	var paymentMethods []models.PaymentMethod

	err := db.DB.Where("user_id = ?", userID).Find(&paymentMethods)
	if err.Error != nil {
		log.Print(err.Error)
		return paymentMethods, err.Error
	}

	return paymentMethods, nil
}

func RemovePaymentMethod(userID uint, paymentMethodID string) error {
	var paymentMethod models.PaymentMethod

	err := db.DB.Where("user_id = ? AND id = ?", userID, paymentMethodID).First(&paymentMethod)
	if err.Error != nil {
		log.Print(err.Error)
		return err.Error
	}

	err = db.DB.Delete(&paymentMethod)
	if err.Error != nil {
		log.Print(err.Error)
		return err.Error
	}

	return nil
}
