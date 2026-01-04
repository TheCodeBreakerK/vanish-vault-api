// Package service provides various services for the application.
package service

import (
	"time"

	"github.com/TheCodeBreakerK/vanish-vault-api/configs"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// GenerateToken creates a JWT token for the given user ID with an expiration time defined in the config.
func GenerateToken(userID uuid.UUID, cfg *configs.Conf) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID.String(),
		"exp": time.Now().Add(time.Hour * time.Duration(cfg.JWTExpirationHours)).Unix(),
		"iat": time.Now().Unix(),
		"iss": "vanish-vault-api",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := []byte(cfg.JWTSecret)

	return token.SignedString(secret)
}
