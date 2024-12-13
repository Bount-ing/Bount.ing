package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/bount-ing/bount.ing/api/auth"
	"github.com/bount-ing/bount.ing/api/controllers"
	"github.com/bount-ing/bount.ing/api/models"
	"github.com/gin-gonic/gin"
)

func CreateBounty(c *gin.Context) {
	var bounty models.Bounty

	user, err := auth.GetUserFromJwt(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": err})
		return
	}

	bounty.OwnerID = user.ID
	// Manually bind the JSON to the Bounty model
	if err := bounty.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the controller to save the bounty
	if err := controllers.CreateBounty(&bounty); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with success
	c.JSON(http.StatusCreated, gin.H{"message": "Bounty created successfully"})
}

func GetPublicBountiesByIssue(ctx *gin.Context) {
	issues, err := controllers.GetPublicBountiesByIssue()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Bounty not found"})
		return
	}

	log.Printf("Issues: %+v", issues)

	ctx.JSON(http.StatusOK, issues)
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

	user, err := auth.GetUserFromJwt(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": err})
		return
	}

	//convert string to uint
	bountyIDUint, err := strconv.ParseUint(bountyID, 10, 64)

	if err := controllers.DeleteBounty(uint(bountyIDUint), user.ID); err != nil {
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
