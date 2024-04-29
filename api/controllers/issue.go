package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "open-bounties-api/models"
    "open-bounties-api/services"
)

func GetAllIssues(c *gin.Context) {
    issues, err := services.FetchAllIssues()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve issues"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"issues": issues})
}

func CreateIssue(c *gin.Context) {
    var newIssue models.Issue
    if err := c.ShouldBindJSON(&newIssue); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid issue data"})
        return
    }

    issue, err := services.CreateIssue(newIssue)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create issue"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"issue": issue})
}

func GetIssue(c *gin.Context) {
    issueID := c.Param("id")
    issue, err := services.FetchIssueByID(issueID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Issue not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"issue": issue})
}

func UpdateIssue(c *gin.Context) {
    issueID := c.Param("id")
    var issueUpdates models.Issue
    if err := c.ShouldBindJSON(&issueUpdates); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data provided"})
        return
    }

    updatedIssue, err := services.UpdateIssue(issueID, issueUpdates)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update issue"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"issue": updatedIssue})
}


func DeleteIssue(c *gin.Context) {
    issueID := c.Param("id")
    err := services.DeleteIssue(issueID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete issue"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Issue deleted"})
}

