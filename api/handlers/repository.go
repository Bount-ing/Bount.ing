package handlers

import (
	"log"
	"net/http"

	"github.com/bount-ing/bount.ing/api/controllers"
	"github.com/bount-ing/bount.ing/api/models"
	"github.com/gin-gonic/gin"
)

func CreateRepository(ctx *gin.Context) {
	var repository models.Repository

	log.Print("Creating repository")

	if err := ctx.ShouldBindJSON(&repository); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if err := controllers.CreateRepository(repository); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	ctx.JSON(http.StatusCreated, repository)
}

func GetRepository(ctx *gin.Context) {
	repositoryID := ctx.Param("id")

	repository, err := controllers.GetRepository(repositoryID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Repository not found"})
		return
	}

	ctx.JSON(http.StatusOK, repository)
}

func GetRepositories(ctx *gin.Context) {
	repositories, err := controllers.GetRepositories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	ctx.JSON(http.StatusOK, repositories)
}

func UpdateRepository(ctx *gin.Context) {
	var repository models.Repository

	if err := ctx.ShouldBindJSON(&repository); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if err := controllers.UpdateRepository(repository); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	ctx.JSON(http.StatusOK, repository)
}

func DeleteRepository(ctx *gin.Context) {
	repositoryID := ctx.Param("id")

	if err := controllers.DeleteRepository(repositoryID); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Repository not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Repository deleted"})
}
