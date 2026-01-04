package service

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateRandomState creates a secure random string to be used as the OAuth2 state parameter.
func GenerateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
