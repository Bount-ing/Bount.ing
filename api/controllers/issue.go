package controllers

import (
    "github.com/gin-gonic/gin"
    "open-bounties-api/services"
    "net/http"
)

type IssueController struct {
    issueService *services.IssueService
}

func NewIssueController(issueService *services.IssueService) *IssueController {
    return &IssueController{
        issueService: issueService,
    }
}

func (ctl *IssueController) GetAllIssues(c *gin.Context) {
    issues, err := ctl.issueService.FetchAllIssues()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, issues)
}

