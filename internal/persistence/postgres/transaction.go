package postgres

import (
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/jackc/pgx/v5"
)

type TransactionRepo struct {
	c *pgx.Conn
}

func NewTransactionRepo(c *pgx.Conn) *TransactionRepo {
	return &TransactionRepo{c: c}
}

func (t *TransactionRepo) CreateReplenishment(tr dto.Replenishment) (uint64, error) {
	// TODO implement me
	panic("implement me")
}

func (t *TransactionRepo) CreateWithdrawal(tr dto.Withdrawal) (uint64, error) {
	// TODO implement me
	panic("implement me")
}

func (t *TransactionRepo) CreateReservation(tr dto.Reservation) (uint64, error) {
	// TODO implement me
	panic("implement me")
}

func (t *TransactionRepo) CreateRelease(tr dto.Release) (uint64, error) {
	// TODO implement me
	panic("implement me")
}
