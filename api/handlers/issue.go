package handlers

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/bount-ing/bount.ing/api/auth"
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

	issue, err := controllers.CreateIssue(issue)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
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

func GetIssueByUrl(c *gin.Context) {
	rawURL := c.Param("url")

	// Optionally, decode it yourself
	decodedURL, err := url.QueryUnescape(rawURL)
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to decode URL"})
		return
	}

	trimmedURL := strings.TrimPrefix(decodedURL, "/")

	issue, err := controllers.GetIssueByUrl(trimmedURL)
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": "Issue not found"})
		return
	}

	c.JSON(http.StatusOK, issue)
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

func GetMyIssues(ctx *gin.Context) {
	user, err := auth.GetUserFromJwt(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": err})
		return
	}

	issues, err := controllers.GetMyIssues(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	ctx.JSON(http.StatusOK, issues)
}

func GetIssueBounties(ctx *gin.Context) {
	issueID := ctx.Param("id")

	bounties, err := controllers.GetIssueBounties(issueID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Issue not found"})
		return
	}

	ctx.JSON(http.StatusOK, bounties)
}
