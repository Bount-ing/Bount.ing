package tools

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
)

// Generates a random verification for an ad
func RandomString(n uint) (string, error) {
	if n > 4096 {
		return "", errors.New("random strings are up to 4096 chars")
	}
	bytes := make([]byte, 4096)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	s := hex.EncodeToString(bytes)
	return s[:n], nil
}
