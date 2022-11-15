package handlers

import (
	"errors"
	"fmt"
	"github.com/fadyat/avito-internship-2022/internal/models"
	"github.com/fadyat/avito-internship-2022/internal/models/dto"
	"github.com/fadyat/avito-internship-2022/internal/persistence"
	"github.com/fadyat/avito-internship-2022/internal/responses"
	"github.com/fadyat/avito-internship-2022/internal/services"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type OuterServiceHandler struct {
	s services.IOuterServiceService
	l *zap.Logger
}

func NewOuterServiceHandler(
	s services.IOuterServiceService,
	l *zap.Logger,
) *OuterServiceHandler {
	return &OuterServiceHandler{s: s, l: l}
}

// createService godoc
// @tags        OuterService
// @router      /api/v1/service [post]
// @summary     New service
// @description Create new outer service info in the system
// @param       body body     dto.OuterService true "Outer service short info"
// @response    201  {object} responses.ServiceCreated
// @response    400  {object} responses.ErrorResp
// @response    422  {object} responses.ErrorResp
// @response    500  {object} responses.ErrorResp
func (h *OuterServiceHandler) createService(c *fiber.Ctx) error {
	var body dto.OuterService
	if err := c.BodyParser(&body); err != nil {
		h.l.Debug("failed to parse body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(&responses.ErrorResp{
			Message:     "Bad request",
			Description: err.Error(),
		})
	}

	id, err := h.s.CreateService(body)
	var verr *responses.ValidationErrResp
	if errors.As(err, &verr) {
		h.l.Debug("validation failed", zap.Error(err))
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&responses.ErrorResp{
			Message:     "Unprocessable entity",
			Description: err.Error(),
		})
	}

	if err != nil {
		h.l.Error("failed to create service", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(&responses.ErrorResp{
			Message:     "Internal server error",
			Description: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&responses.ServiceCreated{
		ID: id,
	})
}

// getServices godoc
// @tags        OuterService
// @router      /api/v1/service [get]
// @summary     Get all services
// @description Get all outer services info in the system
// @response    200 {object} responses.Services
// @response    500 {object} responses.ErrorResp
func (h *OuterServiceHandler) getServices(c *fiber.Ctx) error {
	svcs, err := h.s.GetAllServices()

	if err != nil {
		h.l.Error("failed to get all services", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(&responses.ErrorResp{
			Message:     "Internal server error",
			Description: err.Error(),
		})
	}

	if svcs == nil {
		svcs = make([]*models.OuterService, 0)
	}

	return c.Status(fiber.StatusOK).JSON(&responses.Services{
		Services: svcs,
	})
}

// getServiceByID godoc
// @tags        OuterService
// @router      /api/v1/service/{id} [get]
// @summary     Get service by id
// @description Get outer service info in the system by id
// @param       id  path     uint64 true "service_id"
// @response    200 {object} models.OuterService
// @response    400 {object} responses.ErrorResp
// @response    404 {object} responses.ErrorResp
// @response    422 {object} responses.ErrorResp
// @response    500 {object} responses.ErrorResp
func (h *OuterServiceHandler) getServiceByID(c *fiber.Ctx) error {
	svc, err := h.s.GetServiceByID(c.Params("id"))
	var verr *responses.ValidationErrResp
	if errors.As(err, &verr) {
		h.l.Debug("validation failed", zap.Error(err))
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&responses.ErrorResp{
			Message:     "Unprocessable entity",
			Description: err.Error(),
		})
	}

	if errors.Is(err, persistence.ErrNotFound) {
		h.l.Debug("service not found", zap.Error(err))
		return c.Status(fiber.StatusNotFound).JSON(&responses.ErrorResp{
			Message:     "Not found",
			Description: fmt.Sprintf("service with id %s not found", c.Params("id")),
		})
	}

	if err != nil {
		h.l.Error("failed to get service by id", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(&responses.ErrorResp{
			Message:     "Internal server error",
			Description: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(svc)
}
