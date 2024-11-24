package handlers

import (
	"github.com/gin-gonic/gin"
)

func IncomingRepositoryWebhook(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Received webhook",
	})
}
