package models

import "errors"

var (
	// Database Errors
	ErrDatabaseConnection     = errors.New("api_error_0: failed to establish database connection")
	ErrDatabaseCheckFailed    = errors.New("api_error_30: failed to check if database exists")
	ErrDatabaseCreationFailed = errors.New("api_error_31: failed to create database")

	ErrUserNotFound  = errors.New("api_error_1: user not found")
	ErrIssueNotFound = errors.New("api_error_2: issue not found")

	ErrBountyNotFound     = errors.New("api_error_3: bounty not found")
	ErrBountyUpdateFailed = errors.New("api_error_4: failed to update bounty")

	ErrClaimNotFound       = errors.New("api_error_5: claim not found")
	ErrClaimCreationFailed = errors.New("api_error_6: failed to create claim")

	ErrClaimCheckCreationFailed = errors.New("api_error_7: failed to create claim check")
	ErrClaimCheckNotFound       = errors.New("api_error_8: claim check not found")
	ErrClaimCheckUpdateFailed   = errors.New("api_error_9: failed to update claim check")

	ErrSponsorClaimCheckCreationFailed = errors.New("api_error_10: failed to create sponsor claim check")
	ErrSystemClaimCheckCreationFailed  = errors.New("api_error_11: failed to create system claim check")

	ErrIdentityCreationFailed = errors.New("api_error_39: failed to create identity")
	ErrUserIdentitiesNotFound = errors.New("api_error_12: user identities not found")
	ErrHostIdentitiesNotFound = errors.New("api_error_13: host identities not found")

	ErrLegalEntityNotFound = errors.New("api_error_42: legal entity not found")

	// HTTP
	ErrHTTPRequestCreationFailed = errors.New("api_error_32: failed to create HTTP request")
	ErrHTTPRequestSendFailed     = errors.New("api_error_33: failed to send HTTP request")

	// Data Validation
	ErrIOReadFailed         = errors.New("api_error_34: failed to read response body")
	ErrJSONUnmarshalFailed  = errors.New("api_error_36: failed to unmarshal JSON")
	ErrInvalidSigningMethod = errors.New("api_error_40: invalid signing method")
	ErrCipherTextTooShort   = errors.New("api_error_41: ciphertext too short")

	ErrDateRangeStartAfterEnd        = errors.New("api_error_14: start date must be before end date")
	ErrBountyAmountCalculationFailed = errors.New("api_error_15: failed to calculate bounty amount")
	ErrSponsorRejectedClaimCheck     = errors.New("api_error_16: claim check was rejected by sponsor")

	ErrBountyAlreadyClosed = errors.New("api_error_43: bounty is already closed")

	// Permissions
	ErrOAuthStateExpired    = errors.New("api_error_17: oauth state expired")
	ErrApproveClaimNotOwned = errors.New("api_error_18: cannot approve a claim that is not owned by the user")

	// Requirements
	ErrClaimWithoutLegalEntity   = errors.New("api_error_19: user must have a tax informations to claim bounties")
	ErrInvoiceWithoutLegalEntity = errors.New("api_error_20: user must have a tax informations to ask for an invoice")

	// Third Party //
	// // GitHub
	ErrGitHubIssueNotFound           = errors.New("api_error_21: github issue not found")
	ErrGitHubInvalidURL              = errors.New("api_error_22: invalid github issue URL")
	ErrGitHubAvatarNotFound          = errors.New("api_error_23: github avatar not found")
	ErrGitHubTokenVerificationFailed = errors.New("api_error_35: failed to verify github token")
	ErrGitHubDuplicateIdentity       = errors.New("api_error_37: user already has a github identity")
	ErrGitHubAlreadyInUse            = errors.New("api_error_38: github identity already in use")
	// // Stripe
	ErrUserIdentitiesStripeAccountMismatch = errors.New("api_error_24: stripe account mismatch")
	ErrStripeAccountAlreadyInUse           = errors.New("api_error_25: stripe account already in use")
	ErrNoPaymentMethods                    = errors.New("api_error_26: no payment methods found for customer")
	ErrPaymentIntent                       = errors.New("api_error_27: failed to process payment intent")
	ErrStripeFailedToCreateCustomer        = errors.New("api_error_28: failed to create stripe customer")
	ErrStripeFailedToVerifyCustomer        = errors.New("api_error_29: failed to verify stripe customer")
)
