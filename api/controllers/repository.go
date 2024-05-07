package controllers

import (
	"net/http"
	"open-bounties-api/models"
	"open-bounties-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RepositoryController struct {
	repositoryService *services.RepositoryService
}

func NewRepositoryController(repositoryService *services.RepositoryService) *RepositoryController {
	return &RepositoryController{
		repositoryService: repositoryService,
	}
}

func (uc *RepositoryController) CreateRepository(c *gin.Context) {
	var newRepository models.Repository
	if err := c.ShouldBindJSON(&newRepository); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	registeredRepository, err := uc.repositoryService.CreateRepository(newRepository)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create repository", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, registeredRepository)
}

func (ctl *RepositoryController) GetAllRepositories(c *gin.Context) {
	repositories, err := ctl.repositoryService.FetchAllRepositories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, repositories)
}

func (uc *RepositoryController) GetRepository(c *gin.Context) {
	repositoryIdStr := c.Param("id")
	repositoryId, _ := strconv.ParseUint(repositoryIdStr, 10, 64) // Convert to uint64

	repository, err := uc.repositoryService.FetchRepositoryById(uint(repositoryId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Repository not found", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, repository)
}

func (uc *RepositoryController) UpdateRepository(c *gin.Context) {
	repositoryIdStr := c.Param("id")
	repositoryId, _ := strconv.ParseUint(repositoryIdStr, 10, 64) // Convert to uint64
	var updateRepository models.Repository
	if err := c.ShouldBindJSON(&updateRepository); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	updatedRepository, err := uc.repositoryService.UpdateRepository(uint(repositoryId), updateRepository)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update repository", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedRepository)
}

func (uc *RepositoryController) DeleteRepository(c *gin.Context) {
	repositoryIdStr := c.Param("id")
	repositoryId, _ := strconv.ParseUint(repositoryIdStr, 10, 64) // Convert to uint64
	err := uc.repositoryService.DeleteRepository(uint(repositoryId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete repository", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Repository deleted successfully"})
}
