package controllers

import (
	"net/http"
	"open-bounties-api/models"
	"open-bounties-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ClaimController struct {
	claimService *services.ClaimService
}

func NewClaimController(claimService *services.ClaimService) *ClaimController {
	return &ClaimController{
		claimService: claimService,
	}
}

func (uc *ClaimController) CreateClaim(c *gin.Context) {
	var newClaim models.Claim
	if err := c.ShouldBindJSON(&newClaim); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	registeredClaim, err := uc.claimService.CreateClaim(newClaim)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create claim", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, registeredClaim)
}

func (ctl *ClaimController) GetAllClaims(c *gin.Context) {
	claims, err := ctl.claimService.FetchAllClaims()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, claims)
}

func (uc *ClaimController) GetClaim(c *gin.Context) {
	claimIdStr := c.Param("id")
	claimId, _ := strconv.ParseUint(claimIdStr, 10, 64) // Convert to uint64

	claim, err := uc.claimService.FetchClaimById(uint(claimId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Claim not found", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, claim)
}

func (uc *ClaimController) UpdateClaim(c *gin.Context) {
	claimIdStr := c.Param("id")
	claimId, _ := strconv.ParseUint(claimIdStr, 10, 64) // Convert to uint64
	var updateClaim models.Claim
	if err := c.ShouldBindJSON(&updateClaim); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	updatedClaim, err := uc.claimService.UpdateClaim(uint(claimId), updateClaim)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update claim", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedClaim)
}

func (uc *ClaimController) DeleteClaim(c *gin.Context) {
	claimIdStr := c.Param("id")
	claimId, _ := strconv.ParseUint(claimIdStr, 10, 64) // Convert to uint64
	err := uc.claimService.DeleteClaim(uint(claimId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete claim", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Claim deleted successfully"})
}
