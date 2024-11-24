package handlers

import (
	"log"
	"net/http"

	"github.com/bount-ing/bount.ing/api/controllers"
	"github.com/bount-ing/bount.ing/api/models"
	"github.com/gin-gonic/gin"
)

func CreateClaim(ctx *gin.Context) {
	var claim models.Claim

	log.Print("Creating claim")

	if err := ctx.ShouldBindJSON(&claim); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if err := controllers.CreateClaim(claim); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	ctx.JSON(http.StatusCreated, claim)
}

func GetClaim(ctx *gin.Context) {
	claimID := ctx.Param("id")

	claim, err := controllers.GetClaim(claimID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Claim not found"})
		return
	}

	ctx.JSON(http.StatusOK, claim)
}

func GetClaims(ctx *gin.Context) {
	claims, err := controllers.GetClaims()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	ctx.JSON(http.StatusOK, claims)
}

func UpdateClaim(ctx *gin.Context) {
	var claim models.Claim

	if err := ctx.ShouldBindJSON(&claim); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if err := controllers.UpdateClaim(claim); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Claim not found"})
		return
	}

}
