package mock

import (
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
)

type OuterServiceRepo struct{}

func (m *OuterServiceRepo) CreateService(_ dto.OuterService) (uint64, error) {
	return 1, nil
}

func (m *OuterServiceRepo) GetServiceByID(id uint64) (*models.OuterService, error) {
	return &models.OuterService{ID: id, Name: "aboba", URL: "https://aboba.com"}, nil
}

func (m *OuterServiceRepo) GetAllServices() ([]*models.OuterService, error) {
	return []*models.OuterService{{ID: 1, Name: "aboba", URL: "https://aboba.com"}}, nil
}
