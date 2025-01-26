package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/bount-ing/bount.ing/api/auth"
	"github.com/bount-ing/bount.ing/api/controllers"
	"github.com/bount-ing/bount.ing/api/models"
	"github.com/gin-gonic/gin"
)

type CreateClaimRequest struct {
	ClaimerID      uint
	IssueID        uint
	PullRequestURL string
	ClaimDetails   string
	ClaimCheck     models.ClaimCheck
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

func ClaimBounty(ctx *gin.Context) {
	u, err := auth.GetUserFromJwt(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": err})
		log.Print("Unauthorized user: ", u.ID)
		return
	}

	var request struct {
		IssueID        uint              `json:"issueId"`
		BountyID       *uint             `json:"bountyId"`
		PullRequestURL string            `json:"prUrl"`
		ClaimDetails   string            `json:"claimDetails"`
		ClaimCheck     models.ClaimCheck `json:"bountyClaimerCheck"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	// Use request.ClaimCheck.ClaimerID instead of a top-level ClaimerID
	err = controllers.ClaimBounty(
		request.ClaimCheck.CheckerID,
		request.IssueID,
		request.PullRequestURL,
		request.ClaimDetails,
		request.ClaimCheck,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Bounty claimed successfully"})
}

func ApproveClaim(ctx *gin.Context) {
	u, err := auth.GetUserFromJwt(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": err})
		log.Print("Unauthorized user: ", u.ID)
		return
	}

	var request struct {
		Status             string            `json:"status"`
		BountyClaimerCheck models.ClaimCheck `json:"bountyClaimerCheck"`
	}

	claimID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid claim ID"})
		return
	}

	uintClaimID := uint(claimID)

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	// Assuming the ApproveClaim controller method needs these parameters
	err = controllers.ApproveClaim(
		request.BountyClaimerCheck.CheckerID,
		uintClaimID,
		request.BountyClaimerCheck,
	)
	if err != nil {
		if errors.Is(err, errors.New("unauthorized: not bounty owner")) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Claim approved successfully", "status": request.Status})
}
