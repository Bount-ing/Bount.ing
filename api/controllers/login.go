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
	JWTSecretKey       = "JWT_SECRET_KEY"
)

func (ctl *LoginController) GithubCallback(c *gin.Context) {
	code := c.Query("code")
	requestData := url.Values{
		"client_id":     {os.Getenv(GithubClientID)},
		"client_secret": {os.Getenv(GithubClientSecret)},
		"code":          {code},
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s?%s", GithubTokenURL, requestData.Encode()), nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Faild to create request"})
		return
	}
	req.Header.Add("accept", "application/json")

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to request GitHub token"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read response body"})
		return
	}

	var githubToken struct {
		AccessToken string `json:"access_token"`
		Scope       string `json:"scope"`
		TokenType   string `json:"token_type"`
	}
	if err := json.Unmarshal(body, &githubToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse response body"})
		return
	}

	u, err := ctl.userService.VerifyGitHubToken(githubToken.AccessToken)
	if err != nil {
		return
	}
	jwtToken, err := generateJWT(u.ID, githubToken.AccessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": jwtToken})
}

func generateJWT(userId uint, accessToken string) (string, error) {
	claims := jwt.MapClaims{
		"access_token": accessToken,
		"user_id":      userId,
		"exp":          time.Now().Add(72 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv(JWTSecretKey)))
}
