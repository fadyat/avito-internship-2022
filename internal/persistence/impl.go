package persistence

import (
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
)

type IHealthRepository interface {
	// Ping godoc
	// @description: Ping checks if the database connection is working.
	Ping() error
}

type IUserWalletRepository interface {
	// CreateUserWallet godoc
	// @description: CreateUserWallet creates a new wallet for a user, user_id is passed from outer service
	CreateUserWallet(w dto.UserWallet) (uint64, error)

	// GetUserWalletByID godoc
	// @description: GetUserWalletByID returns a wallet of a user by its id
	GetUserWalletByID(id uint64) (*models.UserWallet, error)

	// GetAllWallets godoc
	// @description: GetAllWallets returns all wallets of all users
	GetAllWallets() ([]*models.UserWallet, error)
}

type ITransactionRepository interface {
	// CreateReplenishment godoc
	// @description: CreateReplenishment creates a new replenishment transaction
	CreateReplenishment(tr dto.Replenishment) (uint64, error)

	// CreateWithdrawal godoc
	// @description: CreateWithdrawal creates a new withdrawal transaction
	CreateWithdrawal(tr dto.Withdrawal) (uint64, error)

	// CreateReservation godoc
	// @description: CreateReservation creates a new reservation transaction
	CreateReservation(tr dto.Reservation) (uint64, error)

	// CreateRelease godoc
	// @description: CreateRelease creates a new release transaction
	CreateRelease(tr dto.Release) (uint64, error)
}

type IOuterServiceRepository interface {
	// CreateService godoc
	// @description: CreateService creates a new service, that uses this microservice
	CreateService(os dto.OuterService) (uint64, error)

	// GetServiceByID godoc
	// @description: GetServiceByID returns a service by its id
	GetServiceByID(id uint64) (*models.OuterService, error)

	// GetAllServices godoc
	// @description: GetAllServices returns all services, that uses this microservice
	GetAllServices() ([]*models.OuterService, error)
}
