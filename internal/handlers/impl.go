package handlers

import (
	"github.com/fadyat/avito-internship-2022/internal/persistence/postgres"
	"github.com/fadyat/avito-internship-2022/internal/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func InitRoutes(app *fiber.App, psql *pgx.Conn, log *zap.Logger, validate *validator.Validate) {
	v1 := app.Group("/api/v1")
	log.Debug("created v1 group")

	hr := postgres.NewHealthRepo(psql)
	hs := services.NewHealthService(hr)
	hh := NewHealthHandler(hs, validate)
	v1.Get("/health", hh.healthCheck)
	log.Debug("registered health check handler")

	sr := postgres.NewOuterServiceRepo(psql)
	ss := services.NewOuterServiceService(sr, validate)
	sh := NewOuterServiceHandler(ss, log)
	v1.Post("/service", sh.createService)
	v1.Get("/service", sh.getServices)
	v1.Get("/service/:id<int>", sh.getServiceByID)
	log.Debug("registered outer service handlers")

	wr := postgres.NewUserWalletRepo(psql)
	ws := services.NewUserWalletService(wr, validate)
	wh := NewUserWalletHandler(ws, log)
	v1.Post("/wallet", wh.createWallet)
	v1.Get("/wallet", wh.getWallets)
	v1.Get("/wallet/:id<int>", wh.getWalletByID)
	log.Debug("registered user wallet handlers")

	tr := postgres.NewTransactionRepo(psql)
	ts := services.NewTransactionService(tr, validate)
	th := NewTransactionHandler(ts, log)
	v1.Get("/transaction/user/:id<int>", th.getUserTransactions)
	v1.Post("/transaction/replenishment", th.createReplenishment)
	v1.Post("/transaction/withdrawal", th.createWithdrawal)
	v1.Post("/transaction/reservation", th.createReservation)
	v1.Post("/transaction/cancel", th.cancelReservation)
	v1.Post("/transaction/release", th.createRelease)
	log.Debug("registered transaction handlers")
}
