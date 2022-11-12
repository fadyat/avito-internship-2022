package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type HealthRepository struct {
	client *pgx.Conn
}

func NewHealthRepository(client *pgx.Conn) *HealthRepository {
	return &HealthRepository{
		client: client,
	}
}

func (hr *HealthRepository) Ping() error {
	return hr.client.Ping(context.Background())
}
