package models

// Service godoc
// @description: Service is an outer microservice, that uses this microservice.
type Service struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`

	// @description: Link is a link to the implementation of the service.
	Link string `json:"link"`
}
