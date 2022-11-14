package dto

type UserWallet struct {
	UserID uint64 `json:"user_id" validate:"required,min=1"`
}
