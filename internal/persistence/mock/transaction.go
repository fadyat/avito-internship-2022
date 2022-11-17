package mock

import (
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
)

type TransactionRepo struct{}

func (t *TransactionRepo) CreateReplenishment(_ *dto.Transaction) (uint64, error) {
	return 1, nil
}

func (t *TransactionRepo) CreateWithdrawal(_ *dto.Transaction) (uint64, error) {
	return 1, nil
}

func (t *TransactionRepo) CreateReservation(_ *dto.Reservation) (uint64, error) {
	return 1, nil
}

func (t *TransactionRepo) CreateRelease(_ *dto.Reservation) (uint64, error) {
	return 1, nil
}

func (t *TransactionRepo) CancelReservation(_ *dto.Reservation) (uint64, error) {
	return 1, nil
}

func (t *TransactionRepo) GetUserTransactions(_, _, _ uint64, _ []string) ([]*models.Transaction, error) {
	return make([]*models.Transaction, 0), nil
}

func (t *TransactionRepo) GetUserTransactionsCount(_ uint64) (uint64, error) {
	return 1, nil
}

func (t *TransactionRepo) GetReservationsReport(_ *dto.ReportTime) ([]*models.ReservationReport, error) {
	return make([]*models.ReservationReport, 0), nil
}
