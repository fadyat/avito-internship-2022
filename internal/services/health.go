package services

import "github.com/fadyat/avito-internship-2022/internal/persistence"

type HealthService struct {
	r persistence.IHealthRepository
}

func NewHealthService(r persistence.IHealthRepository) *HealthService {
	return &HealthService{r: r}
}

func (h *HealthService) Ping() error {
	return h.r.Ping()
}
