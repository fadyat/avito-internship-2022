package postgres

import (
	"errors"
	"github.com/fadyat/avito-internship-2022/internal/persistence"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func recastError(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return persistence.ErrUniqueViolation
		case "23503":
			return persistence.ErrForeignKeyViolation
		}
	} else if errors.Is(err, pgx.ErrNoRows) {
		return persistence.ErrNotFound
	}

	return err
}
