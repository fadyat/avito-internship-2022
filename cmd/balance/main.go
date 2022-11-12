package main

import (
	"context"
	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/fadyat/avito-internship-2022/docs"
	"github.com/fadyat/avito-internship-2022/internal/handlers"
	"github.com/fadyat/avito-internship-2022/internal/persistence/postgres"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"time"
)

// @title        Avito Internship 2022 Balance API
// @description  This is a sample server for a balance API.
// @version      1.0.0
// @host         localhost:80
// @BasePath     /api/v1
// @schemes      http
// @consumes     application/json
// @produces     application/json
// @contact.name Artyom Fadeyev
// @contact.url  https://github.com/fadyat
func main() {
	// fixme: move to config
	dbDSN := os.Getenv("DB_DSN")
	if dbDSN == "" {
		dbDSN = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	}

	psql, err := pgx.Connect(context.Background(), dbDSN)
	if err != nil {
		log.Fatal(err)
	}

	// fixme: add more config
	app := fiber.New(fiber.Config{
		ReadTimeout: 5 * time.Second,
		AppName:     "avito-internship-2022",
	})

	app.Use(cors.New())
	app.Use(recover.New())
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:         "/swagger/doc.json",
		DeepLinking: true,
	}))

	v1 := app.Group("/api/v1")

	// fixme: make cleaner
	healthRepo := postgres.NewHealthRepository(psql)
	healthHandler := handlers.NewHealthHandler(healthRepo)
	v1.Get("/health", healthHandler.HealthCheck)

	// fixme: move to config
	port := os.Getenv("PORT")
	log.Panic(app.Listen(":" + port))
}
