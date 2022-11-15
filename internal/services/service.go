package services

import (
	"github.com/fadyat/avito-internship-2022/internal/helpers"
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/persistence"
	"github.com/fadyat/avito-internship-2022/internal/responses"
	"github.com/go-playground/validator/v10"
)

type OuterServiceService struct {
	r persistence.IOuterServiceRepo
	v *validator.Validate
}

func NewOuterServiceService(
	r persistence.IOuterServiceRepo,
	v *validator.Validate,
) *OuterServiceService {
	return &OuterServiceService{r: r, v: v}
}

func (s *OuterServiceService) CreateService(os dto.OuterService) (uint64, error) {
	if err := s.v.Struct(os); err != nil {
		return 0, &responses.ValidationErrResp{Message: err.Error()}
	}

	return s.r.CreateService(os)
}

func (s *OuterServiceService) GetServiceByID(id string) (*models.OuterService, error) {
	uid, err := helpers.ValidateUint64(id, "required,numeric,gte=1", s.v)
	if err != nil {
		return nil, err
	}

	return s.r.GetServiceByID(uid)
}

func (s *OuterServiceService) GetAllServices() ([]*models.OuterService, error) {
	return s.r.GetAllServices()
}
