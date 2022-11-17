package postgres

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/persistence"
	"testing"
)

type testCreateService struct {
	name   string
	s      *dto.OuterService
	expRes uint64
	expErr error
}

func TestOuterServiceRepo_CreateService(t *testing.T) {
	tests := []testCreateService{
		{
			name:   "no error",
			s:      &dto.OuterService{Name: "test"},
			expRes: 1,
			expErr: nil,
		},
		{
			name:   "no errors url",
			s:      &dto.OuterService{Name: "test", URL: "https://test.com"},
			expRes: 1,
			expErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := createConn(t)
			defer func() { _ = db.Close() }()

			mock.ExpectBegin()
			mock.ExpectQuery("INSERT INTO services").
				WithArgs(tt.s.Name, tt.s.URL).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).
					AddRow(1),
				)
			mock.ExpectCommit()

			s := NewOuterServiceRepo(db)
			res, err := s.CreateService(tt.s)

			compareErrors(t, err, tt.expErr)
			compareResults(t, res, tt.expRes, mock)
		})
	}
}

type testGetServiceByID struct {
	name   string
	id     uint64
	expRes *models.OuterService
	expErr error
}

func TestOuterServiceRepo_GetServiceByID(t *testing.T) {
	tests := []testGetServiceByID{
		{
			name:   "no error",
			id:     1,
			expRes: &models.OuterService{ID: 1, Name: "test"},
			expErr: nil,
		},
		{
			name:   "no error URL",
			id:     1,
			expRes: &models.OuterService{ID: 1, Name: "test", URL: "https://test.com"},
			expErr: nil,
		},
		{
			name:   "not found",
			id:     1,
			expRes: nil,
			expErr: persistence.ErrNotFound,
		},
	}
	q := "SELECT (.+) FROM services"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := createConn(t)
			defer func() { _ = db.Close() }()

			mock.ExpectBegin()
			if tt.expErr == nil {
				rows := sqlmock.NewRows([]string{"id", "name", "url"}).
					AddRow(tt.expRes.ID, tt.expRes.Name, tt.expRes.URL)

				createSuccessMock(t, mock, q, rows, tt.id)
				mock.ExpectCommit()
			} else {
				createErrorMock(t, mock, q, tt.expErr, tt.id)
				mock.ExpectRollback()
			}

			s := NewOuterServiceRepo(db)
			res, err := s.GetServiceByID(tt.id)

			compareErrors(t, err, tt.expErr)
			compareResults(t, res, tt.expRes, mock)
		})
	}
}

type testGetAllServices struct {
	name   string
	expRes []*models.OuterService
	expErr error
}

func TestOuterServiceRepo_GetAllServices(t *testing.T) {
	tests := []testGetAllServices{
		{
			name: "no error",
			expRes: []*models.OuterService{
				{ID: 1, Name: "test"},
				{ID: 2, Name: "test2", URL: "https://test2.com"},
			},
			expErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := createConn(t)
			defer func() { _ = db.Close() }()

			mock.ExpectBegin()
			rows := sqlmock.NewRows([]string{"id", "name", "url"})
			for _, s := range tt.expRes {
				rows.AddRow(s.ID, s.Name, s.URL)
			}
			mock.ExpectQuery("SELECT (.+) FROM services").WillReturnRows(rows)
			mock.ExpectCommit()

			s := NewOuterServiceRepo(db)
			res, err := s.GetAllServices()

			compareErrors(t, err, tt.expErr)
			compareResults(t, res, tt.expRes, mock)
		})
	}
}
