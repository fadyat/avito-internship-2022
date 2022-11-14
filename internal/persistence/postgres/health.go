package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type HealthRepository struct {
	c *pgx.Conn
}

func NewHealthRepository(c *pgx.Conn) *HealthRepository {
	return &HealthRepository{c: c}
}

func (hr *HealthRepository) Ping() error {
	return hr.c.Ping(context.Background())
}
