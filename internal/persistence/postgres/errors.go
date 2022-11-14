package postgres

import (
	"database/sql"
	"github.com/jackc/pgx/v5"
)

func recastError(err error) error {
	switch err {
	case pgx.ErrNoRows:
		return sql.ErrNoRows
	default:
		return err
	}
}
