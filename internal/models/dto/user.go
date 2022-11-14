package dto

// UserWallet godoc
// @description: dto.UserWallet is dto for models.UserWallet.
type UserWallet struct {

	// @description: UserID is a unique identifier of the user, that owns this wallet.
	// @example:     1
	UserID uint64 `json:"user_id" validate:"required,min=1"`
}
