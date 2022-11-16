package services

import (
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/persistence/mock"
	"github.com/fadyat/avito-internship-2022/internal/responses"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

var (
	mockedWalletRepo  = &mock.MockedUserWalletRepo{}
	validate          = validator.New()
	userWalletService = NewUserWalletService(mockedWalletRepo, validate)
)

type testCreateUserWallet struct {
	name   string
	wallet dto.UserWallet
	expErr error
	expRes uint64
}

func TestUserWalletService_CreateUserWallet(t *testing.T) {
	tests := []testCreateUserWallet{
		{
			name:   "user_id is gt 0",
			wallet: dto.UserWallet{UserID: 1},
			expErr: nil,
			expRes: 1,
		},
		{
			name:   "user_id is lt 1",
			wallet: dto.UserWallet{UserID: 0},
			expErr: &responses.ValidationErrResp{},
			expRes: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := userWalletService.CreateUserWallet(tt.wallet)
			require.Equal(t, reflect.TypeOf(err), reflect.TypeOf(tt.expErr))
			require.Equal(t, res, tt.expRes)
		})
	}
}

type testGetUserWalletByID struct {
	name   string
	id     string
	expErr error
	expRes *models.UserWallet
}

func TestUserWalletService_GetUserWalletByID(t *testing.T) {
	tests := []testGetUserWalletByID{
		{
			name:   "user_id is gt 0",
			id:     "1",
			expErr: nil,
			expRes: &models.UserWallet{UserID: 1, Balance: 0},
		},
		{
			name:   "user_id is lt 1",
			id:     "0",
			expErr: &responses.ValidationErrResp{},
			expRes: nil,
		},
		{
			name:   "user_id is not a number",
			id:     "a",
			expErr: &responses.ValidationErrResp{},
			expRes: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := userWalletService.GetUserWalletByID(tt.id)
			require.Equal(t, reflect.TypeOf(err), reflect.TypeOf(tt.expErr))
			require.Equal(t, res, tt.expRes)
		})
	}
}

type testGetAllUserWallets struct {
	name   string
	expErr error
	expRes []*models.UserWallet
}

func TestUserWalletService_GetAllUserWallets(t *testing.T) {
	tests := []testGetAllUserWallets{
		{
			name:   "get all wallets",
			expErr: nil,
			expRes: []*models.UserWallet{
				{UserID: 1, Balance: 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := userWalletService.GetAllUserWallets()
			require.Equal(t, reflect.TypeOf(err), reflect.TypeOf(tt.expErr))
			require.Equal(t, res, tt.expRes)
		})
	}
}
