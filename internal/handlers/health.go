package handlers

import (
	"github.com/fadyat/avito-internship-2022/internal/persistence"
	"github.com/fadyat/avito-internship-2022/internal/responses"
	"github.com/gofiber/fiber/v2"
)

type HealthHandler struct {
	repo persistence.HealthRepository
}

func NewHealthHandler(repo persistence.HealthRepository) *HealthHandler {
	return &HealthHandler{
		repo: repo,
	}
}

// HealthCheck godoc
// @Router      /api/v1/health [get]
// @Tags        health
// @Summary     Healthcheck
// @Description Healthcheck endpoint, that checks if the service is alive and database connection is working.
// @Response    200 {object} responses.HealthSuccess
// @Response    500 {object} responses.HealthError
func (h *HealthHandler) HealthCheck(c *fiber.Ctx) error {
	if err := h.repo.Ping(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&responses.HealthError{
			Message:     "Database connection error",
			Description: err.Error(),
			StatusCode:  fiber.StatusInternalServerError,
		})
	}

	return c.Status(fiber.StatusOK).JSON(&responses.HealthSuccess{
		Message:    "OK",
		StatusCode: fiber.StatusOK,
	})
}
