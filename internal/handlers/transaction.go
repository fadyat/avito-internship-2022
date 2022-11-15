package handlers

import (
	"errors"
	"fmt"
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
// @summary     Replenishment of the user's balance
// @description Replenishment of the user's balance by a certain amount and creating a replenishment transaction
// @param       body body     dto.Replenishment true "Replenishment info"
// @response    400  {object} responses.ErrorResp
// @response    422  {object} responses.ErrorResp
// @response    500  {object} responses.ErrorResp
// response    201  {object} // todo: implement
func (h *TransactionHandler) createReplenishment(c *fiber.Ctx) error {
	var body dto.Replenishment
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

	// todo: fix response
	msg := fmt.Sprintf("Replenishment %d created", id)
	return c.Status(fiber.StatusCreated).JSON(`{"message": "` + msg + `"}`)
}

// createWithdrawal godoc
// @tags        Transaction
// @router      /api/v1/transaction/withdrawal [post]
// @summary     Withdrawal of the user's balance
// @description Withdrawal of the user's balance by a certain amount and creating a withdrawal transaction
// @param       body body     dto.Withdrawal true "Withdrawal info"
// @response    400  {object} responses.ErrorResp
// @response    422  {object} responses.ErrorResp
// @response    500  {object} responses.ErrorResp
// response    201  {object} // todo: implement
func (h *TransactionHandler) createWithdrawal(c *fiber.Ctx) error {
	// todo: implement
	panic("implement me")
}

func (h *TransactionHandler) createReservation(c *fiber.Ctx) error {
	// todo: implement
	panic("implement me")
}

func (h *TransactionHandler) createRelease(c *fiber.Ctx) error {
	// todo: implement
	panic("implement me")
}

func (h *TransactionHandler) getTransactions(c *fiber.Ctx) error {
	// todo: implement
	panic("implement me")
}

func (h *TransactionHandler) getTransactionByID(c *fiber.Ctx) error {
	// todo: implement
	panic("implement me")
}
