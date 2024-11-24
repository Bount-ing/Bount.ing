package handlers

import (
	"log"
	"net/http"

	"github.com/bount-ing/bount.ing/api/controllers"
	"github.com/bount-ing/bount.ing/api/models"
	"github.com/gin-gonic/gin"
)

func CreateHost(ctx *gin.Context) {
	var host models.Host

	log.Print("Creating host")

	if err := ctx.ShouldBindJSON(&host); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if err := controllers.CreateHost(host); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	ctx.JSON(http.StatusCreated, host)
}

func GetHost(ctx *gin.Context) {
	hostID := ctx.Param("id")

	host, err := controllers.GetHost(hostID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Host not found"})
		return
	}

	ctx.JSON(http.StatusOK, host)
}

func GetHosts(ctx *gin.Context) {
	hosts, err := controllers.GetHosts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	ctx.JSON(http.StatusOK, hosts)
}

func UpdateHost(ctx *gin.Context) {
	var host models.Host

	if err := ctx.ShouldBindJSON(&host); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
}

func DeleteHost(ctx *gin.Context) {
	hostID := ctx.Param("id")

	if err := controllers.DeleteHost(hostID); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Host not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Host deleted"})
}
