package handlers

import (
	"log"
	"net/http"

	"github.com/bount-ing/bount.ing/api/auth"
	"github.com/bount-ing/bount.ing/api/controllers"
	"github.com/bount-ing/bount.ing/api/models"
	"github.com/gin-gonic/gin"
)

type CreateClaimRequest struct {
	ClaimerID      uint                  `json:"claimerId"`
	IssueID        uint                  `json:"IssueId"`
	PullRequestURL string                `json:"prUrl"`
	ClaimDetails   string                `json:"claimDetails"`
	PRVerification models.PRVerification `json:"prVerification"`
}

func CreateClaim(ctx *gin.Context) {
	var req CreateClaimRequest

	u, err := auth.GetUserFromJwt(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": err})
		return
	}
	log.Print("Creating claim for user: ", u.ID)

	// Bind JSON payload to request struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Create PRVerification first
	if err := controllers.CreatePRVerification(&req.PRVerification); err != nil {
		log.Printf("Error creating PRVerification: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create PRVerification"})
		return
	}

	// Create claims for all bounties associated with the issue
	if err := controllers.CreateClaim(
		u.ID,
		req.IssueID,
		req.PullRequestURL,
		req.ClaimDetails,
		req.PRVerification.ID,
	); err != nil {
		log.Printf("Error creating claims: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create claims"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Claims created successfully for all bounties",
		"issueId": req.IssueID,
	})
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

func DeleteClaim(ctx *gin.Context) {
	claimID := ctx.Param("id")

	if err := controllers.DeleteClaim(claimID); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Claim not found"})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
