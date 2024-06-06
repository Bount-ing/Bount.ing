package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"open-bounties-api/models"
	"strings"

	"gorm.io/gorm"
)

type IssueService struct {
	db *gorm.DB
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

func NewIssueService(db *gorm.DB) *IssueService {
	return &IssueService{db: db}
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

func (s *IssueService) CreateIssue(token string, issue models.Issue) (*models.Issue, error) {
	result := s.db.Create(&issue)
	if result.Error != nil {
		return nil, result.Error
	}

	err := s.subscribeToGitHubWebhook(token, issue)
	if err != nil {
		return nil, err
	}

	return &issue, nil
}

func (s *IssueService) UpdateIssue(id uint, updatedData models.Issue) (*models.Issue, error) {
	var issue models.Issue
	result := s.db.First(&issue, id)
	if result.Error != nil {
		return nil, result.Error
	}

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

func (s *IssueService) subscribeToGitHubWebhook(token string, issue models.Issue) error {
	// parse owner and repo from the GitHub URL
	owner, repo, err := parseOwnerAndRepoFromGitHubURL(issue.GithubURL)
	if err != nil {
		log.Printf("Failed to parse owner and repo from GitHub URL: %s", err)
		return err
	}
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/hooks", owner, repo)
	log.Printf("Subscribing to GitHub webhook at %s", url)
	api_base_url := "https://api.bount.ing" //os.Getenv("API_BASE_URL")
	requestBody := struct {
		Name   string   `json:"name"`
		Active bool     `json:"active"`
		Events []string `json:"events"`
		Config struct {
			URL         string `json:"url"`
			ContentType string `json:"content_type"`
			InsecureSSL string `json:"insecure_ssl"`
		} `json:"config"`
	}{
		Name:   "web",
		Active: true,
		Events: []string{"issues"},
		Config: struct {
			URL         string `json:"url"`
			ContentType string `json:"content_type"`
			InsecureSSL string `json:"insecure_ssl"`
		}{
			URL:         fmt.Sprintf("%s/webhooks/github/issues/%d", api_base_url, issue.ID),
			ContentType: "json",
			InsecureSSL: "0",
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

func FetchIssueByGithubID(githubID int, token string) (*models.Issue, error) {
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
