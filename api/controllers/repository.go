package controllers

import (
    "github.com/gin-gonic/gin"
    "open-bounties-api/services"
    "net/http"
)

type RepositoryController struct {
    repositoryService *services.RepositoryService
}

func NewRepositoryController(repositoryService *services.RepositoryService) *RepositoryController {
    return &RepositoryController{
        repositoryService: repositoryService,
    }
}

func (ctl *RepositoryController) GetAllRepositories(c *gin.Context) {
    repositories, err := ctl.repositoryService.FetchAllRepositories()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, repositories)
}

