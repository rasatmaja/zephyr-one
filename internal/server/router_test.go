package server

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/zephyr-one/internal/logger"
)

func TestInitializeRoute(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := &App{
			server: fiber.New(),
			logger: logger.New(),
		}

		app.InitializeRoute()
	})
}
