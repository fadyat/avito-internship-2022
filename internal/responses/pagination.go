package responses

// Pagination godoc
// @description: Pagination is a pagination response, which is used in paginated responses.
type Pagination struct {

	// @description: PrevPage is a number of the previous page.
	// @example:     1
	PrevPage uint64 `json:"prev_page,omitempty"`

	// @description: Page is a number of the current page.
	// @example:     2
	Page uint64 `json:"page"`

	// @description: NextPage is a number of the next page.
	// @example:     3
	NextPage uint64 `json:"next_page,omitempty"`

	// @description: Found is a number of found items.
	// @example:     10
	Found uint64 `json:"found"`

	// @description: PerPage is a number of items per page.
	// @example:     10
	PerPage uint64 `json:"per_page"`

	// @description: Total is a total number of items.
	// @example:     100
	Total uint64 `json:"total,omitempty"`
}

func NewPagination(perPage, page, found, total uint64) *Pagination {
	basic := &Pagination{
		Found:   found,
		Page:    page,
		PerPage: perPage,
	}

	if total > 0 {
		basic.Total = total
	}

	if page > 1 {
		basic.PrevPage = page - 1
	}

	if page > total/perPage {
		basic.PrevPage = total / perPage
	}

	if found == perPage && total > page*perPage {
		basic.NextPage = page + 1
	}

	return basic
}
