package responses

// HealthSuccess godoc
// @description: HealthSuccess is valid response for health check.
type HealthSuccess struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}
