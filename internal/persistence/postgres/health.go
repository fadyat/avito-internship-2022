package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type HealthRepo struct {
	c *pgx.Conn
}

func NewHealthRepo(c *pgx.Conn) *HealthRepo {
	return &HealthRepo{c: c}
}

func (hr *HealthRepo) Ping() error {
	return hr.c.Ping(context.Background())
}
