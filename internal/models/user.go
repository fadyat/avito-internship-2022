package models

// UserWallet godoc
// @description: UserWallet is a wallet of a user.
type UserWallet struct {
	// @description: ID is a unique identifier of the wallet, in real cases it's a card number or something like that.
	ID uint `json:"id"`

	// @description: UserID is a unique identifier of the user, that owns this wallet.
	UserID uint `json:"user_id"`

	// @description: Balance is a current balance of the wallet.
	Balance uint `json:"balance"`
}
