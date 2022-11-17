package postgres

import (
	"database/sql"
	"errors"
	"github.com/fadyat/avito-internship-2022/internal/persistence"
	"github.com/lib/pq"
)

func recastError(err error) error {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		switch pqErr.Code {
		case "23505":
			return persistence.ErrUniqueViolation
		case "23503":
			return persistence.ErrForeignKeyViolation
		case "42703":
			return persistence.ErrInvalidColumn
		}
	} else if errors.Is(err, sql.ErrNoRows) {
		return persistence.ErrNotFound
	}

	return err
}
