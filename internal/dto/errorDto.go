package dto

// ErrorResponse defines the standard structure for all API error messages.
type ErrorResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}
