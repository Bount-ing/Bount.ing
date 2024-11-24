package handlers

import (
	"log"
	"net/http"

	"github.com/bount-ing/bount.ing/api/controllers"
	"github.com/bount-ing/bount.ing/api/models"
	"github.com/gin-gonic/gin"
)

func CreateBounty(ctx *gin.Context) {
	var bounty models.Bounty

	log.Print("Creating bounty")

	if err := ctx.ShouldBindJSON(&bounty); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if err := controllers.CreateBounty(bounty); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	ctx.JSON(http.StatusCreated, bounty)
}

func GetBounty(ctx *gin.Context) {
	bountyID := ctx.Param("id")

	bounty, err := controllers.GetBounty(bountyID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Bounty not found"})
		return
	}

	ctx.JSON(http.StatusOK, bounty)
}

func GetBounties(ctx *gin.Context) {
	bounties, err := controllers.GetBounties()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	ctx.JSON(http.StatusOK, bounties)
}

func UpdateBounty(ctx *gin.Context) {
	var bounty models.Bounty

	if err := ctx.ShouldBindJSON(&bounty); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if err := controllers.UpdateBounty(bounty); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Bounty not found"})
		return
	}

	ctx.JSON(http.StatusOK, bounty)
}

func DeleteBounty(ctx *gin.Context) {
	bountyID := ctx.Param("id")

	if err := controllers.DeleteBounty(bountyID); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Bounty not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Bounty deleted"})
}

func GetAllUnconfirmedBounties(ctx *gin.Context) {
	bounties, err := controllers.GetAllUnconfirmedBounties()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	ctx.JSON(http.StatusOK, bounties)
}

func FinalizeBounty(ctx *gin.Context) {
	bountyID := ctx.Param("id")

	if err := controllers.FinalizeBounty(bountyID); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Bounty not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Bounty finalized"})
}
