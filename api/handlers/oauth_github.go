package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/bount-ing/bount.ing/api/controllers"
	"github.com/bount-ing/bount.ing/api/models"
	"github.com/gin-gonic/gin"
)

const GithubTokenURL = "https://github.com/login/oauth/access_token"

func OAuthGithubCallback(c *gin.Context) {
	AppRedirURL := os.Getenv("GITHUB_REDIRECT_URL")
	code := c.DefaultQuery("code", "")
	state := c.DefaultQuery("state", "")
	log.Printf("Received code: %s, state: %s", code, state)

	// extract the state json payload (bs64 encoded)
	statePayload, err := controllers.DecryptAndReadOAuthState(state, "https://github.com")
	if err != nil {
		log.Printf("Error decoding state: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode state"})
		return
	}
	log.Printf("Decoded state: %s", state)
	log.Printf("Decoded state payload: %+v", statePayload)

	// Verify the state is not expired
	if statePayload.ExpiresAt.Before(time.Now()) {
		log.Printf("State expired: %v", statePayload.ExpiresAt)
		c.JSON(http.StatusBadRequest, gin.H{"error": "State expired"})
		return
	}

	// Proceed with token exchange if state is valid
	requestData := url.Values{
		"client_id":     {os.Getenv("GITHUB_CLIENT_ID")},
		"client_secret": {os.Getenv("GITHUB_CLIENT_SECRET")},
		"code":          {code},
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s?%s", GithubTokenURL, requestData.Encode()), nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create request"})
		return
	}
	req.Header.Add("accept", "application/json")

	httpClient := http.Client{}
	log.Printf("Requesting GitHub token: %+v", req)
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Printf("Error making request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to request GitHub token"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read response body"})
		return
	}
	log.Printf("Response body: %s", body)

	var githubToken struct {
		AccessToken string `json:"access_token"`
		Scope       string `json:"scope"`
		TokenType   string `json:"token_type"`
	}
	if err := json.Unmarshal(body, &githubToken); err != nil {
		log.Printf("Error parsing response body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse response body"})
		return
	}
	log.Printf("GitHub token: %v", githubToken)

	// Optionally verify the token (to check validity)
	err = VerifyGitHubToken(githubToken.AccessToken, statePayload.UserID)
	if err != nil {
		log.Printf("Error verifying GitHub token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify GitHub token"})
		return
	}

	// Store token in session or secure storage
	c.Set("github_token", githubToken.AccessToken)

	// Redirect the user securely
	c.Redirect(http.StatusFound, fmt.Sprintf("%s?github_token=%s", AppRedirURL, githubToken.AccessToken))
}

func VerifyGitHubToken(token string, stateUserID uint) error {
	const githubUserAPIURL = "https://api.github.com/user"
	const githubEmailsAPIURL = "https://api.github.com/user/emails"

	// Create a new request to the GitHub API to fetch user data
	req, err := http.NewRequest("GET", githubUserAPIURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Add the OAuth token in the Authorization header
	req.Header.Add("Authorization", "Bearer "+token)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request to GitHub: %w", err)
	}
	defer resp.Body.Close()

	// Read and parse the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("GitHub API returned non-OK status: %s", resp.Status)
	}

	// Parse the JSON response into a struct
	var githubUser struct {
		ID    int    `json:"id"`
		Login string `json:"login"`
		Email string `json:"email"`
	}
	if err := json.Unmarshal(body, &githubUser); err != nil {
		return fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	// If email is null, fetch the user's email addresses
	if githubUser.Email == "" {
		req, err := http.NewRequest("GET", githubEmailsAPIURL, nil)
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}

		// Add the OAuth token in the Authorization header
		req.Header.Add("Authorization", "Bearer "+token)

		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("failed to send request to GitHub: %w", err)
		}
		defer resp.Body.Close()

		// Read and parse the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %w", err)
		}

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("GitHub API returned non-OK status: %s", resp.Status)
		}

		// Parse the JSON response into a slice of structs
		var emails []struct {
			Email    string `json:"email"`
			Primary  bool   `json:"primary"`
			Verified bool   `json:"verified"`
		}
		if err := json.Unmarshal(body, &emails); err != nil {
			return fmt.Errorf("failed to unmarshal email response body: %w", err)
		}

		// Find the primary and verified email address
		for _, email := range emails {
			if email.Primary && email.Verified {
				githubUser.Email = email.Email
				break
			}
		}
	}

	//convert githubUser.ID to string
	githubUserID := fmt.Sprint(githubUser.ID)

	// Check 1: stateUserID doesn't already have a github identity
	userIdentities, err := controllers.GetUserIdentities(stateUserID)
	if err != nil {
		return fmt.Errorf("failed to get user identities: %w", err)
	}
	for _, identity := range userIdentities {
		if identity.Host.Address == "https://github.com" && identity.UserExternalID != githubUserID {
			return fmt.Errorf("user already has a GitHub identity associated")
		} else if identity.Host.Address == "https://github.com" && identity.UserExternalID == githubUserID {
			// return ok
			return nil
		}
	}

	// Check 2: githubUser.ID doesn't already have an identity
	hostIdentities, err := controllers.GetHostIdentitiesFromAddress("https://github.com")
	if err != nil {
		return fmt.Errorf("failed to get host identities: %w", err)
	}
	for _, identity := range hostIdentities {
		if identity.UserExternalID == githubUserID {
			return fmt.Errorf("GitHub identity already has a user associated")
		}
	}

	// Create a new identity for the user
	newIdentity := &models.Identity{
		Username:       githubUser.Login,
		Email:          githubUser.Email,
		UserExternalID: githubUserID,
		UserID:         stateUserID,
		HostID:         1, // GitHub host ID
	}

	err = controllers.CreateIdentity(newIdentity)
	if err != nil {
		return fmt.Errorf("failed to create identity: %w", err)
	}

	return nil
}
