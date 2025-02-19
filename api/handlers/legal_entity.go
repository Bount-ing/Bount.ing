package handlers

import (
	"net/http"
	"time"

	"github.com/bount-ing/bount.ing/api/auth"
	"github.com/bount-ing/bount.ing/api/controllers"
	"github.com/bount-ing/bount.ing/api/models"
	"github.com/gin-gonic/gin"
)

func GetCurrentLegalEntityInfo(c *gin.Context) {
	user, err := auth.GetUserFromJwt(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": err})
		return
	}

	legalInfo, err := controllers.GetLegalEntity(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, legalInfo)
}

func UpdateCurrentLegalEntityInfo(c *gin.Context) {
	user, err := auth.GetUserFromJwt(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": err})
		return
	}

	var legalInfo models.LegalEntity
	if err := c.ShouldBindJSON(&legalInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check date is not before today (remove hours, minutes, seconds)
	today := time.Now().Truncate(24 * time.Hour)
	if legalInfo.ConfirmedAt.Before(today) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ConfirmedAt date cannot be before today"})
		return
	}

	if err := controllers.UpdateLegalEntity(user.ID, legalInfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Legal info updated"})
}
