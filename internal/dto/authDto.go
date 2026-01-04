// Package dto contains Data Transfer Objects (DTOs) used to
// standardize the input and output data across the API.
package dto

// LoginResponseDto represents the structure of the authentication URL response.
type LoginResponseDto struct {
	URL string `json:"url"`
}

// CallbackResponseDto represents the structure of the OAuth2 callback response.
type CallbackResponseDto struct {
	Token     string `json:"token"`
	TokenType string `json:"token_type"`
	ExpiryAt  int64  `json:"expiry_at"`
}

// UserInfoResponseDto holds the standardized user profile data from external auth providers.
type UserInfoResponseDto struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}
