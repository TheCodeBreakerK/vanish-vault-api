package dto

import "time"

// HealthCheckResponseDto represents the structure of the health check response.
type HealthCheckResponseDto struct {
	Code    int       `json:"code"`
	Status  string    `json:"status"`
	TS      time.Time `json:"ts"`
	Service string    `json:"service"`
}
