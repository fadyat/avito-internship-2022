package dto

// OuterService godoc
// @description: dto.OuterService is dto for models.OuterService.
type OuterService struct {

	// @description: Name is a name of the service.
	// @example:     aboba-service
	Name string `json:"name" validate:"required,min=1,max=255"`

	// @description: URL is a link to the implementation of the service.
	// @example:     http://aboba-service.com
	URL string `json:"url" validate:"omitempty,url,max=255"`
}
