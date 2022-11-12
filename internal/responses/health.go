package responses

type HealthSuccess struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type HealthError struct {
	Message     string `json:"message"`
	StatusCode  int    `json:"status_code"`
	Description string `json:"description"`
}
