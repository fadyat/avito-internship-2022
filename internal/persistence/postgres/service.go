package postgres

import (
	"context"
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/jackc/pgx/v5"
)

type OuterServiceRepository struct {
	c *pgx.Conn
}

func NewOuterServiceRepository(c *pgx.Conn) *OuterServiceRepository {
	return &OuterServiceRepository{c: c}
}

func (s *OuterServiceRepository) CreateService(os dto.OuterService) (uint64, error) {
	id, err := s.createService(os)
	return id, recastError(err)
}

func (s *OuterServiceRepository) createService(os dto.OuterService) (uint64, error) {
	tx, err := s.c.Begin(context.Background())
	if err != nil {
		return 0, err
	}
	defer func() { _ = tx.Rollback(context.Background()) }()

	var id uint64
	q := "INSERT INTO services (name, url) VALUES ($1, $2) RETURNING id"
	err = tx.QueryRow(context.Background(), q, os.Name, os.URL).Scan(&id)
	if err != nil {
		return 0, err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return 0, err
	}

	return id, nil
}

func (s *OuterServiceRepository) GetServiceByID(id uint64) (*models.OuterService, error) {
	svcs, err := s.getServiceByID(id)
	return svcs, recastError(err)
}

func (s *OuterServiceRepository) getServiceByID(id uint64) (*models.OuterService, error) {
	tx, err := s.c.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback(context.Background()) }()

	var os models.OuterService
	q := "SELECT * FROM services WHERE id = $1 LIMIT 1"
	err = tx.QueryRow(context.Background(), q, id).Scan(&os.ID, &os.Name, &os.URL)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return nil, err
	}

	return &os, nil
}

func (s *OuterServiceRepository) GetAllServices() ([]*models.OuterService, error) {
	svcs, err := s.getAllServices()
	return svcs, recastError(err)
}

func (s *OuterServiceRepository) getAllServices() ([]*models.OuterService, error) {
	tx, err := s.c.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback(context.Background()) }()

	q := "SELECT * FROM services"
	rows, err := tx.Query(context.Background(), q)
	if err != nil {
		return nil, err
	}

	var services []*models.OuterService
	for rows.Next() {
		var os models.OuterService
		_ = rows.Scan(&os.ID, &os.Name, &os.URL)
		services = append(services, &os)
	}

	return services, nil
}
