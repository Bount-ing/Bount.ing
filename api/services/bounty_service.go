package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"open-bounties-api/models"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stripe/stripe-go/v74/invoice"
	"gorm.io/gorm"
)

type BountyService struct {
	db             *gorm.DB
	issueService   *IssueService
	discordService *DiscordService
}

// NewBountyService creates a new BountyService
func NewBountyService(db *gorm.DB, issueService *IssueService, discordService *DiscordService) *BountyService {
	return &BountyService{
		db:             db,
		issueService:   issueService,
		discordService: discordService,
	}
}

// FetchAllBounties returns all bounties from the database
func (s *BountyService) FetchAllBounties() ([]models.Bounty, error) {
	var bounties []models.Bounty
	if err := s.db.Find(&bounties).Error; err != nil {
		return nil, err
	}
	return bounties, nil
}

// FetchAllUncofirmedBounties returns all unconfirmed bounties from the database
func (s *BountyService) FetchAllUncofirmedBounties() ([]models.Bounty, error) {
	var bounties []models.Bounty
	if err := s.db.Find(&bounties).Where("status = unconfirmed").Error; err != nil {
		return nil, err
	}
	return bounties, nil
}

// FetchBountyByID retrieves a bounty by its ID from the database
func (s *BountyService) FetchBountyById(id uint) (*models.Bounty, error) {
	var bounty models.Bounty
	if err := s.db.First(&bounty, id).Error; err != nil {
		return nil, err
	}
	return &bounty, nil
}

// CreateBounty creates a new bounty in the database
func (s *BountyService) findOrCreateIssue(c *gin.Context, bounty models.Bounty) (*models.Issue, error) {
	var issue models.Issue
	err := s.db.Where("github_id = ?", bounty.IssueGithubID).First(&issue).Error
	if err == nil {
		// Issue found in the database
		return &issue, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Some other error occurred
		log.Printf("Database error: %s", err)
		return nil, err
	}

	log.Printf("Issue not found in database, fetching from GitHub for ID: %d", bounty.IssueGithubID)
	githubToken, err := s.extractGithubTokenFromContext(c)
	if err != nil {
		log.Printf("Failed to extract GitHub token: %s", err)
		return nil, err
	}

	githubIssue, err := fetchGitHubIssueDetails(bounty.IssueGithubURL, githubToken)
	if err != nil {
		log.Printf("Failed to fetch GitHub issue details: %s", err)
		return nil, err
	}

	// Create the new issue in the database
	issue = models.Issue{
		GithubID:    githubIssue.ID,
		GithubURL:   bounty.IssueGithubURL,
		Title:       githubIssue.Title,
		Description: githubIssue.Body,
		Status:      githubIssue.State,
	}

	createdIssue, err := s.issueService.CreateIssue(c, issue)
	if err != nil {
		log.Printf("Failed to create issue in database: %s", err)
		return nil, err
	}

	return createdIssue, nil
}

func fetchGitHubIssueDetails(issueURL, token string) (*GitHubIssue, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", issueURL, nil)
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

	return &githubIssue, nil
}

func (s *BountyService) CreateBounty(c *gin.Context, bounty models.Bounty) (*models.Bounty, error) {
	log.Println("Creating bounty")

	// Validate the bounty type
	if err := models.ValidateBountyType(bounty.BountyType); err != nil {
		log.Printf("Validation error: %s", err)
		return nil, err
	}

	// Check if the issue exists in the database
	issue, err := s.findOrCreateIssue(c, bounty)
	if err != nil {
		log.Printf("Failed to find or create issue: %s", err)
		return nil, err
	}

	// Set the IssueID in the bounty
	bounty.IssueID = issue.ID

	// Create the bounty in the database
	if err := s.db.Create(&bounty).Error; err != nil {
		log.Printf("Failed to create bounty in database: %s", err)
		return nil, err
	}

	log.Println("Bounty created successfully")
	s.bountyCreationHook(bounty)
	return &bounty, nil
}

// extractGithubTokenFromContext extracts the GitHub token from the request context
func (s *BountyService) extractGithubTokenFromContext(c *gin.Context) (string, error) {
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
		return "", errors.New("github token not found in token claims")
	}

	return githubToken, nil
}

// UpdateBounty updates an existing bounty in the database
func (s *BountyService) UpdateBounty(id uint, updatedData models.Bounty) (*models.Bounty, error) {
	var bounty models.Bounty
	if err := s.db.First(&bounty, id).Error; err != nil {
		return nil, err
	}

	if err := s.db.Model(&bounty).Updates(updatedData).Error; err != nil {
		return nil, err
	}
	return &bounty, nil
}

// DeleteBounty deletes a bounty from the database
func (s *BountyService) DeleteBounty(id uint) error {
	var bounty models.Bounty
	if err := s.db.First(&bounty, id).Error; err != nil {
		return err
	}
	if err := s.db.Delete(&bounty).Error; err != nil {
		return err
	}
	return nil
}

func (s *BountyService) bountyCreationHook(bounty models.Bounty) {
	// Get the issue details
	issue_id := bounty.IssueID
	issue := models.Issue{}
	if err := s.db.Where("id = ?", issue_id).First(&issue).Error; err != nil {
		log.Printf("Failed to fetch issue details: %s", err)
		return
	}
	user := models.User{}
	if err := s.db.Where("id = ?", bounty.OwnerID).First(&user).Error; err != nil {
		log.Printf("Failed to fetch user details: %s", err)
		return
	}

	// Send a notification to discord
	s.discordService.SendBountyCreationNotification(bounty, issue, user)
}

func (s *BountyService) CancelBounty(bountyID uint) error {
	// get the bounty from db
	var dbBounty models.Bounty
	if err := s.db.Where("id = ?", bountyID).First(&dbBounty).Error; err != nil {
		log.Printf("Failed to fetch bounty details: %s", err)
		return err
	}

	dbBounty.Status = "cancelled"
	if err := s.db.Save(&dbBounty).Error; err != nil {
		log.Printf("Failed to update bounty status: %s", err)
		return err
	}
	return nil
}

func (s *BountyService) FinalizeBounty(bountyID uint) error {
	// get the bounty from db
	var dbBounty models.Bounty
	if err := s.db.Where("id = ?", bountyID).First(&dbBounty).Error; err != nil {
		log.Printf("Failed to fetch bounty details: %s", err)
		return err
	}

	_, err := invoice.FinalizeInvoice(dbBounty.StripeInvoiceID, nil)
	if err != nil {
		log.Printf("Failed to fetch bounty details: %s", err)
		return err
	}

	dbBounty.Status = "finalized"
	if err := s.db.Save(&dbBounty).Error; err != nil {
		log.Printf("Failed to update bounty status: %s", err)
		return err
	}
	return nil
}
