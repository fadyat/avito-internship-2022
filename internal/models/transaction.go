package models

import "time"

type TransactionType string

const (
	Replenishment TransactionType = "replenishment"
	Withdrawal    TransactionType = "withdrawal"
)

// Transaction godoc
// @description: Transaction is an operation between two wallets.
type Transaction struct {

	// @description: ID is a unique identifier of the transaction.
	// @example:     1
	ID uint64 `json:"id"`

	// @description: UserID is a unique identifier of the user, that owns this transaction.
	// @example:     1
	UserID uint64 `json:"user_id"`

	// @description: Amount is the amount of money, that was transferred.
	// @example:     100
	Amount uint64 `json:"amount"`

	// @description: Type is the type of the transaction.
	// @example:     Replenishment
	Type TransactionType `json:"type"`

	// @description: CreatedAt is the time, when the transaction was created.
	// @example:     2021-01-01T00:00:00Z
	CreatedAt time.Time `json:"created_at"`
}

type ReservationStatus string

const (
	Pending  ReservationStatus = "pending"
	Released ReservationStatus = "approved" // TODO: rename to "released"
	Canceled ReservationStatus = "rejected" // TODO: rename to canceled
)

type Reservation struct {

	// @description: ID is a unique identifier of the reservation.
	// @example:     1
	ID uint64 `json:"id"`

	// @description: UserID is a unique identifier of the user, that owns this reservation.
	// @example:     1
	UserID uint64 `json:"user_id"`

	// @description: ServiceID is a unique identifier of the service, that made this reservation.
	// @example:     1
	ServiceID uint64 `json:"service_id"`

	// @description: OrderID is a unique identifier of the order, that belongs to the service.
	// @example:     1
	OrderID uint64 `json:"order_id"`

	// @description: Amount is the amount of money, that was transferred.
	// @example:     100
	Amount uint64 `json:"amount"`

	// @description: Status is the status of the reservation.
	// @example:     Pending
	Status ReservationStatus `json:"status"`
}
