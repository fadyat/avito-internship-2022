package postgres

import (
	"context"
	"fmt"
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/persistence"
	"strings"
)

func (t *TransactionRepo) CreateReservation(tr dto.Reservation) (uint64, error) {
	id, err := t.createReservation(tr)
	return id, recastError(err)
}

func (t *TransactionRepo) createReservation(tr dto.Reservation) (uint64, error) {
	tx, err := t.c.Begin(context.Background())
	defer func() { _ = tx.Rollback(context.Background()) }()
	if err != nil {
		return 0, err
	}

	// extra validation if validation is not done on the service level
	if tr.Amount <= 0 {
		return 0, persistence.ErrNegativeAmount
	}

	var id uint64
	rq := `INSERT INTO reservations (user_id, service_id, order_id, amount, status)
		   VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = tx.QueryRow(context.Background(), rq, tr.UserID, tr.ServiceID, tr.OrderID, tr.Amount, models.Pending).Scan(&id)
	if err != nil {
		return 0, err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return 0, err
	}

	return id, nil
}

func (t *TransactionRepo) CreateRelease(tr dto.Reservation) (uint64, error) {
	id, err := t.createRelease(tr)
	return id, recastError(err)
}

func (t *TransactionRepo) createRelease(tr dto.Reservation) (uint64, error) {
	// TODO implement me
	panic("implement me")
}

func (t *TransactionRepo) CancelReservation(tr dto.Reservation) (uint64, error) {
	id, err := t.cancelReservation(tr)
	return id, recastError(err)
}

func (t *TransactionRepo) cancelReservation(tr dto.Reservation) (uint64, error) {
	// TODO implement me
	panic("implement me")
}

func (t *TransactionRepo) GetUserTransactions(userID, page, perPage uint64, orderBy []string) ([]*models.Transaction, error) {
	ts, err := t.getUserTransactions(userID, page, perPage, orderBy)
	return ts, recastError(err)
}

func (t *TransactionRepo) getUserTransactions(userID, page, perPage uint64, orderBy []string) ([]*models.Transaction, error) {
	tx, err := t.c.Begin(context.Background())
	defer func() { _ = tx.Rollback(context.Background()) }()
	if err != nil {
		return nil, err
	}

	q := fmt.Sprintf(
		"SELECT * FROM transactions WHERE user_id = $1 ORDER BY %s LIMIT $2 OFFSET $3",
		strings.Join(orderBy, ", "),
	)
	rows, err := tx.Query(context.Background(), q, userID, perPage, (page-1)*perPage)
	if err != nil {
		return nil, err
	}

	var ts []*models.Transaction
	for rows.Next() {
		var tr models.Transaction
		_ = rows.Scan(&tr.ID, &tr.UserID, &tr.Amount, &tr.Type, &tr.CreatedAt)
		ts = append(ts, &tr)
	}

	if err := tx.Commit(context.Background()); err != nil {
		return nil, err
	}

	return ts, nil
}

func (t *TransactionRepo) GetUserTransactionsCount(userID uint64) (uint64, error) {
	count, err := t.getUserTransactionsCount(userID)
	return count, recastError(err)
}

func (t *TransactionRepo) getUserTransactionsCount(userID uint64) (uint64, error) {
	tx, err := t.c.Begin(context.Background())
	defer func() { _ = tx.Rollback(context.Background()) }()
	if err != nil {
		return 0, err
	}

	q := `SELECT COUNT(*) FROM transactions WHERE user_id = $1`
	var cnt uint64
	err = tx.QueryRow(context.Background(), q, userID).Scan(&cnt)
	if err != nil {
		return 0, err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return 0, err
	}

	return cnt, nil
}
