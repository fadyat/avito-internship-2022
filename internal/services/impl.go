package services

import (
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/persistence"
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

// todo: replace repository interface with service interface

type ITransactionService interface {
	persistence.ITransactionRepo
}
