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
	"github.com/rasatmaja/zephyr-one/internal/database/models"
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

func TestContact(t *testing.T) {
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

	defer app.Shutdown()

	t.Run("error-empty-user-id", func(t *testing.T) {
		app.Get("/contact", handler.Contact)
		req := httptest.NewRequest("GET", "/contact", nil)
		resp, err := app.Test(req)

		// begin assert response from http test
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})

	// mock custom middleware to set context value
	app.Use(func(c *fiber.Ctx) error {
		c.Locals(constant.AuthIDContext, "12345")
		return c.Next()
	})

	t.Run("error-get-contact", func(t *testing.T) {
		// start repo mock
		repo := &repository.Mock{}
		var contact []*models.Contact
		repo.On("Contacts", mock.Anything, mock.Anything, mock.Anything).Return(contact, fmt.Errorf("Error DB"))
		handler.repo = repo

		app.Get("/contacts", handler.Contact)
		req := httptest.NewRequest("GET", "/contacts", nil)
		resp, err := app.Test(req)

		// begin assert response from http test
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		// start repo mock
		repo := &repository.Mock{}
		var contact []*models.Contact
		repo.On("Contacts", mock.Anything, mock.Anything, mock.Anything).Return(contact, nil)
		handler.repo = repo

		app.Get("/contacts", handler.Contact)
		req := httptest.NewRequest("GET", "/contacts", nil)
		resp, err := app.Test(req)

		// begin assert response from http test
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

}

func TestSetPrimaryContact(t *testing.T) {
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

	defer app.Shutdown()

	t.Run("error-empty-user-id", func(t *testing.T) {
		app.Put("/contact/primary/:contact", handler.SetPrimaryContact)
		req := httptest.NewRequest("PUT", "/contact/primary/test@test.test", nil)
		resp, err := app.Test(req)

		// begin assert response from http test
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})

	// mock custom middleware to set context value
	app.Use(func(c *fiber.Ctx) error {
		c.Locals(constant.AuthIDContext, "12345")
		return c.Next()
	})

	t.Run("error-empty-user-id", func(t *testing.T) {

		// start repo mock
		repo := &repository.Mock{}
		repo.On("SetPrimaryContact", mock.Anything, mock.Anything, mock.Anything).Return(fmt.Errorf("Error DB"))
		handler.repo = repo

		app.Put("/user/contact/primary/:contact", handler.SetPrimaryContact)
		req := httptest.NewRequest("PUT", "/user/contact/primary/test@test.test", nil)
		resp, err := app.Test(req)

		// begin assert response from http test
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("success", func(t *testing.T) {

		// start repo mock
		repo := &repository.Mock{}
		repo.On("SetPrimaryContact", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		handler.repo = repo

		app.Put("/user/contact/primary/:contact", handler.SetPrimaryContact)
		req := httptest.NewRequest("PUT", "/user/contact/primary/test@test.test", nil)
		resp, err := app.Test(req)

		// begin assert response from http test
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})
}

func TestRemovePrimaryContact(t *testing.T) {
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

	defer app.Shutdown()

	t.Run("error-empty-user-id", func(t *testing.T) {
		app.Delete("/contact/:contact", handler.RemoveContact)
		req := httptest.NewRequest("DELETE", "/contact/test@test.test", nil)
		resp, err := app.Test(req)

		// begin assert response from http test
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})

	// mock custom middleware to set context value
	app.Use(func(c *fiber.Ctx) error {
		c.Locals(constant.AuthIDContext, "12345")
		return c.Next()
	})

	t.Run("error-repo-delete-contact", func(t *testing.T) {

		// start repo mock
		repo := &repository.Mock{}
		repo.On("DeleteContact", mock.Anything, mock.Anything, mock.Anything).Return(fmt.Errorf("Error DB"))
		handler.repo = repo

		app.Delete("/user/contact/:contact", handler.RemoveContact)
		req := httptest.NewRequest("DELETE", "/user/contact/test@test.test", nil)
		resp, err := app.Test(req)

		// begin assert response from http test
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("success", func(t *testing.T) {

		// start repo mock
		repo := &repository.Mock{}
		repo.On("DeleteContact", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		handler.repo = repo

		app.Delete("/user/contact/:contact", handler.RemoveContact)
		req := httptest.NewRequest("DELETE", "/user/contact/test@test.test", nil)
		resp, err := app.Test(req)

		// begin assert response from http test
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})
}
