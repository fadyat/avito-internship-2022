package responses

// ErrorResp godoc
// @description: ErrorResp is an base error response for all endpoints.
type ErrorResp struct {

	// @description: Message is a message of the error.
	// @example:     invalid request
	Message string `json:"message"`

	// @description: Description is a description of the error.
	// @example:     invalid request with id=1
	Description string `json:"description"`
}

// ValidationErrResp godoc
// @description: ValidationErrResp is an error response for validation errors.
type ValidationErrResp struct {

	// @description: Message is a message of the error.
	// @example:     validation failed
	Message string `json:"message"`
	error
}

func (e *ValidationErrResp) Error() string {
	return e.Message
}
