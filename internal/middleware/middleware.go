package middleware

import (
	"time"
	"user-api/internal/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func RequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := uuid.New().String()
		c.Set("X-Request-ID", id)
		c.Locals("requestid", id)
		return c.Next()
	}
}

func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start)

		id := c.Get("X-Request-ID")
		// If not set by middleware? (Should be set if chained correctly)

		status := c.Response().StatusCode()

		logger.Log.Info("Request",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", status),
			zap.Duration("duration", duration),
			zap.String("request_id", id),
		)

		return err
	}
}
