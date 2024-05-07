package controllers

import (
	"net/http"
	"open-bounties-api/models"
	"open-bounties-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BountyController struct {
	bountyService *services.BountyService
}

func NewBountyController(bountyService *services.BountyService) *BountyController {
	return &BountyController{
		bountyService: bountyService,
	}
}

func (uc *BountyController) CreateBounty(c *gin.Context) {
	var newBounty models.Bounty
	if err := c.ShouldBindJSON(&newBounty); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	registeredBounty, err := uc.bountyService.CreateBounty(newBounty)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bounty", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, registeredBounty)
}

func (ctl *BountyController) GetAllBounties(c *gin.Context) {
	bounties, err := ctl.bountyService.FetchAllBounties()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bounties)
}

func (uc *BountyController) GetBounty(c *gin.Context) {
	bountyIdStr := c.Param("id")
	bountyId, _ := strconv.ParseUint(bountyIdStr, 10, 64) // Convert to uint64

	bounty, err := uc.bountyService.FetchBountyById(uint(bountyId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Bounty not found", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bounty)
}

func (uc *BountyController) UpdateBounty(c *gin.Context) {
	bountyIdStr := c.Param("id")
	bountyId, _ := strconv.ParseUint(bountyIdStr, 10, 64) // Convert to uint64
	var updateBounty models.Bounty
	if err := c.ShouldBindJSON(&updateBounty); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	updatedBounty, err := uc.bountyService.UpdateBounty(uint(bountyId), updateBounty)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update bounty", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedBounty)
}

func (uc *BountyController) DeleteBounty(c *gin.Context) {
	bountyIdStr := c.Param("id")
	bountyId, _ := strconv.ParseUint(bountyIdStr, 10, 64) // Convert to uint64
	err := uc.bountyService.DeleteBounty(uint(bountyId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete bounty", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bounty deleted successfully"})
}
