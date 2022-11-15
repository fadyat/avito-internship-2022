package responses

import "github.com/fadyat/avito-internship-2022/internal/models"

// TransactionPaginated godoc
// @description TransactionPaginated is a response for paginated transactions
type TransactionPaginated struct {

	// @description Transactions is a list of transactions, which are paginated by page and perPage
	// @example     [{"id":1,"user_id":1,"service_id":1,"amount":100,"type":"replenishment","created_at":"2021-10-01T00:00:00Z"}]
	Transactions []*models.Transaction `json:"transactions"`

	// @description Pagination is a pagination object, which have info about pages
	// @example     {"prev_page":1,"page":2,"next_page":3,"found":10,"limit":10,"total":100}
	Pagination *Pagination `json:"pagination"`
}

// TransactionCreated godoc
// @description TransactionCreated is a response for transaction creation
type TransactionCreated struct {

	// @description ID is given unique identifier of the transaction
	// @example     1
	ID uint64 `json:"id"`
}
