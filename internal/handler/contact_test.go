package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/zephyr-one/internal/config"
	"github.com/rasatmaja/zephyr-one/internal/constant"
	"github.com/rasatmaja/zephyr-one/internal/database/repository"
	zosql "github.com/rasatmaja/zephyr-one/internal/database/sql"
	"github.com/rasatmaja/zephyr-one/internal/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddContact(t *testing.T) {
	env := config.LoadENV()
	env.LogLevel = "disable" // disable logging

	// setup handler
	handler := &Endpoint{
		log: logger.New(),
	}

	// setup fiber app
	app := fiber.New(
		fiber.Config{
			ErrorHandler: handler.Error,
		},
	)
	// mock custom middleware to set context value
	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("middleware")
		c.Locals(constant.AuthIDContext, "12345")
		return c.Next()
	})
	app.Post("/contact", handler.AddContact)
	defer app.Shutdown()

	t.Run("error-parse-request-body", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/contact", nil)
		resp, err := app.Test(req)

		// begin assert response from http test
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("error-create-contact", func(t *testing.T) {

		// start repo mock
		repo := &repository.Mock{}
		repo.On("CreateContact", mock.Anything, mock.Anything).Return(fmt.Errorf("error get auth"))
		handler.repo = repo

		// build body request
		body := &AddContactReq{
			Contact: "test",
			Type:    "test",
		}
		sbody, _ := json.Marshal(body)

		req := httptest.NewRequest("POST", "/contact", bytes.NewReader(sbody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		// begin assert response from http test
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("error-create-contact", func(t *testing.T) {

		// start repo mock
		repo := &repository.Mock{}
		repo.On("CreateContact", mock.Anything, mock.Anything).Return(zosql.ErrDataDuplicate)
		handler.repo = repo

		// build body request
		body := &AddContactReq{
			Contact: "test",
			Type:    "test",
		}
		sbody, _ := json.Marshal(body)

		req := httptest.NewRequest("POST", "/contact", bytes.NewReader(sbody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		// begin assert response from http test
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("success", func(t *testing.T) {

		// start repo mock
		repo := &repository.Mock{}
		repo.On("CreateContact", mock.Anything, mock.Anything).Return(nil)
		handler.repo = repo

		// build body request
		body := &AddContactReq{
			Contact: "test",
			Type:    "test",
		}
		sbody, _ := json.Marshal(body)

		req := httptest.NewRequest("POST", "/contact", bytes.NewReader(sbody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		// begin assert response from http test
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})
}
