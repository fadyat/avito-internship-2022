package services

import (
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
)

type IHealthService interface {
	Ping() error
}

type IUserWalletService interface {
	CreateUserWallet(w *dto.UserWallet) (uint64, error)
	GetAllUserWallets() ([]*models.UserWallet, error)
	GetUserWalletByID(id string) (*models.UserWallet, error)
}

type IOuterServiceService interface {
	CreateService(os dto.OuterService) (uint64, error)
	GetAllServices() ([]*models.OuterService, error)
	GetServiceByID(id string) (*models.OuterService, error)
}

type ITransactionService interface {
	CreateReplenishment(tr dto.Transaction) (uint64, error)
	CreateWithdrawal(tr dto.Transaction) (uint64, error)
	GetUserTransactions(userID string, page, perPage uint64, orderBy []string) ([]*models.Transaction, error)
	GetUserTransactionsCount(userID string) (uint64, error)
	iReservationService
}

type iReservationService interface {
	CreateReservation(tr dto.Reservation) (uint64, error)
	CreateRelease(tr dto.Reservation) (uint64, error)
	CancelReservation(tr dto.Reservation) (uint64, error)
	GetReservationsReport(tm dto.ReportTime) ([]*models.ReservationReport, error)
}
