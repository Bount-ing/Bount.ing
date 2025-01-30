package handlers

import (
	"net/http"
	"time"

	"github.com/bount-ing/bount.ing/api/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/oauth"
)

func OAuthStripeCallback(c *gin.Context) {
	code := c.DefaultQuery("code", "")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing code"})
		return
	}

	state := c.DefaultQuery("state", "")

	// extract the state json payload (bs64 encoded)
	statePayload, err := controllers.DecryptAndReadOAuthState(state, "https://connect.stripe.com")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode state"})
		return
	}

	// Verify the state is not expired
	if statePayload.ExpiresAt.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "State expired"})
		return
	}

	// Exchange the authorization code for an access token
	params := &stripe.OAuthTokenParams{
		GrantType: stripe.String("authorization_code"),
		Code:      stripe.String(code),
	}

	// Make the OAuth token request
	token, err := oauth.New(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange code for token"})
		return
	}

	// Save the connected account ID to the user's profile
	err = controllers.SaveStripeConnectedAccountID(statePayload.UserID, token.StripeUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save connected account ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Connected account ID saved"})
}
