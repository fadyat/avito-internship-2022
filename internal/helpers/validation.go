package helpers

import (
	"fmt"
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/responses"
	"github.com/go-playground/validator/v10"
	"strconv"
)

func ValidateUint64(n, valTag string, v *validator.Validate) (uint64, error) {
	number, err := strconv.ParseUint(n, 10, 64)
	if err != nil {
		return 0, &responses.ValidationErrResp{Message: fmt.Sprintf("failed to parse %s as uint64", n)}
	}

	if err := v.Var(number, valTag); err != nil {
		return 0, &responses.ValidationErrResp{Message: err.Error()}
	}

	return number, nil
}

func MakePaginationParamsCorrect(p *models.Pagination) {
	if p.Page == 0 {
		p.Page = 1
	}

	if p.PerPage == 0 {
		p.PerPage = 10
	}

	if p.PerPage > 100 {
		p.PerPage = 100
	}

	if len(p.OrderBy) == 0 {
		p.OrderBy = []string{"created_at", "amount"}
	}
}
