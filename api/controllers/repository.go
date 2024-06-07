package controllers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"open-bounties-api/models"
	"open-bounties-api/services"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RepositoryController struct {
	repositoryService *services.RepositoryService
	issueService      *services.IssueService
	db                *gorm.DB
}

func NewRepositoryController(db *gorm.DB, repositoryService *services.RepositoryService, issueService *services.IssueService) *RepositoryController {
	return &RepositoryController{
		repositoryService: repositoryService,
		issueService:      issueService,
		db:                db,
	}
}

func (uc *RepositoryController) CreateRepository(c *gin.Context) {
	var newRepository models.Repository
	if err := c.ShouldBindJSON(&newRepository); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	registeredRepository, err := uc.repositoryService.CreateRepository(c, newRepository)
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
func verifyWebhookSignature(secret, payload, signature string) bool {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payload))
	expectedMAC := mac.Sum(nil)
	expectedSignature := "sha256=" + hex.EncodeToString(expectedMAC)
	return hmac.Equal([]byte(signature), []byte(expectedSignature))
}

func (uc *RepositoryController) IssueGithubWebhook(c *gin.Context) {
	secret := os.Getenv("GITHUB_WEBHOOK_SECRET")

	// Read the signature from the headers
	signature := c.GetHeader("X-Hub-Signature-256")
	if signature == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Signature header missing"})
		log.Println("Signature header missing")
		return
	}

	// Read the payload
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read body"})
		log.Println("Unable to read body")
		return
	}

	// Verify the payload signature
	if !verifyWebhookSignature(secret, string(body), signature) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid signature"})
		log.Println("Invalid signature")
		return
	}

	// Parse the JSON payload
	var payload map[string]interface{}
	if err := json.Unmarshal(body, &payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		log.Println("Invalid payload")
		return
	}
	log.Printf("Received payload: %v", payload)

	// Ensure the payload contains the "issue" key
	issueData, ok := payload["issue"].(map[string]interface{})
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload: missing issue data"})
		log.Println("Invalid payload: missing issue data")
		return
	}

	// Retrieve the issue ID from the GitHub payload
	issueGithubID, ok := issueData["id"].(float64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload: missing issue ID"})
		log.Println("Invalid payload: missing issue ID")
		return
	}

	// Fetch the existing issue from the database using the GitHub ID
	var issue models.Issue
	if err := uc.db.Where("github_id = ?", int(issueGithubID)).First(&issue).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Issue not found"})
			log.Println("Issue not found")
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch issue", "details": err.Error()})
			log.Printf("Failed to fetch issue: %s", err)
		}
		return
	}

	// Update the issue with the new data from the webhook payload
	uc.issueService.UpdateIssueFromGithubPayload(c, &issue, issueData)

	c.JSON(http.StatusOK, gin.H{"status": "success", "issue": issue})
}
