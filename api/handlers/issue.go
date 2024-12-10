package handlers

import (
	"log"
	"net/http"

	"github.com/bount-ing/bount.ing/api/controllers"
	"github.com/bount-ing/bount.ing/api/models"
	"github.com/gin-gonic/gin"
)

func CreateIssue(ctx *gin.Context) {
	var issue models.Issue

	log.Print("Creating issue")

	if err := ctx.ShouldBindJSON(&issue); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if err := controllers.CreateIssue(issue); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	ctx.JSON(http.StatusCreated, issue)
}

func GetIssue(ctx *gin.Context) {
	issueID := ctx.Param("id")

	issue, err := controllers.GetIssue(issueID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Issue not found"})
		return
	}

	ctx.JSON(http.StatusOK, issue)
}

func GetIssues(ctx *gin.Context) {
	issues, err := controllers.GetIssues()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	ctx.JSON(http.StatusOK, issues)
}

func UpdateIssue(ctx *gin.Context) {
	var issue models.Issue

	if err := ctx.ShouldBindJSON(&issue); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if err := controllers.UpdateIssue(issue); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	ctx.JSON(http.StatusOK, issue)
}

func DeleteIssue(ctx *gin.Context) {
	issueID := ctx.Param("id")
	if err := controllers.DeleteIssue(issueID); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Issue not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Issue deleted"})
}
