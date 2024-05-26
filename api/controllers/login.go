package controllers

import (
	"io/ioutil"
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
	GithubAuthorizeURL = "https://github.com/login/oauth/authorize"
	GithubTokenURL     = "https://github.com/login/oauth/access_token"
	GithubClientID     = "GITHUB_CLIENT_ID"
	GithubClientSecret = "GITHUB_CLIENT_SECRET"
	JWTSecretKey       = "your_secret_key"
	RedirectURL        = "https://yourapp.com/oauth/callback"
)

func (ctl *LoginController) LoginWithGithub(c *gin.Context) {
	params := url.Values{
		"client_id":    {os.Getenv(GithubClientID)},
		"scope":        {"read:user"},
		"redirect_uri": {RedirectURL},
	}
	redirectURL := GithubAuthorizeURL + "?" + params.Encode()
	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}

func (ctl *LoginController) GithubCallback(c *gin.Context) {
	code := c.Query("code")
	requestData := url.Values{
		"client_id":     {os.Getenv(GithubClientID)},
		"client_secret": {os.Getenv(GithubClientSecret)},
		"code":          {code},
	}

	resp, err := http.PostForm(GithubTokenURL, requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to request GitHub token"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read response body"})
		return
	}

	jwtToken, err := generateJWT(string(body))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": jwtToken})
}

func generateJWT(accessToken string) (string, error) {
	claims := jwt.MapClaims{
		"access_token": accessToken, // Use the actual data claims relevant for your app
		"exp":          time.Now().Add(72 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWTSecretKey))
}
