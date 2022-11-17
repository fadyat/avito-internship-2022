package postgres

import (
	"github.com/jmoiron/sqlx"
)

type HealthRepo struct {
	c *sqlx.DB
}

func NewHealthRepo(c *sqlx.DB) *HealthRepo {
	return &HealthRepo{c: c}
}

func (hr *HealthRepo) Ping() error {
	return hr.c.Ping()
}
