package services

import (
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/persistence"
	"github.com/fadyat/avito-internship-2022/internal/responses"
	"github.com/go-playground/validator/v10"
)

type TransactionService struct {
	r persistence.ITransactionRepo
	v *validator.Validate
}

func NewTransactionService(
	r persistence.ITransactionRepo,
	v *validator.Validate,
) *TransactionService {
	return &TransactionService{r: r, v: v}
}

func (s *TransactionService) CreateReplenishment(tr dto.Transaction) (uint64, error) {
	if err := s.v.Struct(tr); err != nil {
		return 0, &responses.ValidationErrResp{Message: err.Error()}
	}

	return s.r.CreateReplenishment(tr)
}

func (s *TransactionService) CreateWithdrawal(tr dto.Transaction) (uint64, error) {
	if err := s.v.Struct(tr); err != nil {
		return 0, &responses.ValidationErrResp{Message: err.Error()}
	}

	return s.r.CreateWithdrawal(tr)
}
