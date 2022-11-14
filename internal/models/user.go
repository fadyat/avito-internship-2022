package models

// UserWallet godoc
// @description: UserWallet is a wallet of a user.
type UserWallet struct {
	// @description: UserID is a unique identifier of the user, that owns this wallet.
	UserID uint `json:"user_id"`

	// @description: Balance is a current balance of the wallet.
	Balance uint `json:"balance"`
}
