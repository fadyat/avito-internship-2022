package persistence

import (
	"errors"
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
)

var (
	ErrNotFound            = errors.New("not found")
	ErrUniqueViolation     = errors.New("unique violation")
	ErrForeignKeyViolation = errors.New("foreign key violation")
	ErrNegativeAmount      = errors.New("negative amount")
	ErrInvalidColumn       = errors.New("invalid column")
	ErrInsufficientFunds   = errors.New("insufficient funds")
)

type IHealthRepo interface {
	Ping() error
}

type IUserWalletRepo interface {
	CreateUserWallet(w *dto.UserWallet) (uint64, error)
	GetUserWalletByID(id uint64) (*models.UserWallet, error)
	GetAllWallets() ([]*models.UserWallet, error)
}

type ITransactionRepo interface {
	CreateReplenishment(tr dto.Transaction) (uint64, error)
	CreateWithdrawal(tr dto.Transaction) (uint64, error)
	GetUserTransactions(userID, page, perPage uint64, orderBy []string) ([]*models.Transaction, error)
	GetUserTransactionsCount(userID uint64) (uint64, error)
	iReservationRepo
}

type iReservationRepo interface {
	CreateReservation(tr dto.Reservation) (uint64, error)
	CreateRelease(tr dto.Reservation) (uint64, error)
	CancelReservation(tr dto.Reservation) (uint64, error)
	GetReservationsReport(tm dto.ReportTime) ([]*models.ReservationReport, error)
}

type IOuterServiceRepo interface {
	CreateService(os *dto.OuterService) (uint64, error)
	GetServiceByID(id uint64) (*models.OuterService, error)
	GetAllServices() ([]*models.OuterService, error)
}
