package models

// OuterService godoc
// @description: OuterService is an outer microservice, that uses this microservice.
type OuterService struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`

	// @description: Link is a link to the implementation of the service.
	Link string `json:"link,omitempty"`
}
