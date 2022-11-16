package services

import (
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/responses"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

type testCreateReservation struct {
	name   string
	res    dto.Reservation
	errRes error
	expRes uint64
}

var reservationTests = []testCreateReservation{
	{
		name: "user_id is lt 1",
		res: dto.Reservation{
			UserID:    0,
			Amount:    1,
			ServiceID: 1,
			OrderID:   1,
		},
		errRes: &responses.ValidationErrResp{},
		expRes: 0,
	},
	{
		name: "amount is lt 1",
		res: dto.Reservation{
			UserID:    1,
			Amount:    0,
			ServiceID: 1,
			OrderID:   1,
		},
		errRes: &responses.ValidationErrResp{},
		expRes: 0,
	},
	{
		name: "service_id is lt 1",
		res: dto.Reservation{
			UserID:    1,
			Amount:    1,
			ServiceID: 0,
			OrderID:   1,
		},
		errRes: &responses.ValidationErrResp{},
		expRes: 0,
	},
	{
		name: "order_id is lt 1",
		res: dto.Reservation{
			UserID:    1,
			Amount:    1,
			ServiceID: 1,
			OrderID:   0,
		},
		errRes: &responses.ValidationErrResp{},
		expRes: 0,
	},
	{
		name: "amount is gt 1000000",
		res: dto.Reservation{
			UserID:    1,
			Amount:    1000001,
			ServiceID: 1,
			OrderID:   1,
		},
		errRes: &responses.ValidationErrResp{},
		expRes: 0,
	},
	{
		name: "missing user_id",
		res: dto.Reservation{
			Amount:    1,
			ServiceID: 1,
			OrderID:   1,
		},
		errRes: &responses.ValidationErrResp{},
		expRes: 0,
	},
	{
		name: "missing amount",
		res: dto.Reservation{
			UserID:    1,
			ServiceID: 1,
			OrderID:   1,
		},
		errRes: &responses.ValidationErrResp{},
		expRes: 0,
	},
	{
		name: "missing service_id",
		res: dto.Reservation{
			UserID:  1,
			Amount:  1,
			OrderID: 1,
		},
		errRes: &responses.ValidationErrResp{},
	},
	{
		name: "missing order_id",
		res: dto.Reservation{
			UserID:    1,
			Amount:    1,
			ServiceID: 1,
		},
		errRes: &responses.ValidationErrResp{},
		expRes: 0,
	},
	{
		name: "no errors",
		res: dto.Reservation{
			UserID:    1,
			Amount:    1,
			ServiceID: 1,
			OrderID:   1,
		},
		errRes: nil,
		expRes: 1,
	},
}

func TestTransactionService_CreateReservation(t *testing.T) {
	for _, tt := range reservationTests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := transactionService.CreateReservation(tt.res)
			require.Equal(t, reflect.TypeOf(err), reflect.TypeOf(tt.errRes))
			require.Equal(t, id, tt.expRes)
		})
	}
}

func TestTransactionService_CreateRelease(t *testing.T) {
	for _, tt := range reservationTests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := transactionService.CreateRelease(tt.res)
			require.Equal(t, reflect.TypeOf(err), reflect.TypeOf(tt.errRes))
			require.Equal(t, id, tt.expRes)
		})
	}
}

func TestTransactionService_CancelReservation(t *testing.T) {
	for _, tt := range reservationTests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := transactionService.CancelReservation(tt.res)
			require.Equal(t, reflect.TypeOf(err), reflect.TypeOf(tt.errRes))
			require.Equal(t, id, tt.expRes)
		})
	}
}

type testReservationReport struct {
	name   string
	res    dto.ReportTime
	errRes error
	expRes []*models.ReservationReport
}

func TestTransactionService_GetReservationsReport(t *testing.T) {
	var tests = []testReservationReport{
		{
			name: "no errors",
			res: dto.ReportTime{
				Year:  2021,
				Month: 1,
			},
			errRes: nil,
			expRes: make([]*models.ReservationReport, 0),
		},
		{
			name: "year is lt 1",
			res: dto.ReportTime{
				Year:  0,
				Month: 1,
			},
			errRes: &responses.ValidationErrResp{},
			expRes: nil,
		},
		{
			name: "month is lt 1",
			res: dto.ReportTime{
				Year:  2021,
				Month: 0,
			},
			errRes: &responses.ValidationErrResp{},
			expRes: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := transactionService.GetReservationsReport(tt.res)
			require.Equal(t, reflect.TypeOf(err), reflect.TypeOf(tt.errRes))
		})
	}
}
