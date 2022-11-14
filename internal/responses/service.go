package responses

import "github.com/fadyat/avito-internship-2022/internal/models"

type ServiceCreated struct {
	ID string `json:"id"`
}

type Services struct {
	Services []*models.OuterService `json:"services"`
}
