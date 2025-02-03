package handlers

import (
	"net/http"
	"os"

	"github.com/bount-ing/bount.ing/api/auth"
	"github.com/bount-ing/bount.ing/api/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentmethod"
	"github.com/stripe/stripe-go/setupintent"
	"gorm.io/gorm"
)

// CreateSetupIntent creates a Stripe setup intent
func CreateSetupIntent(c *gin.Context) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	intent, err := setupintent.New(&stripe.SetupIntentParams{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"clientSecret": intent.ClientSecret})
}

// ConfirmSetup stores the payment method after setup
func ConfirmSetup(c *gin.Context) {
	var req struct {
		SetupIntentID string `json:"setupIntentId"`
		UserID        uint   `json:"userId"`
	}

	user, err := auth.GetUserFromJwt(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": err})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	// Retrieve the SetupIntent
	intent, err := setupintent.Get(req.SetupIntentID, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve setup intent"})
		return
	}

	// Get the attached payment method
	pm, err := paymentmethod.Get(intent.PaymentMethod.ID, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve payment method"})
		return
	}

	err = controllers.SavePaymentMethod(user.ID, pm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save payment method"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment method saved"})
}

// GetPaymentMethods retrieves all saved payment methods for a user
func GetPaymentMethods(c *gin.Context) {

	user, err := auth.GetUserFromJwt(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": err})
		return
	}

	paymentMethods, err := controllers.GetPaymentMethods(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve payment methods"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"paymentMethods": paymentMethods})
}

// RemovePaymentMethod deletes a saved payment method
func RemovePaymentMethod(c *gin.Context) {

	user, err := auth.GetUserFromJwt(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": err})
		return
	}

	paymentMethodID := c.Param("id")
	err = controllers.RemovePaymentMethod(user.ID, paymentMethodID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Payment method not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove payment method"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment method removed successfully"})
}
