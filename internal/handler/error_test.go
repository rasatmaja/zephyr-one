package handler

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/zephyr-one/internal/config"
	"github.com/rasatmaja/zephyr-one/internal/logger"
	"github.com/rasatmaja/zephyr-one/internal/password"
	"github.com/rasatmaja/zephyr-one/internal/response"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	// setup handler
	handler := &Endpoint{
		log:      logger.New(),
		password: password.Factory(),
	}

	// setup fiber app
	app := fiber.New(
		fiber.Config{
			ErrorHandler: handler.Error,
		},
	)
	env := config.LoadENV()
	env.LogLevel = "disable" // disable logging

	t.Run("erros-is-response-struct", func(t *testing.T) {
		app.Get("/error", func(c *fiber.Ctx) error {
			return response.Factory().BadRequest("Testing")
		})
		resp, err := app.Test(httptest.NewRequest("GET", "/error", nil))

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

	})

	t.Run("erros-isnt-response-struct", func(t *testing.T) {
		app.Get("/errors", func(c *fiber.Ctx) error {
			return fmt.Errorf("Testing")
		})
		resp, err := app.Test(httptest.NewRequest("GET", "/errors", nil))

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

	})
}
