package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"open-bounties-api/services"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/oauth"
)

type LoginController struct {
	userService *services.UserService
}

func NewLoginController(userService *services.UserService) *LoginController {
	return &LoginController{
		userService: userService,
	}
}

// Constants for API URLs and client settings
const (
	GithubTokenURL = "https://github.com/login/oauth/access_token"

	GithubClientID     = "GITHUB_CLIENT_ID"
	GithubClientSecret = "GITHUB_CLIENT_SECRET"
	GithubRedirect     = "GITHUB_REDIRECT_URL"
	JWTSecretKey       = "JWT_SECRET_KEY"
)

func (ctl *LoginController) GithubCallback(c *gin.Context) {
	AppRedirULR := os.Getenv(GithubRedirect)
	code := c.Query("code")
	log.Printf("Received code: %s", code)

	requestData := url.Values{
		"client_id":     {os.Getenv(GithubClientID)},
		"client_secret": {os.Getenv(GithubClientSecret)},
		"code":          {code},
	}
	log.Printf("Request data: %v", requestData)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s?%s", GithubTokenURL, requestData.Encode()), nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Faild to create request"})
		return
	}
	req.Header.Add("accept", "application/json")

	httpClient := http.Client{}
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

	u, err := ctl.userService.VerifyGitHubToken(githubToken.AccessToken)
	if err != nil {
		log.Printf("Error verifying GitHub token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify GitHub token"})
		return

	}
	log.Printf("Verified user: %v", u)
	go ctl.syncGithubInfo(githubToken.AccessToken)

	jwtToken, err := generateJWT(u.ID, githubToken.AccessToken)
	if err != nil {
		log.Printf("Error generating JWT: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	log.Printf("Generated JWT: %s", jwtToken)

	c.Redirect(http.StatusFound, fmt.Sprintf("%s?token=%s", AppRedirULR, jwtToken))
}

func (ctl *LoginController) syncGithubInfo(accessToken string) {
	// Fetch all bounties created by the user
	// Retrieve all bountie's issues from github
	// update the issues in the database
}

func generateJWT(userId uint, accessToken string) (string, error) {
	claims := jwt.MapClaims{
		"access_token": accessToken,
		"user_id":      userId,
		"exp":          time.Now().Add(72 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv(JWTSecretKey)))
	if err != nil {
		log.Printf("Error signing JWT: %v", err)
		return "", err
	}
	return signedToken, nil
}

// Constants for API URLs and client settings
const (
	StripeSecret   = "STRIPE_CLIENT_SECRET"
	StripeRedirect = "STRIPE_REDIRECT_URL"
)

func (ctl *LoginController) StripeCallback(c *gin.Context) {
	AppRedirULR := os.Getenv(StripeRedirect)
	code := c.Query("code")
	log.Printf("Received code: %s", code)

	params := stripe.OAuthTokenParams{
		GrantType: stripe.String("authorization_code"),
		Code:      stripe.String(code),
	}

	resp, err := oauth.New(&params)
	if err != nil {
		log.Printf("Error oauth Stripe: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Stripe Oauth Failed"})
		return
	}
	log.Println("User Stripe Id:", resp.StripeUserID)
	c.Redirect(http.StatusFound, fmt.Sprintf("%s?account_id=%s", AppRedirULR, resp.StripeUserID))
}
