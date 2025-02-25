package controllers

import (
	"github.com/bount-ing/bount.ing/api/db"
	"github.com/bount-ing/bount.ing/api/models"
)

func GetUserIdentities(userID uint) ([]models.Identity, error) {
	var identities []models.Identity

	//preload host
	if err := db.DB.Preload("Host").Where("user_id = ?", userID).Find(&identities).Error; err != nil {
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

func SaveStripeConnectedAccountID(stateUserID uint, stripeUserAccountID, stripeUserCustomerID string) error {

	// Check 1: stateUserID doesn't already have a github identity
	userIdentities, err := GetUserIdentities(stateUserID)
	if err != nil {
		return models.ErrUserIdentitiesNotFound
	}
	for _, identity := range userIdentities {
		if identity.Host.Address == "https://stripe.com" && identity.UserExternalID != stripeUserAccountID {
			return models.ErrUserIdentitiesStripeAccountMismatch
		} else if identity.Host.Address == "https://stripe.com" && identity.UserExternalID == stripeUserAccountID {
			// return ok
			return nil
		}
	}

	//fetch host id
	host, err := GetHostFromAddress("https://stripe.com")

	// Check 2: stripeUser.ID doesn't already have an identity
	hostIdentities, err := GetHostIdentitiesFromAddress("https://stripe.com")
	if err != nil {
		return models.ErrHostIdentitiesNotFound
	}
	for _, identity := range hostIdentities {
		if identity.UserExternalID == stripeUserAccountID && identity.UserID != stateUserID {
			return models.ErrStripeAccountAlreadyInUse
		}
	}

	// Create a new identity for the user
	newIdentity := &models.Identity{
		UserExternalID:          stripeUserAccountID,
		UserExternalSecondaryID: stripeUserCustomerID,
		UserID:                  stateUserID,
		HostID:                  host.ID, // Stripe host ID
	}

	err = CreateIdentity(newIdentity)
	if err != nil {
		return err
	}

	return nil
}
