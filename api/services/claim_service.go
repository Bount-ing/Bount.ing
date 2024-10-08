package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"open-bounties-api/models"
	"os"
	"strings"
	"time"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/invoice"
	"gorm.io/gorm"
)

type ClaimService struct {
	db *gorm.DB
}

type PRDetails struct {
	ID          int    `json:"id"`           // The unique ID of the PR
	URL         string `json:"url"`          // The URL of the PR
	Number      int    `json:"number"`       // The number of the PR in the repository
	State       string `json:"state"`        // The state of the PR (e.g., open, closed, merged)
	Title       string `json:"title"`        // The title of the PR
	Body        string `json:"body"`         // The body description of the PR
	Author      string `json:"author"`       // The author (username) of the PR
	Merged      bool   `json:"merged"`       // Indicates if the PR has been merged
	MergeCommit string `json:"merge_commit"` // The merge commit hash if the PR is merged
	CreatedAt   string `json:"created_at"`   // The creation date of the PR
	UpdatedAt   string `json:"updated_at"`   // The last update date of the PR
	ClosedAt    string `json:"closed_at"`    // The date the PR was closed, if applicable
	MergedAt    string `json:"merged_at"`    // The date the PR was merged, if applicable
	User        struct {
		Login string `json:"login"`
	} `json:"user"`
}

func NewClaimService(db *gorm.DB) *ClaimService {
	return &ClaimService{db: db}
}

// FetchAllClaims returns all claims from the database
func (s *ClaimService) FetchAllClaims() ([]models.Claim, error) {
	var claims []models.Claim
	result := s.db.Find(&claims)
	return claims, result.Error
}

// FetchClaimByID retrieves an claim by its ID from the database
func (s *ClaimService) FetchClaimById(id uint) (*models.Claim, error) {
	var claim models.Claim
	result := s.db.First(&claim, id)
	return &claim, result.Error
}

// CreateClaim creates a new claim in the database
func (s *ClaimService) CreateClaim(claim models.Claim) (*models.Claim, error) {
	result := s.db.Create(&claim)
	return &claim, result.Error
}

func (s *ClaimService) UpdateClaim(id uint, updatedData models.Claim) (*models.Claim, error) {
	var claim models.Claim
	result := s.db.First(&claim, id)
	if result.Error != nil {
		return nil, result.Error
	}

	saveResult := s.db.Save(&claim)
	if saveResult.Error != nil {
		return nil, saveResult.Error
	}
	return &claim, nil
}

func (s *ClaimService) DeleteClaim(id uint) error {
	var claim models.Claim
	result := s.db.First(&claim, id)
	if result.Error != nil {
		return result.Error
	}
	deleteResult := s.db.Delete(&claim)
	return deleteResult.Error
}

func (s *ClaimService) SolveClaimByPullRequest(url string, issue models.Issue) (bool, error) {
	// Fetch PR details from GitHub
	prDetails, err := fetchPRDetails(url)
	if err != nil {
		return false, err
	}

	// Check if PR is merged
	if !prDetails.Merged {
		return false, fmt.Errorf("PR is not merged")
	}

	// Check if the user exists in the database
	var user models.User
	if err := s.db.Where("username = ?", prDetails.User.Login).First(&user).Error; err != nil {
		return false, err
	}

	// Check if the claim exists, create if not
	var claim models.Claim
	result := s.db.Where("issue_id = ? AND owner_id = ?", issue.ID, user.ID).First(&claim)
	if result.Error != nil {
		// Create a new claim if not found
		claim = models.Claim{IssueID: issue.ID, OwnerID: user.ID, Status: "Open"}
		if err := s.db.Create(&claim).Error; err != nil {
			return false, err
		}
	}

	// Retrieve all bounties for the issue, associate bounties to the claim
	var bounties []models.Bounty
	s.db.Where("issue_id = ?", issue.ID).Find(&bounties)

	for _, bounty := range bounties {
		if bounty.Status != "open" {
			log.Printf("Bounty %d is not open", bounty.ID)
			continue
		}
		pr_created_date, _ := time.Parse(time.RFC3339, prDetails.CreatedAt)
		if pr_created_date.After(bounty.CreatedAt) {
			log.Printf("Bounty %d was created after the PR", bounty.ID)
			continue
		}
		if pr_created_date.Before(bounty.StartAt) {
			log.Printf("Pull request was created before the bounty start date")
			continue
		}
		if pr_created_date.After(bounty.EndAt) {
			log.Printf("Pull request was created after the bounty end date")
			continue
		}
		bounty.Status = "waiting for payment"
		s.db.Model(&claim).Association("Bounties").Append(&bounty)
	}

	// Pay the user and update the bounties status to paid
	if err := s.payUserAndUpdateBounties(user, bounties); err != nil {
		return false, err
	}

	// Update the claim status to solved
	s.db.Model(&claim).Update("status", "Solved")

	return true, nil
}

// fetchPRDetails fetches the details of a pull request from GitHub.
func fetchPRDetails(url string) (*PRDetails, error) {
	// Convert GitHub PR URL to API URL
	apiURL, err := convertToAPIURL(url)
	if err != nil {
		log.Printf("Failed to convert URL: %s", err)
		return nil, err
	}

	// Create a new HTTP request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Optional: Set up authentication using a personal access token
	token := os.Getenv("GITHUB_TOKEN")
	if token != "" {
		req.Header.Set("Authorization", "token "+token)
	}

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	// Read and parse the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for non-200 status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response: %s", resp.Status)
	}

	var details PRDetails
	if err := json.Unmarshal(body, &details); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &details, nil
}

// convertToAPIURL converts a GitHub PR URL to the corresponding API URL.
func convertToAPIURL(inputURL string) (string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}

	// Split the path to get the repository details
	parts := strings.Split(strings.Trim(parsedURL.Path, "/"), "/")
	if len(parts) < 4 || parts[2] != "pull" {
		log.Printf("Invalid URL format: %s", inputURL)
		return "", fmt.Errorf("invalid URL format: %s", inputURL)
	}

	// Extract owner, repo, and pull request number
	owner := parts[0]
	repo := parts[1]
	prNumber := parts[3]

	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls/%s", owner, repo, prNumber)
	return apiURL, nil
}

func (c ClaimService) payUserAndUpdateBounties(user models.User, bounties []models.Bounty) error {
	var owner models.User
	// Placeholder for payment logic
	for _, bounty := range bounties {
		if bounty.Status != "waiting for payment" {
			log.Printf("Bounty %d is not waiting for payment", bounty.ID)
			continue
		}
		if owner.ID != bounty.OwnerID {
			status := c.db.First(&owner, bounty.OwnerID)
			if status.Error != nil {
				return status.Error
			}
		}
		// Amount must be in cents
		i, err := createInvoice(owner.StipeAccountID, user.StipeAccountID, bounty.IssueGithubURL, int64(bounty.Amount*100))
		if err != nil {
			return err
		}
		bounty.Status = "unconfirmed"
		bounty.StripeInvoiceID = i.ID
		status := c.db.Save(&bounty)
		if status.Error != nil {
			return status.Error
		}
		log.Printf("Paid bounty %d to user %s", bounty.ID, user.Username)
		log.Printf("Bounty ID: %d, Issue ID: %d, Owner ID: %d, Amount: %f, Status: %s, Created At: %s, Updated At: %s", bounty.ID, bounty.IssueID, bounty.OwnerID, bounty.Amount, bounty.Status, bounty.CreatedAt, bounty.UpdatedAt)
	}
	return nil
}

// createInvoice Creates an Invoice that takes a fee of 8%
// Recives IDs of the sender and destination accounts, url of the issue and the amount of the bounty
func createInvoice(sender, dest, url string, amount int64) (*stripe.Invoice, error) {
	params := stripe.InvoiceParams{
		Customer:             stripe.String(sender),
		ApplicationFeeAmount: stripe.Int64(amount * 8 / 100),
		AutoAdvance:          stripe.Bool(false),
		Description:          stripe.String("[Bount.ing] Bounty claimed for issue: " + url),
		TransferData: &stripe.InvoiceTransferDataParams{
			Amount:      stripe.Int64(amount),
			Destination: stripe.String(dest),
		},
	}
	i, err := invoice.New(&params)
	if err != nil {
		return nil, err
	}
	return i, nil
}
