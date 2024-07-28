package controllers

import (
	"log"
	"net/http"
	"open-bounties-api/models"
	"open-bounties-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BountyController struct {
	db            *gorm.DB
	bountyService *services.BountyService
}

func NewBountyController(db *gorm.DB, bountyService *services.BountyService) *BountyController {
	return &BountyController{
		db:            db,
		bountyService: bountyService,
	}
}

func (uc *BountyController) CreateBounty(c *gin.Context) {
	var newBounty models.Bounty

	log.Print("Creating bounty")
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	// Convert userID to uint
	userIDFloat64, ok := userID.(float64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID type"})
		return
	}
	userIDUint := uint(userIDFloat64)

	// Ensure the user exists in the database
	var user models.User
	if err := uc.db.First(&user, userIDUint).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error", "details": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&newBounty); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	newBounty.OwnerID = userIDUint
	registeredBounty, err := uc.bountyService.CreateBounty(c, newBounty)
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

func (ctl *BountyController) GetAllUnconfirmedBounties(c *gin.Context) {
	bounties, err := ctl.bountyService.FetchAllUncofirmedBounties()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bounties)
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
	bountyId, _ := strconv.ParseUint(bountyIdStr, 10, 64)

	if err := uc.bountyService.CancelBounty(uint(bountyId)); err != nil {
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bounty canceled successfully"})
}

func (uc *BountyController) FinalizeBounty(c *gin.Context) {
	bountyIdStr := c.Param("id")
	bountyId, _ := strconv.ParseUint(bountyIdStr, 10, 64)

	if err := uc.bountyService.FinalizeBounty(uint(bountyId)); err != nil {
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bounty finalized successfully"})
}
