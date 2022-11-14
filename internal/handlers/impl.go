package handlers

import (
	"github.com/fadyat/avito-internship-2022/internal/persistence/postgres"
	"github.com/fadyat/avito-internship-2022/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func InitRoutes(app *fiber.App, psql *pgx.Conn, log *zap.Logger) {
	v1 := app.Group("/api/v1")
	log.Debug("created v1 group")

	hr := postgres.NewHealthRepository(psql)
	hs := services.NewHealthService(hr)
	hh := NewHealthHandler(hs)
	v1.Get("/health", hh.healthCheck)
	log.Debug("registered health check handler")

	sr := postgres.NewOuterServiceRepository(psql)
	ss := services.NewOuterServiceService(sr)
	sh := NewOuterServiceHandler(ss, *log)
	v1.Post("/service", sh.createService)
	v1.Get("/service", sh.getServices)
	v1.Get("/service/:id<int>", sh.getServiceByID)
	log.Debug("registered outer service handlers")

}
