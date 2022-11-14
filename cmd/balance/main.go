package main

import (
	"context"
	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/fadyat/avito-internship-2022/docs"
	"github.com/fadyat/avito-internship-2022/internal/config"
	"github.com/fadyat/avito-internship-2022/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
	logger := initLogger()
	defer func() {
		_ = logger.Sync()
	}()

	cfg, err := initConfig(logger)
	if err != nil {
		logger.Fatal("failed to init config", zap.Error(err))
	}

	psql, err := initDB(cfg)
	defer func() {
		if err = psql.Close(context.Background()); err != nil {
			logger.Fatal("failed to close db connection", zap.Error(err))
		}
	}()
	if err != nil {
		logger.Fatal("failed to init db", zap.Error(err))
	}
	logger.Debug("db initialized")

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

	handlers.InitRoutes(app, psql)

	err = app.Listen(":" + cfg.HTTPPort)
	defer func() {
		if err = app.Shutdown(); err != nil {
			logger.Fatal("failed to shutdown server", zap.Error(err))
		}
	}()
	if err != nil {
		logger.Fatal("failed to start server", zap.Error(err))
	}
}

func initConfig(logger *zap.Logger) (*config.HTTPConfig, error) {
	var cfg config.HTTPConfig
	if err := godotenv.Load(".env"); err != nil {
		logger.Debug("failed to load .env file", zap.Error(err))
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

func initLogger() *zap.Logger {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	logLevel := zap.NewAtomicLevelAt(zap.DebugLevel)

	loggerCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		logLevel,
	)
	logger := zap.New(loggerCore)
	return logger
}
