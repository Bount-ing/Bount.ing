package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/bount-ing/bount.ing/api/models"
)

type GitHubController struct {
	client *http.Client
}

func NewGitHubController() *GitHubController {
	return &GitHubController{
		client: &http.Client{},
	}
}

func (gc *GitHubController) FetchGitHubIssue(url string) (*models.GitHubIssue, error) {
	// Extract owner, repo, and issue number from URL
	re := regexp.MustCompile(`^https://github\.com/([^/]+)/([^/]+)/issues/(\d+)$`)
	matches := re.FindStringSubmatch(url)
	if matches == nil {
		return nil, errors.New("invalid GitHub issue URL")
	}

	owner, repo, issueNum := matches[1], matches[2], matches[3]
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%s", owner, repo, issueNum)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	// Add GitHub token if available
	// req.Header.Set("Authorization", "token YOUR_GITHUB_TOKEN")

	resp, err := gc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned status: %d", resp.StatusCode)
	}

	var issue models.GitHubIssue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}

	return &issue, nil
}
