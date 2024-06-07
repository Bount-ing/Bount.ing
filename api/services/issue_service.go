package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"open-bounties-api/models"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type IssueService struct {
	db                *gorm.DB
	repositoryService *RepositoryService
}

type GitHubWebhookSubscriptionRequest struct {
	Events []string `json:"events"`
	Config struct {
		URL         string `json:"url"`
		ContentType string `json:"content_type"`
	} `json:"config"`
}

type GitHubIssue struct {
	Repository struct {
		FullName string `json:"full_name"`
		Owner    struct {
			Login string `json:"login"`
		} `json:"owner"`
	} `json:"repository"`
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	State string `json:"state"`
	URL   string `json:"url"`
}

func NewIssueService(db *gorm.DB, repositoryService *RepositoryService) *IssueService {
	return &IssueService{
		db:                db,
		repositoryService: repositoryService,
	}
}

// FetchAllIssues returns all issues from the database
func (s *IssueService) FetchAllIssues() ([]models.Issue, error) {
	var issues []models.Issue
	result := s.db.Find(&issues)
	return issues, result.Error
}

// FetchIssueByID retrieves an issue by its ID from the database
func (s *IssueService) FetchIssueById(id uint) (*models.Issue, error) {
	var issue models.Issue
	result := s.db.First(&issue, id)
	return &issue, result.Error
}

func (s *IssueService) CreateIssue(c *gin.Context, issue models.Issue) (*models.Issue, error) {
	log.Println("Creating Issue")

	// Check if the repository exists in the database
	repo, err := s.findOrCreateRepo(c, issue)
	if err != nil {
		log.Printf("Failed to find or create repo: %s", err)
		return nil, err
	}

	// Set the RepositoryID in the issue
	issue.RepositoryID = repo.ID

	result := s.db.Create(&issue)
	if result.Error != nil {
		return nil, result.Error
	}

	return &issue, nil
}

// extractGithubTokenFromContext extracts the GitHub token from the request context
func (s *IssueService) extractGithubTokenFromContext(c *gin.Context) (string, error) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return "", errors.New("authorization header not found")
	}

	// Strip the "Bearer " prefix if it exists
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	claims := jwt.MapClaims{}
	jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		log.Printf("Error parsing token: %s", err)
		if err == jwt.ErrSignatureInvalid {
			return "", errors.New("invalid token signature")
		}
		return "", errors.New("invalid token")
	}
	if !token.Valid {
		return "", errors.New("token is not valid")
	}

	githubToken, ok := claims["access_token"].(string)
	if !ok {
		return "", errors.New("GitHub token not found in token claims")
	}

	return githubToken, nil
}

func (s *IssueService) UpdateIssue(id uint, updatedData models.Issue) (*models.Issue, error) {
	var issue models.Issue
	result := s.db.First(&issue, id)
	if result.Error != nil {
		return nil, result.Error
	}

	// Update the fields of the issue with the new data
	issue.Title = updatedData.Title
	issue.Description = updatedData.Description
	issue.Status = updatedData.Status

	saveResult := s.db.Save(&issue)
	if saveResult.Error != nil {
		return nil, saveResult.Error
	}
	return &issue, nil
}

func (s *IssueService) DeleteIssue(id uint) error {
	var issue models.Issue
	result := s.db.First(&issue, id)
	if result.Error != nil {
		return result.Error
	}
	deleteResult := s.db.Delete(&issue)
	return deleteResult.Error
}

func (s *IssueService) subscribeToGitHubWebhook(token string, repo models.Repository) error {
	// Parse owner and repo from the GitHub URL
	owner, repo_name, err := parseOwnerAndRepoFromGitHubURL(repo.GithubURL)
	if err != nil {
		log.Printf("Failed to parse owner and repo from GitHub URL: %s", err)
		return err
	}
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/hooks", owner, repo_name)
	log.Printf("Subscribing to GitHub webhook at %s", url)
	apiBaseURL := "https://api.bount.ing" //os.Getenv("API_BASE_URL")
	secret := os.Getenv("GITHUB_WEBHOOK_SECRET")

	requestBody := struct {
		Name   string   `json:"name"`
		Active bool     `json:"active"`
		Events []string `json:"events"`
		Config struct {
			URL         string `json:"url"`
			ContentType string `json:"content_type"`
			InsecureSSL string `json:"insecure_ssl"`
			Secret      string `json:"secret"`
		} `json:"config"`
	}{
		Name:   "web",
		Active: true,
		Events: []string{"issues"},
		Config: struct {
			URL         string `json:"url"`
			ContentType string `json:"content_type"`
			InsecureSSL string `json:"insecure_ssl"`
			Secret      string `json:"secret"`
		}{
			URL:         fmt.Sprintf("%s/webhooks/github/repos/%d", apiBaseURL, repo.ID),
			ContentType: "json",
			InsecureSSL: "0",
			Secret:      secret,
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		log.Printf("Failed to marshal JSON: %s", err)
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Failed to create request: %s", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to subscribe to GitHub webhook: %s", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := io.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		log.Printf("Failed to subscribe to GitHub webhook: %s. Response body: %s", resp.Status, bodyString)
		return fmt.Errorf("failed to subscribe to GitHub webhook: %s. Response body: %s", resp.Status, bodyString)
	}
	return nil
}

func parseOwnerAndRepoFromGitHubURL(url string) (string, string, error) {
	owner := ""
	repo := ""
	// Split the URL by "/"
	urlParts := strings.Split(url, "/")
	if len(urlParts) < 6 {
		return "", "", fmt.Errorf("invalid GitHub URL: %s", url)
	}
	owner = urlParts[4]
	repo = urlParts[5]
	return owner, repo, nil
}

func FetchRepoByGithubID(githubID int, token string) (*models.Issue, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/issues/%d", "owner", githubID)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Failed to create request: %s", err)
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to fetch issue details: %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to fetch issue details: %s", resp.Status)
		return nil, fmt.Errorf("failed to fetch issue details: %s", resp.Status)
	}

	var githubIssue GitHubIssue
	if err := json.NewDecoder(resp.Body).Decode(&githubIssue); err != nil {
		log.Printf("Failed to decode issue details: %s", err)
		return nil, err
	}

	issue := models.Issue{
		GithubID:    githubIssue.ID,
		GithubURL:   githubIssue.URL,
		Title:       githubIssue.Title,
		Description: githubIssue.Body,
		Status:      githubIssue.State,
	}
	return &issue, nil
}

// findOrCreateRepo finds or creates a repository in the database
func (s *IssueService) findOrCreateRepo(c *gin.Context, issue models.Issue) (*models.Repository, error) {
	var repo models.Repository
	urlParts := strings.SplitN(issue.GithubURL, "/", 6)
	if len(urlParts) < 6 {
		return nil, fmt.Errorf("invalid GitHub URL: %s", issue.GithubURL)
	}
	repoURL := urlParts[4]
	err := s.db.Where("github_url = ?", repoURL).First(&repo).Error
	if err == nil {
		// Repo found in the database
		return &repo, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Some other error occurred
		log.Printf("Database error: %s", err)
		return nil, err
	}

	log.Printf("Repo not found in database, fetching from GitHub for URL: %s", issue.GithubURL)
	githubToken, err := s.extractGithubTokenFromContext(c)
	if err != nil {
		log.Printf("Failed to extract GitHub token: %s", err)
		return nil, err
	}

	githubRepo, err := fetchGitHubRepoDetails(issue.GithubURL, githubToken)
	if err != nil {
		log.Printf("Failed to fetch GitHub repo details: %s", err)
		return nil, err
	}

	// Create the new repository in the database
	repo = models.Repository{
		GithubID:             githubRepo.ID,
		GithubURL:            githubRepo.URL,
		GithubWebhookEnabled: false,
		Name:                 githubRepo.Repository.FullName,
	}

	// Extract GitHub token from context
	token, err := s.extractGithubTokenFromContext(c)
	if err != nil {
		log.Printf("Failed to extract GitHub token: %s", err)
		return nil, err
	}

	createdRepo, err := s.repositoryService.CreateRepository(c, repo)
	if err != nil {
		log.Printf("Failed to create issue in database: %s", err)
		return nil, err
	}

	err = s.subscribeToGitHubWebhook(token, *createdRepo)
	if err != nil {
		log.Printf("Failed to subscribe to GitHub webhook: %s", err)
	} else {
		createdRepo.GithubWebhookEnabled = true
	}

	updatedRepo, err := s.repositoryService.UpdateRepository(createdRepo.ID, *createdRepo)
	if err != nil {
		log.Printf("Failed to update repository: %s", err)
		return nil, err
	}

	return updatedRepo, nil
}

func fetchGitHubRepoDetails(issueURL, token string) (*GitHubIssue, error) {
	//  https://api.github.com/repos/Smojify/Smojify-Android/issues/35
	tmpURL := strings.Split(issueURL, "/")
	repoURL := fmt.Sprintf("https://api.github.com/repos/%s/%s", tmpURL[4], tmpURL[5])
	client := &http.Client{}
	req, err := http.NewRequest("GET", repoURL, nil)
	if err != nil {
		log.Printf("Failed to create request: %s", err)
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to fetch repo details: %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to fetch repo details: %s", resp.Status)
		return nil, fmt.Errorf("failed to fetch repo details: %s", resp.Status)
	}

	var githubRepo GitHubIssue
	if err := json.NewDecoder(resp.Body).Decode(&githubRepo); err != nil {
		log.Printf("Failed to decode repo details: %s", err)
		return nil, err
	}

	return &githubRepo, nil
}

func (s *IssueService) UpdateIssueFromGithubPayload(c *gin.Context, issue *models.Issue, issueData map[string]interface{}) error {
	// Ensure all required fields are present and not nil, with default values if necessary

	log.Println("Updating Issue")
	log.Printf("Issue Data: %v", issueData)
	// Title
	if title, ok := issueData["title"].(string); ok && title != "" {
		issue.Title = title
	} else {
		issue.Title = "No Title Provided" // Default title or handle the absence of a title
	}

	// Description (Body)
	if body, ok := issueData["body"].(string); ok && body != "" {
		issue.Description = body
	} else {
		issue.Description = "No Description Provided" // Default description or handle the absence of a description
	}

	// Status (State)
	if state, ok := issueData["state"].(string); ok && state != "" {
		issue.Status = state
	} else {
		issue.Status = "open" // Default state, assuming it's open if not specified
	}

	// ClosedAt
	if closedAt, ok := issueData["closed_at"].(string); ok && closedAt != "" {
		issue.ClosedAt = closedAt
	} else {
		issue.ClosedAt = "" // Default to an empty string if closed_at is not specified
	}

	// Attempt to save the updated issue to the database
	db := c.MustGet("db").(*gorm.DB)
	result := db.Save(issue)

	// Check for errors during the save operation
	if result.Error != nil {
		log.Printf("Failed to update issue: %v", result.Error) // Log the error
		return result.Error                                    // Return the error to be handled by the caller
	}

	// Successfully updated the issue
	return nil
}
