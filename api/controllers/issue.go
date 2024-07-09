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

	registeredIssue, err := uc.issueService.CreateIssue(c, newIssue)
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
