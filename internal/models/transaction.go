package models

type TransactionType uint64

// Transaction godoc
// @description: Transaction is an operation between two wallets.
type Transaction struct {
	// @description: ID is a unique identifier of the transaction.
	ID uint64 `json:"id"`

	// @description: UserID is a unique identifier of the user, that owns this transaction.
	UserID uint64 `json:"user_id"`

	// @description: ServiceID is a unique identifier of the service, that made this transaction.
	ServiceID uint64 `json:"service_id"`

	// @description: OrderID is a unique identifier of the order, that belongs to the service.
	OrderID uint64 `json:"order_id"`

	// @description: Amount is the amount of money, that was transferred.
	Amount uint64 `json:"amount"`
}
