package services

import (
	"github.com/fadyat/avito-internship-2022/internal/helpers"
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/responses"
)

func (s *TransactionService) CreateReservation(tr dto.Reservation) (uint64, error) {
	if err := s.v.Struct(tr); err != nil {
		return 0, &responses.ValidationErrResp{Message: err.Error()}
	}

	return s.r.CreateReservation(tr)
}

func (s *TransactionService) CreateRelease(tr dto.Reservation) (uint64, error) {
	if err := s.v.Struct(tr); err != nil {
		return 0, &responses.ValidationErrResp{Message: err.Error()}
	}

	return s.r.CreateRelease(tr)
}

func (s *TransactionService) CancelReservation(tr dto.Reservation) (uint64, error) {
	if err := s.v.Struct(tr); err != nil {
		return 0, &responses.ValidationErrResp{Message: err.Error()}
	}

	return s.r.CancelReservation(tr)
}

func (s *TransactionService) GetUserTransactions(userID string, page, perPage uint64, orderBy []string) ([]*models.Transaction, error) {
	uid, err := helpers.ValidateUint64(userID, "required,numeric,gte=1", s.v)
	if err != nil {
		return nil, err
	}

	return s.r.GetUserTransactions(uid, page, perPage, orderBy)
}

func (s *TransactionService) GetUserTransactionsCount(userID string) (uint64, error) {
	uid, err := helpers.ValidateUint64(userID, "required,numeric,gte=1", s.v)
	if err != nil {
		return 0, err
	}

	return s.r.GetUserTransactionsCount(uid)
}
