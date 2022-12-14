package handlers

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/persistence"
	"github.com/fadyat/avito-internship-2022/internal/responses"
	"github.com/gocarina/gocsv"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// createReservation godoc
// @tags        Transaction
// @router      /api/v1/transaction/reservation [post]
// @summary     Reservation of the user's balance
// @description Reservation of the user's balance from another service
// @param       body body     dto.Reservation true "Reservation info"
// @response    201  {object} responses.TransactionCreated
// @response    400  {object} responses.ErrorResp
// @response    404  {object} responses.ErrorResp
// @response    422  {object} responses.ErrorResp
// @response    500  {object} responses.ErrorResp
func (h *TransactionHandler) createReservation(c *fiber.Ctx) error {
	var body dto.Reservation
	if err := c.BodyParser(&body); err != nil {
		h.l.Debug("failed to parse body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(&responses.ErrorResp{
			Message:     "Bad request",
			Description: err.Error(),
		})
	}

	id, err := h.s.CreateReservation(&body)
	var verr *responses.ValidationErrResp
	if errors.As(err, &verr) {
		h.l.Debug("validation failed", zap.Error(err))
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&responses.ErrorResp{
			Message:     "Validation error",
			Description: err.Error(),
		})
	}

	if errors.Is(err, persistence.ErrForeignKeyViolation) {
		h.l.Debug("service or wallet does not exist", zap.Error(err))
		return c.Status(fiber.StatusNotFound).JSON(&responses.ErrorResp{
			Message:     "Service or user wallet not found",
			Description: err.Error(),
		})
	}

	if err != nil {
		h.l.Error("failed to create reservation", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(&responses.ErrorResp{
			Message:     "Internal server error",
			Description: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&responses.TransactionCreated{
		ID: id,
	})
}

// createRelease godoc
// @tags        Transaction
// @router      /api/v1/transaction/release [post]
// @summary     Release of the user's balance
// @description Release of the user's balance to another service
// @param       body body     dto.Reservation true "Release info"
// @response    201  {object} responses.ReservationReleased
// @response    400  {object} responses.ErrorResp
// @response    404  {object} responses.ErrorResp
// @response    422  {object} responses.ErrorResp
// @response    500  {object} responses.ErrorResp
func (h *TransactionHandler) createRelease(c *fiber.Ctx) error {
	var body dto.Reservation
	if err := c.BodyParser(&body); err != nil {
		h.l.Debug("failed to parse body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(&responses.ErrorResp{
			Message:     "Bad request",
			Description: err.Error(),
		})
	}

	id, err := h.s.CreateRelease(&body)
	var verr *responses.ValidationErrResp
	if errors.As(err, &verr) {
		h.l.Debug("validation failed", zap.Error(err))
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&responses.ErrorResp{
			Message:     "Validation error",
			Description: err.Error(),
		})
	}

	if errors.Is(err, persistence.ErrForeignKeyViolation) {
		h.l.Debug("service or wallet does not exist", zap.Error(err))
		return c.Status(fiber.StatusNotFound).JSON(&responses.ErrorResp{
			Message:     "Service or user wallet not found",
			Description: err.Error(),
		})
	}

	if errors.Is(err, persistence.ErrInsufficientFunds) {
		h.l.Debug("insufficient funds", zap.Error(err))
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&responses.ErrorResp{
			Message:     "Insufficient funds",
			Description: err.Error(),
		})
	}

	if errors.Is(err, persistence.ErrNotFound) {
		h.l.Debug("reservation not found", zap.Error(err))
		return c.Status(fiber.StatusNotFound).JSON(&responses.ErrorResp{
			Message:     "Reservation not found",
			Description: err.Error(),
		})
	}

	if err != nil {
		h.l.Error("failed to create release", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(&responses.ErrorResp{
			Message:     "Internal server error",
			Description: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&responses.ReservationReleased{
		ID: id,
	})
}

// cancelReservation godoc
// @tags        Transaction
// @router      /api/v1/transaction/cancel [post]
// @summary     Cancel reservation of the user's balance
// @description Cancel reservation of the user's balance from another service
// @param       body body     dto.Reservation true "Reservation info"
// @response    200  {object} responses.ReservationCancelled
// @response    400  {object} responses.ErrorResp
// @response    404  {object} responses.ErrorResp
// @response    422  {object} responses.ErrorResp
// @response    500  {object} responses.ErrorResp
func (h *TransactionHandler) cancelReservation(c *fiber.Ctx) error {
	var body dto.Reservation
	if err := c.BodyParser(&body); err != nil {
		h.l.Debug("failed to parse body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(&responses.ErrorResp{
			Message:     "Bad request",
			Description: err.Error(),
		})
	}

	id, err := h.s.CancelReservation(&body)
	var verr *responses.ValidationErrResp
	if errors.As(err, &verr) {
		h.l.Debug("validation failed", zap.Error(err))
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&responses.ErrorResp{
			Message:     "Validation error",
			Description: err.Error(),
		})
	}

	if errors.Is(err, persistence.ErrForeignKeyViolation) {
		h.l.Debug("service or wallet does not exist", zap.Error(err))
		return c.Status(fiber.StatusNotFound).JSON(&responses.ErrorResp{
			Message:     "Service or user wallet not found",
			Description: err.Error(),
		})
	}

	if errors.Is(err, persistence.ErrNotFound) {
		h.l.Debug("reservation not found", zap.Error(err))
		return c.Status(fiber.StatusNotFound).JSON(&responses.ErrorResp{
			Message:     "Reservation not found",
			Description: err.Error(),
		})
	}

	if err != nil {
		h.l.Error("failed to cancel reservation", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(&responses.ErrorResp{
			Message:     "Internal server error",
			Description: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&responses.ReservationCancelled{
		ID: id,
	})
}

// getReservationReport godoc
// @tags        Transaction
// @router      /api/v1/transaction/reservation/report [get]
// @summary     Get reservation report
// @description Get reservation report
// @produce     json,text/csv
// @param       body   body     dto.ReportTime true  "Reservation report time"
// @param       format query    string         false "Report format" Enums(csv, json)
// @response    400    {object} responses.ErrorResp
// @response    422    {object} responses.ErrorResp
// @response    500    {object} responses.ErrorResp
func (h *TransactionHandler) getReservationReport(ctx *fiber.Ctx) error {
	var body dto.ReportTime
	if err := ctx.BodyParser(&body); err != nil {
		h.l.Debug("failed to parse body", zap.Error(err))
		return ctx.Status(fiber.StatusBadRequest).JSON(&responses.ErrorResp{
			Message:     "Bad request",
			Description: err.Error(),
		})
	}

	r, err := h.s.GetReservationsReport(&body)
	var verr *responses.ValidationErrResp
	if errors.As(err, &verr) {
		h.l.Debug("validation failed", zap.Error(err))
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&responses.ErrorResp{
			Message:     "Validation error",
			Description: err.Error(),
		})
	}

	if err != nil {
		h.l.Error("failed to get reservation report", zap.Error(err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(&responses.ErrorResp{
			Message:     "Internal server error",
			Description: err.Error(),
		})
	}

	format := ctx.Query("format", "csv")
	if r == nil {
		r = make([]*models.ReservationReport, 0)
	}

	if format == "csv" {
		filename := fmt.Sprintf("reservation_report_%d-%d.csv", body.Year, body.Month)
		ctx.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
		ctx.Set("Content-Type", "text/csv")
		w := csv.NewWriter(ctx)
		if err := gocsv.MarshalCSV(r, w); err != nil {
			h.l.Error("failed to marshal csv", zap.Error(err))
			return ctx.Status(fiber.StatusInternalServerError).JSON(&responses.ErrorResp{
				Message:     "Internal server error",
				Description: err.Error(),
			})
		}

		w.Flush()
		if err := w.Error(); err != nil {
			h.l.Error("failed to flush csv", zap.Error(err))
			return ctx.Status(fiber.StatusInternalServerError).JSON(&responses.ErrorResp{
				Message:     "Internal server error",
				Description: err.Error(),
			})
		}

		return nil
	}

	return ctx.Status(fiber.StatusOK).JSON(r)
}
