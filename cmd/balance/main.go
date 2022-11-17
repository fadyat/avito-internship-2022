package main

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/fadyat/avito-internship-2022/docs"
	"github.com/fadyat/avito-internship-2022/internal/config"
	"github.com/fadyat/avito-internship-2022/internal/handlers"
	"github.com/fadyat/avito-internship-2022/internal/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

// @title        Avito Internship 2022 Balance API
// @description  This is a sample server for a balance API.
// @version      1.0.0
// @host         localhost:80
// @schemes      http
// @consumes     application/json
// @produces     application/json
// @contact.name Artyom Fadeyev
// @contact.url  https://github.com/fadyat
func main() {
	validate := initValidator()

	logs := initLogger()
	defer func() { _ = logs.Sync() }()

	cfg, err := initConfig(logs, validate)
	if err != nil {
		logs.Fatal("failed to init config", zap.Error(err))
	}

	psql, err := initDB(cfg)
	defer func() {
		if err = psql.Close(); err != nil {
			logs.Fatal("failed to close db connection", zap.Error(err))
		}
	}()

	if err != nil {
		logs.Fatal("failed to init db", zap.Error(err))
	}
	logs.Debug("db initialized")

	app := initApp(logs)
	logs.Debug("app initialized")

	handlers.InitRoutes(app, psql, logs, validate)
	logs.Debug("routes initialized")

	err = app.Listen(":" + cfg.HTTPPort)
	defer func() {
		if err = app.Shutdown(); err != nil {
			logs.Fatal("failed to shutdown server", zap.Error(err))
		}
	}()
	if err != nil {
		logs.Fatal("failed to start server", zap.Error(err))
	}
}

func initConfig(logs *zap.Logger, validate *validator.Validate) (*config.HTTPConfig, error) {
	var cfg config.HTTPConfig
	if err := godotenv.Load(".env"); err != nil {
		logs.Debug("failed to load .env file", zap.Error(err))
	}

	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	err := validate.Struct(cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func initDB(cfg *config.HTTPConfig) (*sqlx.DB, error) {
	psql, err := sqlx.Connect("postgres", cfg.GetConnectionString())
	if err != nil {
		return nil, err
	}

	return psql, nil
}

func initLogger() *zap.Logger {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	logLevel := zap.NewAtomicLevelAt(zap.DebugLevel)

	logs := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		logLevel,
	))
	return logs
}

func initApp(l *zap.Logger) *fiber.App {
	app := fiber.New(fiber.Config{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		AppName:      "avito-internship-2022",
	})

	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	app.Use(middleware.IncludeLogs(l))

	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:         "/swagger/doc.json",
		DeepLinking: true,
	}))

	return app
}

func initValidator() *validator.Validate {
	return validator.New()
}
