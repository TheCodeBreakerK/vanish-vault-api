package dto

import "time"

type HealthCheckResponse struct {
	Code    int       `json:"code"`
	Status  string    `json:"status"`
	TS      time.Time `json:"ts"`
	Service string    `json:"service"`
}
