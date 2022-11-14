package postgres

import (
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/jackc/pgx/v5"
)

type TransactionRepository struct {
	c *pgx.Conn
}

func NewTransactionRepository(c *pgx.Conn) *TransactionRepository {
	return &TransactionRepository{c: c}
}

func (t *TransactionRepository) CreateReplenishment(tr dto.Replenishment) (uint64, error) {
	// TODO implement me
	panic("implement me")
}

func (t *TransactionRepository) CreateWithdrawal(tr dto.Withdrawal) (uint64, error) {
	// TODO implement me
	panic("implement me")
}

func (t *TransactionRepository) CreateReservation(tr dto.Reservation) (uint64, error) {
	// TODO implement me
	panic("implement me")
}

func (t *TransactionRepository) CreateRelease(tr dto.Release) (uint64, error) {
	// TODO implement me
	panic("implement me")
}
