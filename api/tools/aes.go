package tools

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// EncryptAES encrypts data using AES-GCM and returns base64 encoded string
func EncryptAES(plaintext []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, 12) // 12 bytes for GCM nonce
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nil, nonce, plaintext, nil)
	finalData := append(nonce, ciphertext...) // Prepend nonce

	return base64.URLEncoding.EncodeToString(finalData), nil
}

// DecryptAES decrypts base64 encoded string using AES-GCM
func DecryptAES(encryptedStr string, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	encryptedData, err := base64.URLEncoding.DecodeString(encryptedStr)
	if err != nil {
		return nil, err
	}

	if len(encryptedData) < 12 { // 12 is nonce size
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce := encryptedData[:12]
	ciphertext := encryptedData[12:]

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
