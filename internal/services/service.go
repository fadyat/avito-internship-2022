package services

import (
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/persistence"
	"github.com/fadyat/avito-internship-2022/internal/validators"
	"strconv"
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

func (s *OuterServiceService) GetServiceByID(id string) (*models.OuterService, error) {
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}

	if err := validators.ValidateServiceID(uid); err != nil {
		return nil, err
	}

	return s.repo.GetServiceByID(uid)
}

func (s *OuterServiceService) GetAllServices() ([]*models.OuterService, error) {
	return s.repo.GetAllServices()
}
