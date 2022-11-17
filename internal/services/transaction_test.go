package services

import (
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/persistence/mock"
	"github.com/fadyat/avito-internship-2022/internal/responses"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

var (
	mockedTransactionRepo = &mock.TransactionRepo{}
	transactionService    = NewTransactionService(mockedTransactionRepo, validate)
)

type testCreateTransaction struct {
	name   string
	tr     *dto.Transaction
	expErr error
	expRes uint64
}

var createTransactionTests = []testCreateTransaction{
	{
		name:   "amount is lt 1",
		tr:     &dto.Transaction{UserID: 1, Amount: 0},
		expErr: &responses.ValidationErrResp{},
		expRes: 0,
	},
	{
		name:   "user id is lt 1",
		tr:     &dto.Transaction{UserID: 0, Amount: 1},
		expErr: &responses.ValidationErrResp{},
		expRes: 0,
	},
	{
		name:   "amount is gt 1000000",
		tr:     &dto.Transaction{UserID: 1, Amount: 1000001},
		expErr: &responses.ValidationErrResp{},
		expRes: 0,
	},
	{
		name:   "missing user id",
		tr:     &dto.Transaction{Amount: 1},
		expErr: &responses.ValidationErrResp{},
		expRes: 0,
	},
	{
		name:   "missing amount",
		tr:     &dto.Transaction{UserID: 1},
		expErr: &responses.ValidationErrResp{},
		expRes: 0,
	},
	{
		name:   "no errors",
		tr:     &dto.Transaction{UserID: 1, Amount: 1},
		expErr: nil,
		expRes: 1,
	},
}

func TestTransactionService_CreateReplenishment(t *testing.T) {
	for _, tt := range createTransactionTests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := transactionService.CreateReplenishment(tt.tr)
			require.Equal(t, reflect.TypeOf(err), reflect.TypeOf(tt.expErr))
			require.Equal(t, res, tt.expRes)
		})
	}
}

func TestTransactionService_CreateWithdrawal(t *testing.T) {
	for _, tt := range createTransactionTests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := transactionService.CreateWithdrawal(tt.tr)
			require.Equal(t, reflect.TypeOf(err), reflect.TypeOf(tt.expErr))
			require.Equal(t, res, tt.expRes)
		})
	}
}

type testGetUserTransactions struct {
	name   string
	userID string
	pag    *models.Pagination
	expErr error
	expRes []*models.Transaction
}

func TestTransactionService_GetUserTransactions(t *testing.T) {
	tests := []testGetUserTransactions{
		{
			name:   "user id is lt 1",
			userID: "0",
			pag: &models.Pagination{
				Page:    1,
				PerPage: 10,
				OrderBy: []string{"created_at"},
			},
			expErr: &responses.ValidationErrResp{},
			expRes: nil,
		},
		{
			name:   "user id is gte 1",
			userID: "1",
			pag: &models.Pagination{
				Page:    1,
				PerPage: 10,
				OrderBy: []string{"created_at"},
			},
			expErr: nil,
			expRes: make([]*models.Transaction, 0),
		},
		{
			name:   "user id is not a number",
			userID: "a",
			pag: &models.Pagination{
				Page:    1,
				PerPage: 10,
				OrderBy: []string{"created_at"},
			},
			expErr: &responses.ValidationErrResp{},
			expRes: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := transactionService.GetUserTransactions(tt.userID, tt.pag)
			require.Equal(t, reflect.TypeOf(err), reflect.TypeOf(tt.expErr))
			require.Equal(t, res, tt.expRes)
		})
	}
}

type testGetUserTransactionsCount struct {
	name   string
	userID string
	expErr error
	expRes uint64
}

func TestTransactionService_GetUserTransactionsCount(t *testing.T) {
	var tests = []testGetUserTransactionsCount{
		{
			name:   "user id is lt 1",
			userID: "0",
			expErr: &responses.ValidationErrResp{},
			expRes: 0,
		},
		{
			name:   "user id is gte 1",
			userID: "1",
			expErr: nil,
			expRes: 1,
		},
		{
			name:   "user id is not a number",
			userID: "a",
			expErr: &responses.ValidationErrResp{},
			expRes: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := transactionService.GetUserTransactionsCount(tt.userID)
			require.Equal(t, reflect.TypeOf(err), reflect.TypeOf(tt.expErr))
			require.Equal(t, res, tt.expRes)
		})
	}
}
