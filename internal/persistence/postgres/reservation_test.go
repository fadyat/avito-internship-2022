package postgres

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/persistence"
	"testing"
)

type testReservation struct {
	name     string
	r        *dto.Reservation
	expErr   error
	expRes   uint64
	fillMock func(mock sqlmock.Sqlmock, expRes uint64, r *dto.Reservation)
}

func TestTransactionRepo_CreateReservation(t *testing.T) {
	rq := "INSERT INTO reservations"

	tests := []testReservation{
		{
			name: "no error",
			r: &dto.Reservation{
				UserID:    1,
				Amount:    100,
				ServiceID: 1,
				OrderID:   1,
			},
			expRes: 2,
			expErr: nil,
			fillMock: func(mock sqlmock.Sqlmock, expRes uint64, r *dto.Reservation) {
				mock.ExpectBegin()
				mock.ExpectQuery(rq).
					WithArgs(r.UserID, r.ServiceID, r.OrderID, r.Amount, models.Pending).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).
						AddRow(expRes),
					)
				mock.ExpectCommit()
			},
		},
		{
			name: "amount is zero",
			r: &dto.Reservation{
				UserID:    1,
				Amount:    0,
				ServiceID: 1,
				OrderID:   1,
			},
			expErr: persistence.ErrNegativeAmount,
			expRes: 0,
			fillMock: func(mock sqlmock.Sqlmock, expRes uint64, r *dto.Reservation) {
				mock.ExpectBegin()
				mock.ExpectRollback()
			},
		},
		{
			name: "service not found",
			r: &dto.Reservation{
				UserID:    1,
				Amount:    100,
				ServiceID: 10,
				OrderID:   1,
			},
			expErr: persistence.ErrForeignKeyViolation,
			expRes: 0,
			fillMock: func(mock sqlmock.Sqlmock, expRes uint64, r *dto.Reservation) {
				mock.ExpectBegin()
				mock.ExpectQuery(rq).
					WithArgs(r.UserID, r.ServiceID, r.OrderID, r.Amount, models.Pending).
					WillReturnError(persistence.ErrForeignKeyViolation)
				mock.ExpectRollback()
			},
		},
		{
			name: "user not found",
			r: &dto.Reservation{
				UserID:    10,
				Amount:    100,
				ServiceID: 1,
				OrderID:   1,
			},
			expErr: persistence.ErrForeignKeyViolation,
			expRes: 0,
			fillMock: func(mock sqlmock.Sqlmock, expRes uint64, r *dto.Reservation) {
				mock.ExpectBegin()
				mock.ExpectQuery(rq).
					WithArgs(r.UserID, r.ServiceID, r.OrderID, r.Amount, models.Pending).
					WillReturnError(persistence.ErrForeignKeyViolation)
				mock.ExpectRollback()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := createConn(t)
			defer func() { _ = db.Close() }()

			tt.fillMock(mock, tt.expRes, tt.r)

			u := NewTransactionRepo(db)
			res, err := u.CreateReservation(tt.r)

			compareErrors(t, err, tt.expErr)
			compareResults(t, res, tt.expRes, mock)
		})
	}
}

func TestTransactionRepo_CreateRelease(t *testing.T) {
	t.Skip("not implemented")
}

func TestTransactionRepo_CancelReservation(t *testing.T) {
	t.Skip("not implemented")
}

type testReservationReport struct {
	name     string
	r        *dto.ReportTime
	expErr   error
	expRes   []*models.ReservationReport
	fillMock func(mock sqlmock.Sqlmock, expRes []*models.ReservationReport, r *dto.ReportTime)
}

func TestTransactionRepo_GetReservationsReport(t *testing.T) {
	t.Skip("not implemented")

	// todo: implement
	q := ``

	tests := []testReservationReport{
		{
			name: "no error",
			r: &dto.ReportTime{
				Year:  2021,
				Month: 1,
			},
			expRes: []*models.ReservationReport{
				{
					ServiceID: 1,
					OrderID:   1,
					Amount:    100,
					Count:     20,
				},
				{
					ServiceID: 2,
					OrderID:   2,
					Amount:    200,
					Count:     10,
				},
			},
			expErr: nil,
			fillMock: func(mock sqlmock.Sqlmock, expRes []*models.ReservationReport, r *dto.ReportTime) {
				rows := sqlmock.NewRows([]string{"service_id", "order_id", "amount", "count"})
				for _, v := range expRes {
					rows.AddRow(v.ServiceID, v.OrderID, v.Amount, v.Count)
				}
				mock.ExpectBegin()
				mock.ExpectQuery(q).
					WithArgs(r.Year, r.Month, models.Released).
					WillReturnRows(rows)
				mock.ExpectCommit()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := createConn(t)
			defer func() { _ = db.Close() }()

			tt.fillMock(mock, tt.expRes, tt.r)

			u := NewTransactionRepo(db)
			res, err := u.GetReservationsReport(tt.r)

			compareErrors(t, err, tt.expErr)
			compareResults(t, res, tt.expRes, mock)
		})
	}
}
