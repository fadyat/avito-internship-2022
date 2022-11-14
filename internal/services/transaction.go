package services

import (
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/persistence"
)

type TransactionService struct {
	r persistence.ITransactionRepository
}

func NewTransactionService(r persistence.ITransactionRepository) *TransactionService {
	return &TransactionService{r: r}
}

func (t TransactionService) CreateReplenishment(tr dto.Replenishment) (uint64, error) {
	// TODO implement me
	panic("implement me")
}

func (t TransactionService) CreateWithdrawal(tr dto.Withdrawal) (uint64, error) {
	// TODO implement me
	panic("implement me")
}

func (t TransactionService) CreateReservation(tr dto.Reservation) (uint64, error) {
	// TODO implement me
	panic("implement me")
}

func (t TransactionService) CreateRelease(tr dto.Release) (uint64, error) {
	// TODO implement me
	panic("implement me")
}
