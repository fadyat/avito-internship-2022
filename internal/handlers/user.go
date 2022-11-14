package handlers

import (
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
	// todo: implement
	return nil
}

// getWallets godoc
// @tags        UserWallet
// @router      /api/v1/wallet [get]
// @summary     Get wallets
// @description Get wallets from the system
// @response    200 {object} responses.UserWallets
// @response    404 {object} responses.ErrorResp // todo: remove??
// @response    500 {object} responses.ErrorResp
func (h *UserWalletHandler) getWallets(c *fiber.Ctx) error {
	// todo: implement
	return nil
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
	// todo: implement
	return nil
}
