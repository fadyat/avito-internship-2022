package services

import (
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/persistence"
	"github.com/fadyat/avito-internship-2022/internal/validators"
)

type OuterServiceService struct {
	repo persistence.IOuterServiceRepository
}

func NewOuterServiceService(repo persistence.IOuterServiceRepository) *OuterServiceService {
	return &OuterServiceService{repo: repo}
}

func (s *OuterServiceService) CreateService(os dto.OuterService) (uint64, error) {
	if err := validators.ValidateServiceRequest(os); err != nil {
		return 0, err
	}

	return s.repo.CreateService(os)
}

func (s *OuterServiceService) GetServiceByID(id uint64) (*models.OuterService, error) {
	// todo: add parsing id
	if err := validators.ValidateServiceID(id); err != nil {
		return nil, err
	}

	return s.repo.GetServiceByID(id)
}

func (s *OuterServiceService) GetAllServices() ([]*models.OuterService, error) {
	return s.repo.GetAllServices()
}
