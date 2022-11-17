package postgres

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/persistence"
	"testing"
)

type testCreateUserWallet struct {
	name   string
	w      dto.UserWallet
	expRes uint64
	expErr error
}

func TestUserWalletRepo_CreateUserWallet(t *testing.T) {
	tests := []testCreateUserWallet{
		{
			name:   "no error",
			w:      dto.UserWallet{UserID: 1},
			expRes: 1,
			expErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := createConn(t)
			defer func() { _ = db.Close() }()

			mock.ExpectBegin()
			mock.ExpectQuery("INSERT INTO user_wallets").
				WithArgs(tt.w.UserID, 0).
				WillReturnRows(sqlmock.NewRows([]string{"user_id"}).
					AddRow(tt.expRes),
				)
			mock.ExpectCommit()

			u := NewUserWalletRepo(db)
			res, err := u.CreateUserWallet(tt.w)

			compareErrors(t, err, tt.expErr)
			compareResults(t, res, tt.expRes, mock)
		})
	}
}

type testGetUserWalletByID struct {
	name   string
	id     uint64
	expRes *models.UserWallet
	expErr error
}

func TestUserWalletRepo_GetUserWalletByID(t *testing.T) {
	tests := []testGetUserWalletByID{
		{
			name:   "no error",
			id:     1,
			expRes: &models.UserWallet{UserID: 1, Balance: 0},
			expErr: nil,
		},
		{
			name:   "not found",
			id:     2,
			expRes: nil,
			expErr: persistence.ErrNotFound,
		},
	}
	q := "SELECT (.+) FROM user_wallets"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := createConn(t)
			defer func() { _ = db.Close() }()

			mock.ExpectBegin()
			if tt.expErr == nil {
				rows := sqlmock.NewRows([]string{"user_id", "balance"}).
					AddRow(tt.expRes.UserID, tt.expRes.Balance)
				createSuccessMock(t, mock, q, rows, tt.id)
				mock.ExpectCommit()
			} else {
				createErrorMock(t, mock, q, tt.expErr, tt.id)
				mock.ExpectRollback()
			}

			u := NewUserWalletRepo(db)
			res, err := u.GetUserWalletByID(tt.id)

			compareErrors(t, err, tt.expErr)
			compareResults(t, res, tt.expRes, mock)
		})
	}

}

type testGetUserAllUserWallets struct {
	name   string
	expRes []*models.UserWallet
	expErr error
}

func TestUserWalletRepo_GetAllWallets(t *testing.T) {
	tests := []testGetUserAllUserWallets{
		{
			name: "no error",
			expRes: []*models.UserWallet{
				{UserID: 1, Balance: 0},
				{UserID: 2, Balance: 100},
			},
			expErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := createConn(t)
			defer func() { _ = db.Close() }()

			mock.ExpectBegin()
			rows := sqlmock.NewRows([]string{"user_id", "balance"})
			for _, w := range tt.expRes {
				rows.AddRow(w.UserID, w.Balance)
			}
			mock.ExpectQuery("SELECT (.+) FROM user_wallets").WillReturnRows(rows)
			mock.ExpectCommit()

			u := NewUserWalletRepo(db)
			res, err := u.GetAllWallets()

			compareErrors(t, err, tt.expErr)
			compareResults(t, res, tt.expRes, mock)
		})
	}
}
