package responses

import "github.com/fadyat/avito-internship-2022/internal/models"

// UserWallerCreated godoc
// @description: UserWallerCreated is a response for user wallet creation.
type UserWallerCreated struct {

	// @description: ID is given unique identifier of the wallet.
	// @example:     1
	ID uint64 `json:"id"`
}

// UserWallets godoc
// @description: UserWallets is a response for getting all user wallets.
type UserWallets struct {

	// @description: UserWallets is a list of user wallets.
	// @example:     [{"user_id":1,"balance":100}]
	Wallets []*models.UserWallet `json:"wallets"`
}
