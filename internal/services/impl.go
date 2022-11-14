package services

import "github.com/fadyat/avito-internship-2022/internal/persistence"

// todo: replace repository interface with service interface

type IHealthService interface {
	persistence.IHealthRepository
}

// todo: replace repository interface with service interface

type IUserWalletService interface {
	persistence.IUserWalletRepository
}

// todo: replace repository interface with service interface

type IOuterServiceService interface {
	persistence.IOuterServiceRepository
}

// todo: replace repository interface with service interface

type ITransactionService interface {
	persistence.ITransactionRepository
}
