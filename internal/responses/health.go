package responses

// HealthSuccess godoc
// @description: HealthSuccess is valid response for health check.
type HealthSuccess struct {

	// @description: Message is a success message.
	// @example:     OK
	Message string `json:"message"`
}
