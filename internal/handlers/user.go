package handlers

import (
	"errors"
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/persistence"
	"github.com/fadyat/avito-internship-2022/internal/responses"
	"github.com/fadyat/avito-internship-2022/internal/services"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UserWalletHandler struct {
	s *services.UserWalletService
	l *zap.Logger
}

func NewUserWalletHandler(
	s *services.UserWalletService,
	l *zap.Logger,
) *UserWalletHandler {
	return &UserWalletHandler{s: s, l: l}
}

// createWallet godoc
// @tags        UserWallet
// @router      /api/v1/wallet [post]
// @summary     New wallet
// @description Create new wallet in the system
// @param       body body     dto.UserWallet true "Wallet info"
// @response    201  {object} responses.UserWalletCreated
// @response    400  {object} responses.ErrorResp
// @response    422  {object} responses.ErrorResp
// @response    500  {object} responses.ErrorResp
func (h *UserWalletHandler) createWallet(c *fiber.Ctx) error {
	var body dto.UserWallet
	if err := c.BodyParser(&body); err != nil {
		h.l.Debug("failed to parse body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(&responses.ErrorResp{
			Message:     "Bad request",
			Description: err.Error(),
		})
	}

	id, err := h.s.CreateUserWallet(body)
	var verr *responses.ValidationErrResp
	if errors.As(err, &verr) {
		h.l.Debug("validation failed", zap.Error(err))
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&responses.ErrorResp{
			Message:     "Validation error",
			Description: err.Error(),
		})
	}

	if errors.Is(err, persistence.ErrUniqueViolation) {
		h.l.Debug("wallet already exists", zap.Error(err))
		return c.Status(fiber.StatusConflict).JSON(&responses.ErrorResp{
			Message:     "Conflict",
			Description: err.Error(),
		})
	}

	if err != nil {
		h.l.Error("failed to create wallet", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(&responses.ErrorResp{
			Message:     "Internal server error",
			Description: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&responses.UserWalletCreated{
		ID: id,
	})
}

// getWallets godoc
// @tags        UserWallet
// @router      /api/v1/wallet [get]
// @summary     Get wallets
// @description Get wallets from the system
// @response    200 {object} responses.UserWallets
// @response    500 {object} responses.ErrorResp
func (h *UserWalletHandler) getWallets(c *fiber.Ctx) error {
	ws, err := h.s.GetAllUserWallets()

	if err != nil {
		h.l.Error("failed to get wallets", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(&responses.ErrorResp{
			Message:     "Internal server error",
			Description: err.Error(),
		})
	}

	if ws == nil {
		ws = make([]*models.UserWallet, 0)
	}

	return c.Status(fiber.StatusOK).JSON(&responses.UserWallets{
		Wallets: ws,
	})
}

// getWallet godoc
// @tags        UserWallet
// @router      /api/v1/wallet/{id} [get]
// @summary     Get wallet
// @description Get user wallet from the system by id
// @param       id  path     int true "Wallet id"
// @response    200 {object} models.UserWallet
// @response    400 {object} responses.ErrorResp
// @response    404 {object} responses.ErrorResp
// @response    422 {object} responses.ErrorResp
// @response    500 {object} responses.ErrorResp
func (h *UserWalletHandler) getWalletByID(c *fiber.Ctx) error {
	w, err := h.s.GetUserWalletByID(c.Params("id"))
	var verr *responses.ValidationErrResp
	if errors.As(err, &verr) {
		h.l.Debug("validation failed", zap.Error(err))
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&responses.ErrorResp{
			Message:     "Validation error",
			Description: err.Error(),
		})
	}

	if errors.Is(err, persistence.ErrNotFound) {
		h.l.Debug("no wallet found", zap.Error(err))
		return c.Status(fiber.StatusNotFound).JSON(&responses.ErrorResp{
			Message:     "Not found",
			Description: err.Error(),
		})
	}

	if err != nil {
		h.l.Error("failed to get wallet", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(&responses.ErrorResp{
			Message:     "Internal server error",
			Description: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(w)
}
