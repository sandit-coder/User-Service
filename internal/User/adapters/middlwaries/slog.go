package middleware

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func NewSlogMiddleware(log *slog.Logger) fiber.Handler {
	return func(c fiber.Ctx) error {
		reqID := uuid.NewString()

		reqLog := log.With(
			slog.String("request_id", reqID),
			slog.String("method", c.Method()),
			slog.String("path", c.Path()),
		)

		c.Locals("logger", reqLog)

		reqLog.Info("request started")

		err := c.Next()

		status := c.Response().StatusCode()

		reqLog.Info("request finished",
			"status", status,
		)

		return err
	}
}
