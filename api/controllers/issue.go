package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"

	"github.com/bount-ing/bount.ing/api/db"
	"github.com/bount-ing/bount.ing/api/models"
	"gorm.io/gorm"
)

func CreateIssue(issue models.Issue) (models.Issue, error) {
	err := db.DB.Create(&issue)
	if err.Error != nil {
		log.Print(err.Error)
		return issue, err.Error
	}
	return issue, nil
}

func GetIssue(issueID string) (models.Issue, error) {
	var issue models.Issue
	err := db.DB.First(&issue, issueID)
	if err.Error != nil {
		log.Print(err.Error)
		return issue, err.Error
	}
	return issue, nil
}

func GetIssueByUrl(issueUrl string) (models.Issue, error) {
	var issue models.Issue
	log.Print("Issue URL: ", issueUrl)
	err := db.DB.Where("url = ?", issueUrl).First(&issue)
	if err.Error != nil {
		log.Print(err.Error)
		return issue, err.Error
	}
	return issue, nil
}

func GetIssues() ([]models.Issue, error) {
	var issues []models.Issue
	err := db.DB.Find(&issues)
	if err.Error != nil {
		log.Print(err.Error)
		return issues, err.Error
	}
	return issues, nil
}

func UpdateIssue(issue models.Issue) error {
	err := db.DB.Save(&issue)
	if err.Error != nil {
		log.Print(err.Error)
		return err.Error
	}
	return nil
}

func DeleteIssue(issueID string) error {
	var issue models.Issue
	err := db.DB.First(&issue, issueID)
	if err.Error != nil {
		log.Print(err.Error)
		return err.Error
	}
	return nil
}

// controllers/issue.go

func GetMyIssues(userID uint) ([]models.Issue, error) {
	var issues []models.Issue

	// Preload relationships, including specific PRVerification fields
	err := db.DB.Preload("Bounties").
		Preload("Bounties.Variables").
		Preload("Bounties.Claims", "status = ?", "pending"). // Only pending claims
		Preload("Bounties.Claims.BountySponsorCheck").       // Sponsor PR verification
		Preload("Bounties.Claims.BountyClaimerCheck").       // Author PR verification
		Find(&issues).Error

	if err != nil {
		log.Printf("Error fetching issues for user %d: %v", userID, err)
		return nil, err
	}

	// Filter issues based on user's sponsorship of bounties
	var myIssues []models.Issue
	for _, issue := range issues {
		var myBounties []models.Bounty
		for _, bounty := range issue.Bounties {
			if bounty.SponsorID == userID {
				// Filter and organize claims directly
				bounty.Claims = filterAndOrganizeClaims(bounty.Claims)
				myBounties = append(myBounties, bounty)
			}
		}

		if len(myBounties) > 0 {
			issue.Bounties = myBounties
			myIssues = append(myIssues, issue)
		}
	}

	return myIssues, nil
}

func GetIssueBounties(issueID string) (models.Issue, error) {
	var issue models.Issue

	err := db.DB.Preload("Bounties").
		Preload("Bounties.Variables").
		Preload("Bounties.Claims", "status = ?", "pending").
		Preload("Bounties.Claims.BountyClaimerCheck"). // Changed from PRVerification
		Preload("Bounties.Claims.BountySponsorCheck"). // Add other checks if needed
		Preload("Bounties.Claims.BountySystemCheck").  // Add other checks if needed
		First(&issue, issueID)

	if err.Error != nil {
		log.Print(err.Error)
		return issue, err.Error
	}

	// Organize claims for each bounty
	for i := range issue.Bounties {
		issue.Bounties[i].Claims = filterAndOrganizeClaims(issue.Bounties[i].Claims)
	}

	return issue, nil
}

// Helper function to filter and organize claims
func filterAndOrganizeClaims(claims []models.Claim) []models.Claim {
	// You might want to sort claims by date, filter by status, etc.
	// This is just a basic example
	if claims == nil {
		return []models.Claim{}
	}

	// Sort claims by creation date (newest first)
	sort.Slice(claims, func(i, j int) bool {
		return claims[i].CreatedAt.After(claims[j].CreatedAt)
	})

	return claims
}

func GetIssueByID(issueID uint) (models.Issue, error) {
	var issue models.Issue
	err := db.DB.First(&issue, issueID)
	if err.Error != nil {
		log.Print(err.Error)
		return issue, err.Error
	}
	return issue, nil
}

// First, let's add a method to extract sponsor and repo from the URL
func ParseGitHubURL(url string) (sponsor, repo string, err error) {
	// GitHub issue URLs are in the format: https://github.com/sponsor/repo/issues/number
	parts := strings.Split(url, "/")
	if len(parts) < 5 {
		return "", "", models.ErrGitHubInvalidURL
	}
	return parts[3], parts[4], nil
}

type GitHubUser struct {
	AvatarURL string `json:"avatar_url"`
}

func GetGitHubAvatarURL(username string) (string, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", models.ErrGitHubAvatarNotFound
	}

	var user GitHubUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return "", err
	}

	return user.AvatarURL, nil
}

func CreateIssueFromGitHub(url string) (*models.Issue, error) {
	githubCtrl := NewGitHubController()

	// Fetch issue from GitHub
	githubIssue, err := githubCtrl.FetchGitHubIssue(url)
	if err != nil {
		return nil, err
	}

	//extract login from githubIssueurl
	sponsorLogin := strings.Split(url, "/")[3]

	avatarURL, err := GetGitHubAvatarURL(sponsorLogin)
	if err != nil {
		return nil, err
	}

	// Convert GitHub issue to our Issue model
	issue := &models.Issue{
		ForeignID:   fmt.Sprintf("github_%d", githubIssue.ID),
		HostID:      1,
		URL:         githubIssue.HTMLURL,
		AvatarURL:   avatarURL,
		Title:       githubIssue.Title,
		Description: githubIssue.Body,
		Status:      githubIssue.State,
	}

	// Check if issue already exists
	var existingIssue models.Issue
	err = db.DB.Where("foreign_id = ?", issue.ForeignID).First(&existingIssue).Error

	// If the error is anything other than "record not found", return the error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// If we found an existing issue, return it
	if err == nil {
		return &existingIssue, nil
	}

	// Create the issue
	createdIssue, err := CreateIssue(*issue)
	if err != nil {
		return nil, err
	}

	return &createdIssue, nil
}

func GetIssueByURL(url string) (*models.Issue, error) {
	var issue models.Issue
	err := db.DB.Where("url = ?", url).First(&issue)
	if err.Error != nil {
		return nil, err.Error
	}
	return &issue, nil
}
