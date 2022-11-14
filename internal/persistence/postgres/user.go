package postgres

import (
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/jackc/pgx/v5"
)

type UserWalletRepository struct {
	c *pgx.Conn
}

func NewUserWalletRepository(c *pgx.Conn) *UserWalletRepository {
	return &UserWalletRepository{c: c}
}

func (u *UserWalletRepository) CreateUserWallet(w dto.UserWallet) (uint64, error) {
	id, err := u.createUserWallet(w)
	return id, recastError(err)
}

func (u *UserWalletRepository) createUserWallet(w dto.UserWallet) (uint64, error) {
	// TODO implement me
	panic("implement me")
}

func (u *UserWalletRepository) GetUserWalletByID(id uint64) (*models.UserWallet, error) {
	w, err := u.getUserWalletByID(id)
	return w, recastError(err)
}

func (u *UserWalletRepository) getUserWalletByID(id uint64) (*models.UserWallet, error) {
	// TODO implement me
	panic("implement me")
}

func (u *UserWalletRepository) GetAllWallets() ([]*models.UserWallet, error) {
	ws, err := u.getAllWallets()
	return ws, recastError(err)
}

func (u *UserWalletRepository) getAllWallets() ([]*models.UserWallet, error) {
	// TODO implement me
	panic("implement me")
}
