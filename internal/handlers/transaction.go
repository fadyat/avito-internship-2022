package handlers

import (
	"errors"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/persistence"
	"github.com/fadyat/avito-internship-2022/internal/responses"
	"github.com/fadyat/avito-internship-2022/internal/services"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type TransactionHandler struct {
	s services.ITransactionService
	l *zap.Logger
}

func NewTransactionHandler(
	s services.ITransactionService,
	l *zap.Logger,
) *TransactionHandler {
	return &TransactionHandler{s: s, l: l}
}

// createTransaction godoc
// @tags        Transaction
// @router      /api/v1/transaction/replenishment [post]
// @summary     Transaction of the user's balance
// @description Transaction of the user's balance by a certain amount and creating a replenishment transaction
// @param       body body     dto.Transaction true "Transaction info"
// @response    201  {object} responses.TransactionCreated
// @response    400  {object} responses.ErrorResp
// @response    404  {object} responses.ErrorResp
// @response    422  {object} responses.ErrorResp
// @response    500  {object} responses.ErrorResp
func (h *TransactionHandler) createReplenishment(c *fiber.Ctx) error {
	var body dto.Transaction
	if err := c.BodyParser(&body); err != nil {
		h.l.Debug("failed to parse body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(&responses.ErrorResp{
			Message:     "Bad request",
			Description: err.Error(),
		})
	}

	id, err := h.s.CreateReplenishment(body)
	var verr *responses.ValidationErrResp
	if errors.As(err, &verr) {
		h.l.Debug("validation failed", zap.Error(err))
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&responses.ErrorResp{
			Message:     "Validation error",
			Description: err.Error(),
		})
	}

	// it will be better to handle this error in service layer, but for rn it's ok
	if errors.Is(err, persistence.ErrForeignKeyViolation) {
		h.l.Debug("wallet does not exist", zap.Error(err))
		return c.Status(fiber.StatusNotFound).JSON(&responses.ErrorResp{
			Message:     "User wallet not found",
			Description: err.Error(),
		})
	}

	if err != nil {
		h.l.Error("failed to create replenishment", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(&responses.ErrorResp{
			Message:     "Internal server error",
			Description: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&responses.TransactionCreated{
		ID: id,
	})
}

// createWithdrawal godoc
// @tags        Transaction
// @router      /api/v1/transaction/withdrawal [post]
// @summary     Transaction of the user's balance
// @description Transaction of the user's balance by a certain amount and creating a withdrawal transaction
// @param       body body     dto.Transaction true "Transaction info"
// @response    201  {object} responses.TransactionCreated
// @response    400  {object} responses.ErrorResp
// @response    404  {object} responses.ErrorResp
// @response    422  {object} responses.ErrorResp
// @response    500  {object} responses.ErrorResp
func (h *TransactionHandler) createWithdrawal(c *fiber.Ctx) error {
	var body dto.Transaction
	if err := c.BodyParser(&body); err != nil {
		h.l.Debug("failed to parse body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(&responses.ErrorResp{
			Message:     "Bad request",
			Description: err.Error(),
		})
	}

	id, err := h.s.CreateWithdrawal(body)
	var verr *responses.ValidationErrResp
	if errors.As(err, &verr) {
		h.l.Debug("validation failed", zap.Error(err))
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&responses.ErrorResp{
			Message:     "Validation error",
			Description: err.Error(),
		})
	}

	if errors.Is(err, persistence.ErrForeignKeyViolation) {
		h.l.Debug("wallet does not exist", zap.Error(err))
		return c.Status(fiber.StatusNotFound).JSON(&responses.ErrorResp{
			Message:     "User wallet not found",
			Description: err.Error(),
		})
	}

	if errors.Is(err, persistence.ErrInsufficientFunds) {
		h.l.Debug("insufficient funds", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(&responses.ErrorResp{
			Message:     "Insufficient funds",
			Description: err.Error(),
		})
	}

	if err != nil {
		h.l.Error("failed to create withdrawal", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(&responses.ErrorResp{
			Message:     "Internal server error",
			Description: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&responses.TransactionCreated{
		ID: id,
	})
}
