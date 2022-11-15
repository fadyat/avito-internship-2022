package services

import (
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
)

type IHealthService interface {
	Ping() error
}

type IUserWalletService interface {
	CreateUserWallet(w dto.UserWallet) (uint64, error)
	GetAllUserWallets() ([]*models.UserWallet, error)
	GetUserWalletByID(id string) (*models.UserWallet, error)
}

type IOuterServiceService interface {
	CreateService(os dto.OuterService) (uint64, error)
	GetAllServices() ([]*models.OuterService, error)
	GetServiceByID(id string) (*models.OuterService, error)
}

type ITransactionService interface {
	CreateReplenishment(tr dto.Replenishment) (uint64, error)
	CreateWithdrawal(tr dto.Withdrawal) (uint64, error)
	CreateReservation(tr dto.Reservation) (uint64, error)
	CreateRelease(tr dto.Release) (uint64, error)
	GetAllTransactions() ([]*models.Transaction, error)
	GetTransactionByID(id string) (*models.Transaction, error)
}
