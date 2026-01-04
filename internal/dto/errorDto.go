package dto

// ErrorResponseDto defines the standard structure for all API error messages.
type ErrorResponseDto struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Status  string `json:"status"`
}
