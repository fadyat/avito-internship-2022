package responses

// ErrorResp godoc
// @description: ErrorResp is an base error response for all endpoints.
type ErrorResp struct {
	Message     string `json:"message"`
	StatusCode  int    `json:"status_code"`
	Description string `json:"description"`
}
