package models

type TransactionType uint

const (
	Replenishment TransactionType = iota
	Withdrawal
	Reservation
	Release
)

// Transaction godoc
// @description: Transaction is an operation between two wallets.
type Transaction struct {
	// @description: ID is a unique identifier of the transaction.
	ID uint `json:"id"`

	// @description: UserID is a unique identifier of the user, that owns this transaction.
	UserID uint `json:"user_id"`

	// @description: ServiceID is a unique identifier of the service, that made this transaction.
	ServiceID uint `json:"service_id"`

	// @description: OrderID is a unique identifier of the order, that belongs to the service.
	OrderID uint `json:"order_id"`

	// @description: Amount is the amount of money, that was transferred.
	Amount uint `json:"amount"`

	// @description: Type is a type of the transaction.
	Type TransactionType `json:"type"`
}
