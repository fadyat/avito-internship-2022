package postgres

import (
	"fmt"
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/persistence"
	"github.com/jmoiron/sqlx"
	"strings"
)

type TransactionRepo struct {
	c *sqlx.DB
}

func NewTransactionRepo(c *sqlx.DB) *TransactionRepo {
	return &TransactionRepo{c: c}
}

func (t *TransactionRepo) CreateReplenishment(tr dto.Transaction) (uint64, error) {
	id, err := t.createReplenishment(tr)
	return id, recastError(err)
}

func (t *TransactionRepo) createReplenishment(tr dto.Transaction) (uint64, error) {
	tx, err := t.c.Begin()
	if err != nil {
		return 0, err
	}
	defer func() { _ = tx.Rollback() }()

	if tr.Amount <= 0 {
		return 0, persistence.ErrNegativeAmount
	}

	wq := `UPDATE user_wallets SET balance = balance + $1 WHERE user_id = $2`
	_, err = tx.Exec(wq, tr.Amount, tr.UserID)
	if err != nil {
		return 0, err
	}

	var id uint64
	tq := `INSERT INTO transactions (user_id, amount, type) VALUES ($1, $2, $3) RETURNING id`
	err = tx.QueryRow(tq, tr.UserID, tr.Amount, models.Replenishment).Scan(&id)
	if err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return id, nil
}

func (t *TransactionRepo) CreateWithdrawal(tr dto.Transaction) (uint64, error) {
	id, err := t.createWithdrawal(tr)
	return id, recastError(err)
}

func (t *TransactionRepo) createWithdrawal(tr dto.Transaction) (uint64, error) {
	tx, err := t.c.Begin()
	if err != nil {
		return 0, err
	}
	defer func() { _ = tx.Rollback() }()

	if tr.Amount <= 0 {
		return 0, persistence.ErrNegativeAmount
	}

	var balance uint64
	b := `SELECT balance FROM user_wallets WHERE user_id = $1`
	err = tx.QueryRow(b, tr.UserID).Scan(&balance)
	if err != nil {
		return 0, err
	}

	if balance < tr.Amount {
		return 0, persistence.ErrInsufficientFunds
	}

	wq := `UPDATE user_wallets SET balance = balance - $1 WHERE user_id = $2`
	_, err = tx.Exec(wq, tr.Amount, tr.UserID)
	if err != nil {
		return 0, err
	}

	var id uint64
	tq := `INSERT INTO transactions (user_id, amount, type) VALUES ($1, $2, $3) RETURNING id`
	err = tx.QueryRow(tq, tr.UserID, tr.Amount, models.Withdrawal).Scan(&id)
	if err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return id, nil
}

func (t *TransactionRepo) GetUserTransactions(userID, page, perPage uint64, orderBy []string) ([]*models.Transaction, error) {
	ts, err := t.getUserTransactions(userID, page, perPage, orderBy)
	return ts, recastError(err)
}

func (t *TransactionRepo) getUserTransactions(userID, page, perPage uint64, orderBy []string) ([]*models.Transaction, error) {
	tx, err := t.c.Begin()
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback() }()

	q := fmt.Sprintf(
		"SELECT * FROM transactions WHERE user_id = $1 ORDER BY %s LIMIT $2 OFFSET $3",
		strings.Join(orderBy, ", "),
	)
	rows, err := tx.Query(q, userID, perPage, (page-1)*perPage)
	if err != nil {
		return nil, err
	}

	var ts []*models.Transaction
	for rows.Next() {
		var tr models.Transaction
		_ = rows.Scan(&tr.ID, &tr.UserID, &tr.Amount, &tr.Type, &tr.CreatedAt)
		ts = append(ts, &tr)
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return ts, nil
}

func (t *TransactionRepo) GetUserTransactionsCount(userID uint64) (uint64, error) {
	count, err := t.getUserTransactionsCount(userID)
	return count, recastError(err)
}

func (t *TransactionRepo) getUserTransactionsCount(userID uint64) (uint64, error) {
	tx, err := t.c.Begin()
	if err != nil {
		return 0, err
	}
	defer func() { _ = tx.Rollback() }()

	q := `SELECT COUNT(*) FROM transactions WHERE user_id = $1`
	var cnt uint64
	err = tx.QueryRow(q, userID).Scan(&cnt)
	if err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return cnt, nil
}
