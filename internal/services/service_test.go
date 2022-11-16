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
	mockedOuterServiceRepo = &mock.OuterServiceRepo{}
	outerServiceService    = NewOuterServiceService(mockedOuterServiceRepo, validate)
)

type testCreateOuterService struct {
	name    string
	service dto.OuterService
	expErr  error
	expRes  uint64
}

func TestOuterServiceService_CreateService(t *testing.T) {
	tests := []testCreateOuterService{
		{
			name:    "name is empty",
			service: dto.OuterService{Name: "", URL: "https://aboba.com"},
			expErr:  &responses.ValidationErrResp{},
			expRes:  0,
		},
		{
			name:    "url is empty",
			service: dto.OuterService{Name: "aboba", URL: ""},
			expErr:  nil,
			expRes:  1,
		},
		{
			name:    "name and url are empty",
			service: dto.OuterService{Name: "", URL: ""},
			expErr:  &responses.ValidationErrResp{},
			expRes:  0,
		},
		{
			name:    "name and url are not empty",
			service: dto.OuterService{Name: "aboba", URL: "https://aboba.com"},
			expErr:  nil,
			expRes:  1,
		},
		{
			name:    "url is not valid",
			service: dto.OuterService{Name: "aboba", URL: "aboba.com"},
			expErr:  &responses.ValidationErrResp{},
			expRes:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := outerServiceService.CreateService(tt.service)
			require.Equal(t, reflect.TypeOf(err), reflect.TypeOf(tt.expErr))
			require.Equal(t, res, tt.expRes)
		})
	}
}

type testGetServiceByID struct {
	name   string
	id     string
	expErr error
	expRes *models.OuterService
}

func TestOuterServiceService_GetServiceByID(t *testing.T) {
	tests := []testGetServiceByID{
		{
			name:   "id is gt 0",
			id:     "1",
			expErr: nil,
			expRes: &models.OuterService{ID: 1, Name: "aboba", URL: "https://aboba.com"},
		},
		{
			name:   "id is lt 1",
			id:     "0",
			expErr: &responses.ValidationErrResp{},
			expRes: nil,
		},
		{
			name:   "id is not a number",
			id:     "aboba",
			expErr: &responses.ValidationErrResp{},
			expRes: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := outerServiceService.GetServiceByID(tt.id)
			require.Equal(t, reflect.TypeOf(err), reflect.TypeOf(tt.expErr))
			require.Equal(t, res, tt.expRes)
		})
	}
}

type testGetAllServices struct {
	name   string
	expErr error
	expRes []*models.OuterService
}

func TestOuterServiceService_GetAllServices(t *testing.T) {
	tests := []testGetAllServices{
		{
			name:   "no errors",
			expErr: nil,
			expRes: []*models.OuterService{{ID: 1, Name: "aboba", URL: "https://aboba.com"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := outerServiceService.GetAllServices()
			require.Equal(t, reflect.TypeOf(err), reflect.TypeOf(tt.expErr))
			require.Equal(t, res, tt.expRes)
		})
	}
}
