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
	// TODO implement me
	panic("implement me")
}

func (u *UserWalletRepository) GetUserWalletByID(id uint64) (*models.UserWallet, error) {
	// TODO implement me
	panic("implement me")
}

func (u *UserWalletRepository) GetAllWallets() ([]*models.UserWallet, error) {
	// TODO implement me
	panic("implement me")
}