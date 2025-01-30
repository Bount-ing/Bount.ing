package controllers

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/bount-ing/bount.ing/api/tools"
)

type OAuthState struct {
	State     string    `json:"state"`
	ExpiresAt time.Time `json:"expires_at"`
	UserID    uint      `json:"user_id"`
	OAuthHost string    `json:"oauth_host"`
}

var (
	keyMap   = make(map[string]string)
	keyMapMu sync.RWMutex
)

// GenerateOAuthState generates an encrypted OAuth state
func GenerateOAuthState(userID uint, oAuthHost string) (string, error) {
	// Generate random state string
	randomString, err := tools.RandomString(32)
	if err != nil {
		return "", err
	}

	// Create state object
	state := OAuthState{
		State:     randomString,
		ExpiresAt: time.Now().Add(10 * time.Minute),
		UserID:    userID,
		OAuthHost: oAuthHost,
	}

	// Convert to JSON
	statePayloadJSON, err := json.Marshal(state)
	if err != nil {
		return "", err
	}

	// Get encryption key
	key, err := getOrCreateKey(oAuthHost)
	if err != nil {
		return "", err
	}

	// Encrypt the state - EncryptAES now returns (string, error)
	encryptedState, err := tools.EncryptAES(statePayloadJSON, []byte(key))
	if err != nil {
		return "", err
	}

	return encryptedState, nil // Already base64 encoded
}

// DecryptAndReadOAuthState decrypts and parses the OAuth state
func DecryptAndReadOAuthState(encryptedState string, oAuthHost string) (*OAuthState, error) {
	// Get the key
	key, err := getOrCreateKey(oAuthHost)
	if err != nil {
		return nil, err
	}

	// Decrypt the data - DecryptAES now accepts string
	decryptedData, err := tools.DecryptAES(encryptedState, []byte(key))
	if err != nil {
		return nil, err
	}

	// Unmarshal into state struct
	var state OAuthState
	if err := json.Unmarshal(decryptedData, &state); err != nil {
		return nil, err
	}

	// Validate expiration
	if time.Now().After(state.ExpiresAt) {
		return nil, fmt.Errorf("oauth state has expired")
	}

	return &state, nil
}

// getOrCreateKey retrieves or creates an encryption key for a given OAuth host
func getOrCreateKey(oAuthHost string) (string, error) {
	keyMapMu.RLock()
	value, exists := keyMap[oAuthHost]
	keyMapMu.RUnlock()

	if exists {
		return value, nil
	}

	keyMapMu.Lock()
	defer keyMapMu.Unlock()

	// Double-check after acquiring write lock
	if value, exists := keyMap[oAuthHost]; exists {
		return value, nil
	}

	// Generate new key
	randomString, err := tools.RandomString(32)
	if err != nil {
		return "", err
	}

	keyMap[oAuthHost] = randomString
	return randomString, nil
}
