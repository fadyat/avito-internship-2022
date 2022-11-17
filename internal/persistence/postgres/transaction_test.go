package postgres

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/persistence"
	"testing"
	"time"
)

type testReplenishment struct {
	name   string
	t      *dto.Transaction
	expRes uint64
	expErr error
}

func TestTransactionRepo_CreateReplenishment(t *testing.T) {
	tests := []testReplenishment{
		{
			name:   "no error",
			t:      &dto.Transaction{UserID: 1, Amount: 100},
			expRes: 2,
			expErr: nil,
		},
		{
			name:   "amount is zero",
			t:      &dto.Transaction{UserID: 1, Amount: 0},
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
	t        *dto.Transaction
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
			t:      &dto.Transaction{UserID: 1, Amount: 100},
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
			t:      &dto.Transaction{UserID: 1, Amount: 0},
			fillMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectRollback()
			},
		},
		{
			name:   "not enough money",
			expRes: 0,
			expErr: persistence.ErrInsufficientFunds,
			t:      &dto.Transaction{UserID: 1, Amount: 100},
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
	fillMock func(mock sqlmock.Sqlmock, userID uint64, pag *models.Pagination, expRes []*models.Transaction)
	userID   uint64
	pag      *models.Pagination
	expRes   []*models.Transaction
	expErr   error
}

func TestTransactionRepo_GetUserTransactions(t *testing.T) {
	var fixedTime = time.Now()

	tests := []testGetUserTransactions{
		{
			name:   "no error",
			userID: 1,
			pag:    &models.Pagination{Page: 1, PerPage: 10, OrderBy: []string{"created_at"}},
			expRes: []*models.Transaction{
				{ID: 1, UserID: 1, Amount: 100, Type: models.Replenishment, CreatedAt: fixedTime},
				{ID: 2, UserID: 1, Amount: 100, Type: models.Withdrawal, CreatedAt: fixedTime},
			},
			expErr: nil,
			fillMock: func(mock sqlmock.Sqlmock, userID uint64, pag *models.Pagination, expRes []*models.Transaction) {
				//q := fmt.Sprintf(
				//	"SELECT (.+) FROM transactions ORDER BY %s LIMIT ? OFFSET ?",
				//	strings.Join(pag.OrderBy, ", "),
				//)
				q := ``
				//q := regexp.QuoteMeta(`
				//	SELECT (.+) FROM transactions
				//	WHERE ((user_id = ?))
				//	ORDER BY created_at
				//`)
				row := sqlmock.NewRows([]string{"id", "user_id", "amount", "type", "created_at"})
				for _, tr := range expRes {
					row.AddRow(tr.ID, tr.UserID, tr.Amount, tr.Type, tr.CreatedAt)
				}

				mock.ExpectBegin()
				mock.ExpectQuery(q).
					WithArgs(userID, pag.PerPage, (pag.Page-1)*pag.PerPage).
					WillReturnRows(row)

				mock.ExpectCommit()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := createConn(t)
			defer func() { _ = db.Close() }()

			tt.fillMock(mock, tt.userID, tt.pag, tt.expRes)

			u := NewTransactionRepo(db)
			res, err := u.GetUserTransactions(tt.userID, tt.pag)

			compareErrors(t, err, tt.expErr)
			compareResults(t, res, tt.expRes, mock)
		})
	}
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
