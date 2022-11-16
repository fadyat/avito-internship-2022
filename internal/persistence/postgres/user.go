package postgres

import (
	"context"
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/jackc/pgx/v5"
)

type UserWalletRepo struct {
	c *pgx.Conn
}

func NewUserWalletRepo(c *pgx.Conn) *UserWalletRepo {
	return &UserWalletRepo{c: c}
}

func (u *UserWalletRepo) CreateUserWallet(w dto.UserWallet) (uint64, error) {
	id, err := u.createUserWallet(w)
	return id, recastError(err)
}

func (u *UserWalletRepo) createUserWallet(w dto.UserWallet) (uint64, error) {
	tx, err := u.c.Begin(context.Background())
	if err != nil {
		return 0, err
	}
	defer func() { _ = tx.Rollback(context.Background()) }()

	var id uint64
	q := "INSERT INTO user_wallets (user_id, balance) VALUES ($1, $2) RETURNING user_id"
	err = tx.QueryRow(context.Background(), q, w.UserID, 0).Scan(&id)
	if err != nil {
		return 0, err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return 0, err
	}

	return id, nil
}

func (u *UserWalletRepo) GetUserWalletByID(id uint64) (*models.UserWallet, error) {
	w, err := u.getUserWalletByID(id)
	return w, recastError(err)
}

func (u *UserWalletRepo) getUserWalletByID(id uint64) (*models.UserWallet, error) {
	tx, err := u.c.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback(context.Background()) }()

	var w models.UserWallet
	q := "SELECT * FROM user_wallets WHERE user_id = $1 LIMIT 1"
	err = tx.QueryRow(context.Background(), q, id).Scan(&w.UserID, &w.Balance)
	if err != nil {
		return nil, err
	}

	return &w, nil
}

func (u *UserWalletRepo) GetAllWallets() ([]*models.UserWallet, error) {
	ws, err := u.getAllWallets()
	return ws, recastError(err)
}

func (u *UserWalletRepo) getAllWallets() ([]*models.UserWallet, error) {
	tx, err := u.c.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback(context.Background()) }()

	var ws []*models.UserWallet
	q := "SELECT * FROM user_wallets"
	rows, err := tx.Query(context.Background(), q)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var w models.UserWallet
		_ = rows.Scan(&w.UserID, &w.Balance)
		ws = append(ws, &w)
	}

	return ws, nil
}
