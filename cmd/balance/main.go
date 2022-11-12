package main

import (
	"context"
	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/fadyat/avito-internship-2022/docs"
	"github.com/fadyat/avito-internship-2022/internal/config"
	"github.com/fadyat/avito-internship-2022/internal/handlers"
	"github.com/fadyat/avito-internship-2022/internal/persistence/postgres"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
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
	cfg, err := initConfig()
	if err != nil {
		log.Fatal(err)
	}

	psql, err := initDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New(fiber.Config{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		AppName:      "avito-internship-2022",
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

	log.Panic(app.Listen(":" + cfg.HTTPPort))
}

func initConfig() (*config.HTTPConfig, error) {
	var cfg config.HTTPConfig
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("[DEBUG] Error loading .env file: %s", err)
	}

	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func initDB(cfg *config.HTTPConfig) (*pgx.Conn, error) {
	psql, err := pgx.Connect(context.Background(), cfg.GetConnectionString())
	if err != nil {
		return nil, err
	}

	return psql, nil
}
