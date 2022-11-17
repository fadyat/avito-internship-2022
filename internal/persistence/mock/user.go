package mock

import (
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
)

type UserWalletRepo struct{}

func (m *UserWalletRepo) CreateUserWallet(w *dto.UserWallet) (uint64, error) {
	return w.UserID, nil
}

func (m *UserWalletRepo) GetUserWalletByID(id uint64) (*models.UserWallet, error) {
	return &models.UserWallet{UserID: id, Balance: 0}, nil
}

func (m *UserWalletRepo) GetAllWallets() ([]*models.UserWallet, error) {
	return []*models.UserWallet{{UserID: 1, Balance: 0}}, nil
}
