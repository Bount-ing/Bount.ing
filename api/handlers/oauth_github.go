package handlers

import "github.com/gin-gonic/gin"

func OAuthGithubCallback(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Received callback",
	})
}
