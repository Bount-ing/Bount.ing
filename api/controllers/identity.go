package controllers

import (
	"fmt"

	"github.com/bount-ing/bount.ing/api/db"
	"github.com/bount-ing/bount.ing/api/models"
)

func GetUserIdentities(userID uint) ([]models.Identity, error) {
	var identities []models.Identity

	if err := db.DB.Where("user_id = ?", userID).Find(&identities).Error; err != nil {
		return identities, err
	}

	return identities, nil
}
func GetHostIdentitiesFromAddress(hostAddress string) ([]models.Identity, error) {
	var identities []models.Identity

	if err := db.DB.
		Joins("JOIN hosts ON hosts.id = identities.host_id").
		Where("hosts.address = ?", hostAddress).
		Find(&identities).Error; err != nil {
		return identities, err
	}

	return identities, nil
}

func CreateIdentity(identity *models.Identity) error {
	if err := db.DB.Create(identity).Error; err != nil {
		return err
	}

	return nil
}

func SaveStripeConnectedAccountID(stateUserID uint, stripeUserID string) error {

	// Check 1: stateUserID doesn't already have a github identity
	userIdentities, err := GetUserIdentities(stateUserID)
	if err != nil {
		return fmt.Errorf("failed to get user identities: %w", err)
	}
	for _, identity := range userIdentities {
		if identity.Host.Address == "https://stripe.com" && identity.UserExternalID != stripeUserID {
			return fmt.Errorf("user already has another Stripe identity associated")
		} else if identity.Host.Address == "https://stripe.com" && identity.UserExternalID == stripeUserID {
			// return ok
			return nil
		}
	}

	//fetch host id
	host, err := GetHostFromAddress("https://stripe.com")

	// Check 2: stripeUser.ID doesn't already have an identity
	hostIdentities, err := GetHostIdentitiesFromAddress("https://stripe.com")
	if err != nil {
		return fmt.Errorf("failed to get host identities: %w", err)
	}
	for _, identity := range hostIdentities {
		if identity.UserExternalID == stripeUserID && identity.UserID != stateUserID {
			return fmt.Errorf("Stripe identity already has another user associated")
		}
	}

	// Create a new identity for the user
	newIdentity := &models.Identity{
		UserExternalID: stripeUserID,
		UserID:         stateUserID,
		HostID:         host.ID, // Stripe host ID
	}

	err = CreateIdentity(newIdentity)
	if err != nil {
		return fmt.Errorf("failed to create identity: %w", err)
	}

	return nil
}
