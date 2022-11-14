package models

// UserWallet godoc
// @description: UserWallet is a wallet of a user.
type UserWallet struct {

	// @description: UserID is a unique identifier of the user, that owns this wallet.
	// @example:     1
	UserID uint64 `json:"user_id"`

	// @description: Balance is a current balance of the wallet.
	// @example:     100
	Balance uint64 `json:"balance"`
}
