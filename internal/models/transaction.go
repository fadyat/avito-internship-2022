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
	Pending   ReservationStatus = "pending"
	Released  ReservationStatus = "approved" // todo: rename to "released"
	Cancelled ReservationStatus = "cancelled"
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

	// @description: CreatedAt is the time, when the reservation was created.
	// @example:     2021-01-01T00:00:00Z
	CreatedAt time.Time `json:"created_at"`

	// @description: UpdatedAt is the time, when the reservation was updated.
	// @example:     2021-01-01T00:00:00Z
	UpdatedAt time.Time `json:"updated_at"`
}

// ReservationReport is a report about the reservation.
// @description: ReservationReport is a report about the reservation
type ReservationReport struct {

	// @description: ServiceID is a unique identifier of the service, that made this reservation.
	// @example:     1
	ServiceID uint64 `json:"service_id"`

	// @description: OrderID is a unique identifier of the order, that belongs to the service.
	// @example:     1
	OrderID uint64 `json:"order_id"`

	// @description: Amount is the sum of all accepted reservations from all users.
	// @example:     100
	Amount uint64 `json:"amount"`

	// @description: Count is the count of all reservations from all users.
	// @example:     100
	Count uint64 `json:"count"`

	// @description: Status is the status of the reservation.
	// @example:     Pending
	Status ReservationStatus `json:"status"`
}
