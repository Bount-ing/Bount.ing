package handlers

import (
	"log"
	"net/http"
	"os"
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
	statePayload, err := controllers.DecryptAndReadOAuthState(state, "https://stripe.com")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode state"})
		return
	}

	// Verify the state is not expired
	if statePayload.ExpiresAt.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "State expired"})
		return
	}

	//get user from db
	user, err := controllers.GetUserByID(statePayload.UserID)
	if err != nil {
		log.Printf("Failed to fetch user: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	// Set your Stripe secret key
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	// Exchange the authorization code for an access token
	params := &stripe.OAuthTokenParams{
		GrantType: stripe.String("authorization_code"),
		Code:      stripe.String(code),
		// Add these required parameters
		ClientSecret: stripe.String(stripe.Key),
	}

	// Make the OAuth token request
	token, err := oauth.New(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange code for token: " + err.Error()})
		return
	}

	log.Printf("Stripe token: %+v", token)

	accountID := token.StripeUserID
	customerID, err := controllers.GetOrCreateCustomerID(token.AccessToken, token.StripeUserID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch or create customer ID: " + err.Error()})
		return
	}

	// Save the connected account ID to the user's profile
	err = controllers.SaveStripeConnectedAccountID(statePayload.UserID, accountID, customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save connected account ID"})
		return
	}

	// Redirect the user back to the frontend
	redirectURL := os.Getenv("STRIPE_REDIRECT_URL")

	c.Redirect(http.StatusFound, redirectURL)
}
