package controllers

import (
	"net/http"
	"open-bounties-api/models"
	"open-bounties-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IssueController struct {
	issueService *services.IssueService
	db           *gorm.DB
}

func NewIssueController(issueService *services.IssueService, db *gorm.DB) *IssueController {
	return &IssueController{
		issueService: issueService,
		db:           db,
	}
}

func (uc *IssueController) CreateIssue(c *gin.Context) {
	var newIssue models.Issue
	if err := c.ShouldBindJSON(&newIssue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	registeredIssue, err := uc.issueService.CreateIssue("", newIssue)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create issue", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, registeredIssue)
}

func (ctl *IssueController) GetAllIssues(c *gin.Context) {
	issues, err := ctl.issueService.FetchAllIssues()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, issues)
}

func (uc *IssueController) GetIssue(c *gin.Context) {
	issueIdStr := c.Param("id")
	issueId, _ := strconv.ParseUint(issueIdStr, 10, 64) // Convert to uint64

	issue, err := uc.issueService.FetchIssueById(uint(issueId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Issue not found", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, issue)
}

func (uc *IssueController) UpdateIssue(c *gin.Context) {
	issueIdStr := c.Param("id")
	issueId, _ := strconv.ParseUint(issueIdStr, 10, 64) // Convert to uint64
	var updateIssue models.Issue
	if err := c.ShouldBindJSON(&updateIssue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	updatedIssue, err := uc.issueService.UpdateIssue(uint(issueId), updateIssue)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update issue", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedIssue)
}

func (uc *IssueController) DeleteIssue(c *gin.Context) {
	issueIdStr := c.Param("id")
	issueId, _ := strconv.ParseUint(issueIdStr, 10, 64) // Convert to uint64
	err := uc.issueService.DeleteIssue(uint(issueId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete issue", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Issue deleted successfully"})
}

func (uc *IssueController) IssueGithubWebhook(c *gin.Context) {
	var payload map[string]interface{}
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	// Ensure the payload contains the "issue" key
	issueData, ok := payload["issue"].(map[string]interface{})
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload: missing issue data"})
		return
	}

	// Retrieve the issue ID from the GitHub payload
	issueGithubID, ok := issueData["id"].(float64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload: missing issue ID"})
		return
	}

	// Fetch the existing issue from the database using the GitHub ID
	var issue models.Issue
	if err := uc.db.Where("github_id = ?", int(issueGithubID)).First(&issue).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Issue not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch issue", "details": err.Error()})
		}
		return
	}

	// Update the issue with the new data from the webhook payload
	if state, ok := issueData["state"].(string); ok {
		issue.Status = state
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload: missing issue state"})
		return
	}

	// Save the updated issue back to the database
	if err := uc.db.Save(&issue).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update issue", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "issue": issue})
}
