package services

import (
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/persistence"
	"github.com/fadyat/avito-internship-2022/internal/responses"
	"github.com/go-playground/validator/v10"
	"strconv"
)

type OuterServiceService struct {
	r persistence.IOuterServiceRepository
	v *validator.Validate
}

func NewOuterServiceService(
	r persistence.IOuterServiceRepository,
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
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, &responses.ValidationErrResp{Message: "failed to parse id"}
	}

	if err = s.v.Var(uid, "required,numeric,gte=1"); err != nil {
		return nil, &responses.ValidationErrResp{Message: err.Error()}
	}

	return s.r.GetServiceByID(uid)
}

func (s *OuterServiceService) GetAllServices() ([]*models.OuterService, error) {
	return s.r.GetAllServices()
}
