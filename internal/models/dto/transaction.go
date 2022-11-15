package dto

// Replenishment godoc
// @description: Replenishment is a request to replenish a user's wallet.
type Replenishment struct {

	// @description: UserID is a unique identifier of the user, that owns this transaction.
	// @example:     1
	UserID uint64 `json:"user_id" validate:"required,gte=1,numeric"`

	// @description: Amount is the amount of money, that was transferred.
	// @example:     100
	Amount uint64 `json:"amount" validate:"required,gte=1,lte=1000000,numeric"`
}

// Withdrawal godoc
// @description: Withdrawal is a request to withdraw money from the user's wallet.
type Withdrawal struct {

	// @description: UserID is a unique identifier of the user, that owns this transaction.
	// @example:     1
	UserID uint64 `json:"user_id" validate:"required,gte=1,numeric"`

	// @description: Amount is the amount of money, that was transferred.
	// @example:     100
	Amount uint64 `json:"amount" validate:"required,gte=1,lte=1000000,numeric"`
}

// Reservation godoc
// @description: Reservation is a request to reserve money for a future transaction.
type Reservation struct {

	// @description: UserID is a unique identifier of the user, that owns this transaction.
	// @example:     1
	UserID uint64 `json:"user_id" validate:"required,gte=1,numeric"`

	// @description: ServiceID is a unique identifier of the service, that made this transaction.
	// @example:     1
	ServiceID uint64 `json:"service_id" validate:"required,gte=1,numeric"`

	// @description: OrderID is a unique identifier of the order, that belongs to the service.
	// @example:     1
	OrderID uint64 `json:"order_id" validate:"required,gte=1,numeric"`

	// @description: Amount is the amount of money, that was transferred.
	// @example:     100
	Amount uint64 `json:"amount" validate:"required,gte=1,lte=1000000,numeric"`
}

// Release godoc
// @description: Release is a request to release a reservation.
type Release struct {

	// @description: UserID is a unique identifier of the user, that owns this transaction.
	// @example:     1
	UserID uint64 `json:"user_id" validate:"required,gte=1,numeric"`

	// @description: ServiceID is a unique identifier of the service, that made this transaction.
	// @example:     1
	ServiceID uint64 `json:"service_id" validate:"required,gte=1,numeric"`

	// @description: OrderID is a unique identifier of the order, that belongs to the service.
	// @example:     1
	OrderID uint64 `json:"order_id" validate:"required,gte=1,numeric"`

	// @description: Amount is the amount of money, that was transferred.
	// @example:     100
	Amount uint64 `json:"amount" validate:"required,gte=1,lte=1000000,numeric"`
}
