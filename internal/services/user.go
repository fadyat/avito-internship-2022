package services

import (
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/persistence"
)

type UserWalletService struct {
	repo persistence.IUserWalletRepository
}

func NewUserWalletService(repo persistence.IUserWalletRepository) *UserWalletService {
	return &UserWalletService{repo: repo}
}

func (u UserWalletService) CreateUserWallet(w dto.UserWallet) (uint64, error) {
	// TODO implement me
	panic("implement me")
}

func (u UserWalletService) GetUserWalletByID(id uint64) (*models.UserWallet, error) {
	// TODO implement me
	panic("implement me")
}

func (u UserWalletService) GetAllWallets() ([]*models.UserWallet, error) {
	// TODO implement me
	panic("implement me")
}
