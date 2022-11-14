package services

import "github.com/fadyat/avito-internship-2022/internal/persistence"

type HealthService struct {
	repo persistence.IHealthRepository
}

func NewHealthService(repo persistence.IHealthRepository) *HealthService {
	return &HealthService{repo: repo}
}

func (h *HealthService) Ping() error {
	return h.repo.Ping()
}
