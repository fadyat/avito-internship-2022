package models

type Pagination struct {
	Page    uint64   `json:"page" query:"page"`
	PerPage uint64   `json:"per_page" query:"per_page"`
	OrderBy []string `json:"order_by" query:"order_by"`
}
