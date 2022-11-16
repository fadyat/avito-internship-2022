package services

import (
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/responses"
)

func (s *TransactionService) CreateReservation(tr dto.Reservation) (uint64, error) {
	if err := s.v.Struct(tr); err != nil {
		return 0, &responses.ValidationErrResp{Message: err.Error()}
	}

	return s.r.CreateReservation(tr)
}

func (s *TransactionService) CreateRelease(tr dto.Reservation) (uint64, error) {
	if err := s.v.Struct(tr); err != nil {
		return 0, &responses.ValidationErrResp{Message: err.Error()}
	}

	return s.r.CreateRelease(tr)
}

func (s *TransactionService) CancelReservation(tr dto.Reservation) (uint64, error) {
	if err := s.v.Struct(tr); err != nil {
		return 0, &responses.ValidationErrResp{Message: err.Error()}
	}

	return s.r.CancelReservation(tr)
}

func (s *TransactionService) GetReservationsReport(tm dto.ReportTime) ([]*models.ReservationReport, error) {
	if err := s.v.Struct(tm); err != nil {
		return nil, &responses.ValidationErrResp{Message: err.Error()}
	}

	return s.r.GetReservationsReport(tm)
}
