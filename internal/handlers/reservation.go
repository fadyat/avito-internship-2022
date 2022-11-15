package handlers

import (
	"errors"
	"github.com/fadyat/avito-internship-2022/internal/helpers"
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/persistence"
	"github.com/fadyat/avito-internship-2022/internal/responses"
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

	id, err := h.s.CreateReservation(body)
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

	id, err := h.s.CreateRelease(body)
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
// @router      /api/v1/transaction/reservation [delete]
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

	id, err := h.s.CancelReservation(body)
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

// getUserTransactions godoc
// @tags        Transaction
// @router      /api/v1/transaction/user/{id} [get]
// @summary     Get user transactions
// @description Get all user transactions in paginated form
// @param       id       path     int    true  "User ID"     Format(uint64)
// @param       page     query    int    false "Page number" default(1)
// @param       per_page query    int    false "Page size"   default(10)
// @param       order_by query    string false "Order by"    Enums(created_at, amount) default(created_at, amount)
// @response    200      {object} responses.TransactionPaginated
// @response    400      {object} responses.ErrorResp
// @response    422      {object} responses.ErrorResp
// @response    500      {object} responses.ErrorResp
func (h *TransactionHandler) getUserTransactions(c *fiber.Ctx) error {
	var pag models.Pagination
	if err := c.QueryParser(&pag); err != nil {
		h.l.Debug("failed to parse query", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(&responses.ErrorResp{
			Message:     "Bad request",
			Description: err.Error(),
		})
	}
	helpers.MakePaginationParamsCorrect(&pag)

	ts, err := h.s.GetUserTransactions(c.Params("id"), pag.Page, pag.PerPage, pag.OrderBy)
	verr := &responses.ValidationErrResp{}
	if errors.As(err, &verr) {
		h.l.Debug("validation failed", zap.Error(err))
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&responses.ErrorResp{
			Message:     "Validation error",
			Description: err.Error(),
		})
	}

	// fixme: it will be better to handle this error in service layer, but for rn it's ok
	if errors.Is(err, persistence.ErrInvalidColumn) {
		h.l.Debug("invalid column", zap.Error(err))
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&responses.ErrorResp{
			Message:     "Validation error",
			Description: err.Error(),
		})
	}

	if err != nil {
		h.l.Error("failed to get user transactions", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(&responses.ErrorResp{
			Message:     "Internal server error",
			Description: err.Error(),
		})
	}

	if ts == nil {
		ts = make([]*models.Transaction, 0)
	}

	total, err := h.s.GetUserTransactionsCount(c.Params("id"))
	if err != nil {
		h.l.Error("failed to get user transactions count", zap.Error(err))
	}

	return c.Status(fiber.StatusOK).JSON(&responses.TransactionPaginated{
		Transactions: ts,
		Pagination:   responses.NewPagination(pag.PerPage, pag.Page, uint64(len(ts)), total),
	})
}
