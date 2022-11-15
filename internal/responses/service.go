package responses

import "github.com/fadyat/avito-internship-2022/internal/models"

// ServiceCreated godoc
// @description: ServiceCreated is a response for service creation.
type ServiceCreated struct {

	// @description: ID is given unique identifier of the service.
	// @example:     1
	ID uint64 `json:"id"`
}

// Services godoc
// @description: Services is a response for getting all services.
type Services struct {

	// @description: Services is a list of services.
	// @example:     [{"id":1,"name":"aboba-service","url":"http://aboba-service.com"}]
	Services []*models.OuterService `json:"services"`
}
