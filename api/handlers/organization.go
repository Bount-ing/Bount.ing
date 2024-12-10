package handlers

import (
	"log"
	"net/http"

	"github.com/bount-ing/bount.ing/api/controllers"
	"github.com/bount-ing/bount.ing/api/models"
	"github.com/gin-gonic/gin"
)

func CreateOrganization(ctx *gin.Context) {
	var organization models.Organization

	log.Print("Creating organization")

	if err := ctx.ShouldBindJSON(&organization); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if err := controllers.CreateOrganization(organization); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	ctx.JSON(http.StatusCreated, organization)
}

func GetOrganization(ctx *gin.Context) {
	organizationID := ctx.Param("id")

	organization, err := controllers.GetOrganization(organizationID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Organization not found"})
		return
	}

	ctx.JSON(http.StatusOK, organization)
}

func GetOrganizations(ctx *gin.Context) {
	organizations, err := controllers.GetOrganizations()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	ctx.JSON(http.StatusOK, organizations)
}

func UpdateOrganization(ctx *gin.Context) {
	var organization models.Organization

	if err := ctx.ShouldBindJSON(&organization); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if err := controllers.UpdateOrganization(organization); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	ctx.JSON(http.StatusOK, organization)
}

func DeleteOrganization(ctx *gin.Context) {
	organizationID := ctx.Param("id")

	if err := controllers.DeleteOrganization(organizationID); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Organization not found"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
