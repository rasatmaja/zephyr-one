package handler

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/zephyr-one/internal/config"
	"github.com/rasatmaja/zephyr-one/internal/logger"
	"github.com/rasatmaja/zephyr-one/internal/password"
)

func TestRegistration(t *testing.T) {
	// setup fiber app
	app := fiber.New()
	env := config.LoadENV()
	env.LogLevel = "disable" // disable logging

	// setup handler
	handler := &Endpoint{
		log:      logger.New(),
		password: password.Factory(),
	}

	t.Run("success", func(t *testing.T) {
		app.Post("/register", handler.Regitration)
		resp, err := app.Test(httptest.NewRequest("POST", "/register", nil))

		// begin assert response from http test
		if err != nil {
			t.Error("error should be nil")
			t.Fail()
		}

		if resp.StatusCode != fiber.StatusCreated {
			t.Errorf("status code should be %d, but got %d", fiber.StatusOK, resp.StatusCode)
			t.Fail()
		}
	})
}
