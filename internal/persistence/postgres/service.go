package postgres

import (
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/jmoiron/sqlx"
)

type OuterServiceRepo struct {
	c *sqlx.DB
}

func NewOuterServiceRepo(c *sqlx.DB) *OuterServiceRepo {
	return &OuterServiceRepo{c: c}
}

func (s *OuterServiceRepo) CreateService(os dto.OuterService) (uint64, error) {
	id, err := s.createService(os)
	return id, recastError(err)
}

func (s *OuterServiceRepo) createService(os dto.OuterService) (uint64, error) {
	tx, err := s.c.Begin()
	if err != nil {
		return 0, err
	}
	defer func() { _ = tx.Rollback() }()

	q := "INSERT INTO services (name, url) VALUES ($1, $2) RETURNING id"
	var id uint64
	err = tx.QueryRow(q, os.Name, os.URL).Scan(&id)
	if err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return id, nil
}

func (s *OuterServiceRepo) GetServiceByID(id uint64) (*models.OuterService, error) {
	svcs, err := s.getServiceByID(id)
	return svcs, recastError(err)
}

func (s *OuterServiceRepo) getServiceByID(id uint64) (*models.OuterService, error) {
	tx, err := s.c.Begin()
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback() }()

	var os models.OuterService
	q := "SELECT * FROM services WHERE id = $1 LIMIT 1"
	err = tx.QueryRow(q, id).Scan(&os.ID, &os.Name, &os.URL)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &os, nil
}

func (s *OuterServiceRepo) GetAllServices() ([]*models.OuterService, error) {
	svcs, err := s.getAllServices()
	return svcs, recastError(err)
}

func (s *OuterServiceRepo) getAllServices() ([]*models.OuterService, error) {
	tx, err := s.c.Begin()
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback() }()

	q := "SELECT * FROM services"
	rows, err := tx.Query(q)
	if err != nil {
		return nil, err
	}

	var services []*models.OuterService
	for rows.Next() {
		var os models.OuterService
		_ = rows.Scan(&os.ID, &os.Name, &os.URL)
		services = append(services, &os)
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return services, nil
}
