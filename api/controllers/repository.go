package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "open-bounties-api/models"
    "open-bounties-api/services"
)

func GetAllRepositories(c *gin.Context) {
    issues, err := services.FetchAllRepositories()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve issues"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"issues": issues})
}

func CreateRepository(c *gin.Context) {
    var newRepository models.Repository
    if err := c.ShouldBindJSON(&newRepository); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid issue data"})
        return
    }

    issue, err := services.CreateRepository(newRepository)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create issue"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"issue": issue})
}

func GetRepository(c *gin.Context) {
    issueID := c.Param("id")
    issue, err := services.FetchRepositoryByID(issueID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Repository not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"issue": issue})
}

func UpdateRepository(c *gin.Context) {
    issueID := c.Param("id")
    var issueUpdates models.Repository
    if err := c.ShouldBindJSON(&issueUpdates); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data provided"})
        return
    }

    updatedRepository, err := services.UpdateRepository(issueID, issueUpdates)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update issue"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"issue": updatedRepository})
}


func DeleteRepository(c *gin.Context) {
    issueID := c.Param("id")
    err := services.DeleteRepository(issueID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete issue"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Repository deleted"})
}

