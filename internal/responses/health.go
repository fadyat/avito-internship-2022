package responses

// HealthSuccess godoc
// @description: HealthSuccess is valid response for health check.
type HealthSuccess struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// HealthError godoc
// @description: HealthError is an error response for health check.
type HealthError struct {
	Message     string `json:"message"`
	StatusCode  int    `json:"status_code"`
	Description string `json:"description"`
}
