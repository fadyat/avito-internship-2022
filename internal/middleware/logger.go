package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func IncludeLogs(l *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		l.Debug("request", zap.String("method", c.Method()), zap.String("path", c.Path()))
		return c.Next()
	}
}
