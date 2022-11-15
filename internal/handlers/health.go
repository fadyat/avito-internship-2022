package handlers

import (
	"github.com/fadyat/avito-internship-2022/internal/responses"
	"github.com/fadyat/avito-internship-2022/internal/services"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type HealthHandler struct {
	s services.IHealthService
	l *zap.Logger
}

func NewHealthHandler(
	s services.IHealthService,
	l *zap.Logger,
) *HealthHandler {
	return &HealthHandler{s: s, l: l}
}

// healthCheck godoc
// @router      /api/v1/health [get]
// @tags        health
// @summary     Healthcheck
// @description Healthcheck endpoint, that checks if the service is alive and database connection is working.
// @response    200 {object} responses.HealthSuccess
// @response    500 {object} responses.ErrorResp
func (h *HealthHandler) healthCheck(c *fiber.Ctx) error {
	if err := h.s.Ping(); err != nil {
		h.l.Debug("failed to ping database", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(&responses.ErrorResp{
			Message:     "Database connection error",
			Description: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&responses.HealthSuccess{
		Message: "OK",
	})
}
