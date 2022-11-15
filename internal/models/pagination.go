package models

// Pagination godoc
// @description: Pagination is a pagination of a list of items.
type Pagination struct {

	// @description: Page is a current page of the list.
	// @example:     1
	Page uint64 `json:"page" query:"page"`

	// @description: PerPage is a number of items per page.
	// @example:     10
	PerPage uint64 `json:"per_page" query:"per_page"`

	// @description: OrderBy is a field containing a list of elements to sort by.
	// @example:     created_at,amount
	OrderBy []string `json:"order_by" query:"order_by"`
}
