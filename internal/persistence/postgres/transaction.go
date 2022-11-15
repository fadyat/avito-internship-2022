package postgres

import (
	"context"
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/persistence"
	"github.com/jackc/pgx/v5"
)

type TransactionRepo struct {
	c *pgx.Conn
}

func NewTransactionRepo(c *pgx.Conn) *TransactionRepo {
	return &TransactionRepo{c: c}
}

func (t *TransactionRepo) CreateReplenishment(tr dto.Transaction) (uint64, error) {
	id, err := t.createReplenishment(tr)
	return id, recastError(err)
}

func (t *TransactionRepo) createReplenishment(tr dto.Transaction) (uint64, error) {
	tx, err := t.c.Begin(context.Background())
	defer func() { _ = tx.Rollback(context.Background()) }()
	if err != nil {
		return 0, err
	}

	// extra validation if validation is not done on the service level
	if tr.Amount <= 0 {
		return 0, persistence.ErrNegativeAmount
	}

	wq := `UPDATE user_wallets SET balance = balance + $1 WHERE user_id = $2`
	_, err = tx.Exec(context.Background(), wq, tr.Amount, tr.UserID)
	if err != nil {
		return 0, err
	}

	var id uint64
	tq := `INSERT INTO transactions (user_id, amount, type) VALUES ($1, $2, $3) RETURNING id`
	err = tx.QueryRow(context.Background(), tq, tr.UserID, tr.Amount, models.Replenishment).Scan(&id)
	if err != nil {
		return 0, err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return 0, err
	}

	return id, nil
}

func (t *TransactionRepo) CreateWithdrawal(tr dto.Transaction) (uint64, error) {
	id, err := t.createWithdrawal(tr)
	return id, recastError(err)
}

func (t *TransactionRepo) createWithdrawal(tr dto.Transaction) (uint64, error) {
	tx, err := t.c.Begin(context.Background())
	defer func() { _ = tx.Rollback(context.Background()) }()
	if err != nil {
		return 0, err
	}

	// extra validation if validation is not done on the service level
	if tr.Amount <= 0 {
		return 0, persistence.ErrNegativeAmount
	}

	// todo: add a check that the user has enough money
	wq := `UPDATE user_wallets SET balance = balance - $1 WHERE user_id = $2`
	_, err = tx.Exec(context.Background(), wq, tr.Amount, tr.UserID)
	if err != nil {
		return 0, err
	}

	var id uint64
	tq := `INSERT INTO transactions (user_id, amount, type) VALUES ($1, $2, $3) RETURNING id`
	err = tx.QueryRow(context.Background(), tq, tr.UserID, tr.Amount, models.Withdrawal).Scan(&id)
	if err != nil {
		return 0, err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return 0, err
	}

	return id, nil
}
