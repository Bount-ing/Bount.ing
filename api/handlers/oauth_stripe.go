package handlers

import (
	"net/http"

	"github.com/bount-ing/bount.ing/api/controllers"
	"github.com/gin-gonic/gin"
)

func OAuthStripeCallback(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Received callback",
	})
}

func ConnectStripe(c *gin.Context) {
	stripeId := c.Query("id")
	userId, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unautheticated user"})
		return
	}
	if id, ok := userId.(uint); ok {
		controllers.UpdateUserStripeID(id, stripeId)
		c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user", "details": "Unexpected userID type"})
}
