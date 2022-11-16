package dto

// ReportTime godoc
// @description: ReportTime is a model, that specifies the time of the report.
type ReportTime struct {

	// @description: Year is a year of the report.
	// @example:     2021
	Year int `json:"year" validate:"required,gte=0,number"`

	// @description: Month is a month of the report.
	// @example:     1
	Month int `json:"month" validate:"required,gt=0,lte=12,number"`
}
