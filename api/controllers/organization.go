package controllers

import (
	"net/http"
	"open-bounties-api/models"
	"open-bounties-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrganizationController struct {
	organizationService *services.OrganizationService
}

func NewOrganizationController(organizationService *services.OrganizationService) *OrganizationController {
	return &OrganizationController{
		organizationService: organizationService,
	}
}

func (uc *OrganizationController) CreateOrganization(c *gin.Context) {
	var newOrganization models.Organization
	if err := c.ShouldBindJSON(&newOrganization); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	registeredOrganization, err := uc.organizationService.CreateOrganization(newOrganization)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create organization", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, registeredOrganization)
}

func (ctl *OrganizationController) GetAllOrganizations(c *gin.Context) {
	organizations, err := ctl.organizationService.FetchAllOrganizations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, organizations)
}

func (uc *OrganizationController) GetOrganization(c *gin.Context) {
	organizationIdStr := c.Param("id")
	organizationId, _ := strconv.ParseUint(organizationIdStr, 10, 64) // Convert to uint64

	organization, err := uc.organizationService.FetchOrganizationById(uint(organizationId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Organization not found", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, organization)
}

func (uc *OrganizationController) UpdateOrganization(c *gin.Context) {
	organizationIdStr := c.Param("id")
	organizationId, _ := strconv.ParseUint(organizationIdStr, 10, 64) // Convert to uint64
	var updateOrganization models.Organization
	if err := c.ShouldBindJSON(&updateOrganization); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	updatedOrganization, err := uc.organizationService.UpdateOrganization(uint(organizationId), updateOrganization)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update organization", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedOrganization)
}

func (uc *OrganizationController) DeleteOrganization(c *gin.Context) {
	organizationIdStr := c.Param("id")
	organizationId, _ := strconv.ParseUint(organizationIdStr, 10, 64) // Convert to uint64
	err := uc.organizationService.DeleteOrganization(uint(organizationId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete organization", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Organization deleted successfully"})
}
