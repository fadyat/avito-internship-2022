package services

import (
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/persistence"
	"github.com/fadyat/avito-internship-2022/internal/responses"
	"github.com/go-playground/validator/v10"
	"strconv"
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

func (s *UserWalletService) CreateUserWallet(w dto.UserWallet) (uint64, error) {
	if err := s.v.Struct(w); err != nil {
		return 0, &responses.ValidationErrResp{Message: err.Error()}
	}

	return s.r.CreateUserWallet(w)
}

func (s *UserWalletService) GetUserWalletByID(id string) (*models.UserWallet, error) {
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, &responses.ValidationErrResp{Message: "failed to parse id"}
	}

	if err = s.v.Var(uid, "required,numeric,gte=1"); err != nil {
		return nil, &responses.ValidationErrResp{Message: err.Error()}
	}

	return s.r.GetUserWalletByID(uid)
}

func (s *UserWalletService) GetAllWallets() ([]*models.UserWallet, error) {
	return s.r.GetAllWallets()
}
