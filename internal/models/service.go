package models

// OuterService godoc
// @description: OuterService is an outer microservice, that uses this microservice.
type OuterService struct {

	// @description: ID is given unique identifier of the service.
	// @example:     1
	ID uint64 `json:"id"`

	// @description: Name is a name of the service.
	// @example:     aboba-service
	Name string `json:"name"`

	// @description: URL is a link to the implementation of the service.
	// @example:     https://aboba-service.com
	URL string `json:"url,omitempty"`
}
