package services

import (
	"github.com/fadyat/avito-internship-2022/internal/helpers"
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/persistence"
	"github.com/fadyat/avito-internship-2022/internal/responses"
	"github.com/go-playground/validator/v10"
)

type UserWalletService struct {
	r persistence.IUserWalletRepo
	v *validator.Validate
}

func NewUserWalletService(
	r persistence.IUserWalletRepo,
	v *validator.Validate,
) *UserWalletService {
	return &UserWalletService{r: r, v: v}
}

func (s *UserWalletService) CreateUserWallet(w *dto.UserWallet) (uint64, error) {
	if err := s.v.Struct(w); err != nil {
		return 0, &responses.ValidationErrResp{Message: err.Error()}
	}

	return s.r.CreateUserWallet(w)
}

func (s *UserWalletService) GetUserWalletByID(id string) (*models.UserWallet, error) {
	uid, err := helpers.ValidateUint64(id, "required,numeric,gte=1", s.v)
	if err != nil {
		return nil, err
	}

	return s.r.GetUserWalletByID(uid)
}

func (s *UserWalletService) GetAllUserWallets() ([]*models.UserWallet, error) {
	return s.r.GetAllWallets()
}
