package dto

// Replenishment godoc
// @description: Replenishment is a request to replenish a user's wallet.
type Replenishment struct {
	UserID uint64 `json:"user_id" validate:"required,min=1"`
	Amount uint64 `json:"amount" validate:"required,gt=0"`
}

// Withdrawal godoc
// @description: Withdrawal is a request to withdraw money from the user's wallet.
type Withdrawal struct {
	UserID uint64 `json:"user_id" validate:"required,min=1"`
	Amount uint64 `json:"amount" validate:"required,gt=0"`
}

// Reservation godoc
// @description: Reservation is a request to reserve money for a future transaction.
type Reservation struct {
	UserID    uint64 `json:"user_id" validate:"required,min=1"`
	ServiceID uint64 `json:"service_id" validate:"required,min=1"`
	OrderID   uint64 `json:"order_id" validate:"required,min=1"`
	Amount    uint64 `json:"amount" validate:"required,gt=0"`
}

// Release godoc
// @description: Release is a request to release a reservation.
type Release struct {
	UserID    uint64 `json:"user_id" validate:"required,min=1"`
	ServiceID uint64 `json:"service_id" validate:"required,min=1"`
	OrderID   uint64 `json:"order_id" validate:"required,min=1"`
	Amount    uint64 `json:"amount" validate:"required,gt=0"`
}
