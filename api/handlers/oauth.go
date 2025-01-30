package handlers

import (
	"encoding/base64"
	"log"
	"net/http"

	"github.com/bount-ing/bount.ing/api/auth"
	"github.com/bount-ing/bount.ing/api/controllers"
	"github.com/gin-gonic/gin"
)

func GetOAuthState(c *gin.Context) {
	//Get host from URL if not provided return error
	host := c.Param("host")
	if host == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Host not provided"})
		return
	}

	//decode host b64 encoded
	hostDecoded, err := base64.StdEncoding.DecodeString(host)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode host"})
		return
	}

	user, err := auth.GetUserFromJwt(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": err})
		return
	}

	host = string(hostDecoded)
	log.Printf("Generating OAuth state for user %d and host %s", user.ID, host)
	state, err := controllers.GenerateOAuthState(user.ID, host)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"state": state})
}
