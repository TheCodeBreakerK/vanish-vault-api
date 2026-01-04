// Package dto contains Data Transfer Objects (DTOs) used to
// standardize the input and output data across the API.
package dto

// LoginResponse represents the structure of the authentication URL response.
type LoginResponse struct {
	URL string `json:"url"`
}
