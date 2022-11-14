package handlers

import (
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type OuterServiceHandler struct {
	s services.IOuterServiceService
}

func NewOuterServiceHandler(s services.IOuterServiceService) *OuterServiceHandler {
	return &OuterServiceHandler{s: s}
}

// createService godoc
// @tags        OuterService
// @router      /api/v1/service [post]
// @summary     New service
// @description Create new outer service info in the system
// @param       body body dto.OuterService true "Outer service short info"
// todo: add responses
func (h *OuterServiceHandler) createService(c *fiber.Ctx) error {
	// todo: return correct response + add logging

	// todo: specify responses in swagger and add them here, depending on the error
	var body dto.OuterService
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Bad request",
		})
	}

	id, err := h.s.CreateService(body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"message": "OK",
		"id":      id,
	})
}

// getServices godoc
// @tags        OuterService
// @router      /api/v1/service [get]
// @summary     Get all services
// @description Get all outer services info in the system
// todo: add responses
func (h *OuterServiceHandler) getServices(c *fiber.Ctx) error {
	// todo: return correct response + add logging
	// todo: specify responses in swagger and add them here, depending on the error
	svcs, err := h.s.GetAllServices()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message":  "OK",
		"services": svcs,
	})
}

// getServiceByID godoc
// @tags        OuterService
// @router      /api/v1/service/{id} [get]
// @summary     Get service by id
// @description Get outer service info in the system by id
// @param       id path uint64 true "service_id"
func (h *OuterServiceHandler) getServiceByID(c *fiber.Ctx) error {
	// todo: return correct response + add logging
	// todo: specify responses in swagger and add them here, depending on the error
	// todo: move parsing to service layer
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	svc, err := h.s.GetServiceByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "OK",
		"service": svc,
	})
}
