package services

import (
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/persistence"
)

// todo: replace repository interface with service interface

type IHealthService interface {
	persistence.IHealthRepository
}

// todo: replace repository interface with service interface

type IUserWalletService interface {
	CreateUserWallet(w dto.UserWallet) (uint64, error)
	GetAllWallets() ([]*models.UserWallet, error)
	GetUserWalletByID(id string) (*models.UserWallet, error)
}

type IOuterServiceService interface {
	CreateService(os dto.OuterService) (uint64, error)
	GetAllServices() ([]*models.OuterService, error)
	GetServiceByID(id string) (*models.OuterService, error)
}

// todo: replace repository interface with service interface

type ITransactionService interface {
	persistence.ITransactionRepository
}
