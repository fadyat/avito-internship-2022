package postgres

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/persistence"
	"testing"
)

type testReplenishment struct {
	name   string
	t      dto.Transaction
	expRes uint64
	expErr error
}

func TestTransactionRepo_CreateReplenishment(t *testing.T) {
	tests := []testReplenishment{
		{
			name:   "no error",
			t:      dto.Transaction{UserID: 1, Amount: 100},
			expRes: 2,
			expErr: nil,
		},
		{
			name:   "amount is zero",
			t:      dto.Transaction{UserID: 1, Amount: 0},
			expErr: persistence.ErrNegativeAmount,
			expRes: 0,
		},
	}

	tq := "INSERT INTO transactions"
	uq := "UPDATE user_wallets"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := createConn(t)
			defer func() { _ = db.Close() }()

			mock.ExpectBegin()
			if tt.expErr == nil {
				mock.ExpectExec(uq).
					WithArgs(tt.t.Amount, tt.t.UserID).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectQuery(tq).
					WithArgs(tt.t.UserID, tt.t.Amount, models.Replenishment).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).
						AddRow(tt.expRes),
					)
				mock.ExpectCommit()
			} else {
				mock.ExpectRollback()
			}

			u := NewTransactionRepo(db)
			res, err := u.CreateReplenishment(tt.t)

			compareErrors(t, err, tt.expErr)
			compareResults(t, res, tt.expRes, mock)
		})
	}
}

type testWithdrawal struct {
	name     string
	fillMock func(mock sqlmock.Sqlmock)
	t        dto.Transaction
	expRes   uint64
	expErr   error
}

func TestTransactionRepo_CreateWithdrawal(t *testing.T) {
	bq := "SELECT balance FROM user_wallets"
	tq := "INSERT INTO transactions"
	uq := "UPDATE user_wallets"

	tests := []testWithdrawal{
		{
			name:   "no error",
			t:      dto.Transaction{UserID: 1, Amount: 100},
			expRes: 2,
			expErr: nil,
			fillMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(bq).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"balance"}).
						AddRow(1000),
					)

				mock.ExpectExec(uq).
					WithArgs(100, 1).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectQuery(tq).
					WithArgs(1, 100, models.Withdrawal).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).
						AddRow(2),
					)

				mock.ExpectCommit()
			},
		},
		{
			name:   "amount is zero",
			expRes: 0,
			expErr: persistence.ErrNegativeAmount,
			t:      dto.Transaction{UserID: 1, Amount: 0},
			fillMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectRollback()
			},
		},
		{
			name:   "not enough money",
			expRes: 0,
			expErr: persistence.ErrInsufficientFunds,
			t:      dto.Transaction{UserID: 1, Amount: 100},
			fillMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(bq).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"balance"}).
						AddRow(0),
					)

				mock.ExpectRollback()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := createConn(t)
			defer func() { _ = db.Close() }()

			tt.fillMock(mock)

			u := NewTransactionRepo(db)
			res, err := u.CreateWithdrawal(tt.t)

			compareErrors(t, err, tt.expErr)
			compareResults(t, res, tt.expRes, mock)
		})
	}
}

type testGetUserTransactions struct {
	name     string
	fillMock func(mock sqlmock.Sqlmock)
	userID   uint64
	expRes   []*models.Transaction
	expErr   error
}

func TestTransactionRepo_GetUserTransactions(t *testing.T) {
	t.Skip("not implemented")
}

type testGetUserTransactionsCount struct {
	name   string
	expRes uint64
	expErr error
	userID uint64
}

func TestTransactionRepo_GetUserTransactionsCount(t *testing.T) {
	q := "SELECT COUNT(.+) FROM transactions"

	tests := []testGetUserTransactionsCount{
		{
			name:   "no error",
			userID: 1,
			expRes: 2,
			expErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := createConn(t)
			defer func() { _ = db.Close() }()

			mock.ExpectBegin()
			mock.ExpectQuery(q).
				WithArgs(tt.userID).
				WillReturnRows(sqlmock.NewRows([]string{"count"}).
					AddRow(tt.expRes),
				)
			mock.ExpectCommit()

			u := NewTransactionRepo(db)
			res, err := u.GetUserTransactionsCount(tt.userID)

			compareErrors(t, err, tt.expErr)
			compareResults(t, res, tt.expRes, mock)
		})
	}
}
