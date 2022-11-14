package handlers

import (
	"github.com/fadyat/avito-internship-2022/internal/responses"
	"github.com/fadyat/avito-internship-2022/internal/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type HealthHandler struct {
	s services.IHealthService
	v *validator.Validate
}

func NewHealthHandler(
	s services.IHealthService,
	v *validator.Validate,
) *HealthHandler {
	return &HealthHandler{s: s, v: v}
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
		return c.Status(fiber.StatusInternalServerError).JSON(&responses.ErrorResp{
			Message:     "Database connection error",
			Description: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&responses.HealthSuccess{
		Message: "OK",
	})
}
