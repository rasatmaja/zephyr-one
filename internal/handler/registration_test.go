package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/zephyr-one/internal/config"
	"github.com/rasatmaja/zephyr-one/internal/database/models"
	"github.com/rasatmaja/zephyr-one/internal/database/repository"
	"github.com/rasatmaja/zephyr-one/internal/logger"
	"github.com/rasatmaja/zephyr-one/internal/password"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegistration(t *testing.T) {
	// setup fiber app
	app := fiber.New()
	defer app.Shutdown()
	env := config.LoadENV()
	env.LogLevel = "disable" // disable logging

	// setup handler
	handler := &Endpoint{
		log: logger.New(),
	}

	t.Run("error-parse-request-body", func(t *testing.T) {
		app.Post("/register", handler.Regitration)
		resp, err := app.Test(httptest.NewRequest("POST", "/register", nil))

		// begin assert response from http test
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("error-hash-password", func(t *testing.T) {

		// start password mock
		pwd := &password.Mock{}
		pwd.On("Hash", mock.Anything).Return("", fmt.Errorf("Error Hashing passprase")).Once()

		handler.password = pwd

		// build body request
		body := &RegistrationRes{
			Username:   "test-username",
			Passphrase: "test-passphrase",
			Name:       "test-name",
		}
		sbody, _ := json.Marshal(body)
		req := httptest.NewRequest("POST", "/register", bytes.NewReader(sbody))
		req.Header.Set("Content-Type", "application/json")

		app.Post("/register", handler.Regitration)
		resp, err := app.Test(req)

		// begin assert response from http test
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("error-create-auth", func(t *testing.T) {

		// start password mock
		pwd := &password.Mock{}
		pwd.On("Hash", mock.Anything).Return("", nil).Once()
		handler.password = pwd

		// start repo mock
		repo := &repository.Mock{}
		repo.On("CreateAuth", mock.Anything, mock.Anything, mock.Anything).Return(&models.Auth{}, fmt.Errorf("error create auth")).Once()
		handler.repo = repo

		// build body request
		body := &RegistrationRes{
			Username:   "test-username",
			Passphrase: "test-passphrase",
			Name:       "test-name",
		}
		sbody, _ := json.Marshal(body)
		req := httptest.NewRequest("POST", "/register", bytes.NewReader(sbody))
		req.Header.Set("Content-Type", "application/json")

		app.Post("/register", handler.Regitration)
		resp, err := app.Test(req)

		// begin assert response from http test
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("error-create-auth", func(t *testing.T) {

		// start password mock
		pwd := &password.Mock{}
		pwd.On("Hash", mock.Anything).Return("", nil).Once()
		handler.password = pwd

		// start repo mock
		repo := &repository.Mock{}
		repo.On("CreateAuth", mock.Anything, mock.Anything, mock.Anything).Return(&models.Auth{}, nil).Once()
		repo.On("CreateAccountInfo", mock.Anything, mock.Anything, mock.Anything).Return(&models.AccountInfo{}, fmt.Errorf("error create account info")).Once()
		handler.repo = repo

		// build body request
		body := &RegistrationRes{
			Username:   "test-username",
			Passphrase: "test-passphrase",
			Name:       "test-name",
		}
		sbody, _ := json.Marshal(body)
		req := httptest.NewRequest("POST", "/register", bytes.NewReader(sbody))
		req.Header.Set("Content-Type", "application/json")

		app.Post("/register", handler.Regitration)
		resp, err := app.Test(req)

		// begin assert response from http test
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("success", func(t *testing.T) {

		// start password mock
		pwd := &password.Mock{}
		pwd.On("Hash", mock.Anything).Return("", nil).Once()
		handler.password = pwd

		// start repo mock
		repo := &repository.Mock{}
		repo.On("CreateAuth", mock.Anything, mock.Anything, mock.Anything).Return(&models.Auth{}, nil).Once()
		repo.On("CreateAccountInfo", mock.Anything, mock.Anything, mock.Anything).Return(&models.AccountInfo{}, nil).Once()
		handler.repo = repo

		// build body request
		body := &RegistrationRes{
			Username:   "test-username",
			Passphrase: "test-passphrase",
			Name:       "test-name",
		}
		sbody, _ := json.Marshal(body)
		req := httptest.NewRequest("POST", "/register", bytes.NewReader(sbody))
		req.Header.Set("Content-Type", "application/json")

		app.Post("/register", handler.Regitration)
		resp, err := app.Test(req)

		// begin assert response from http test
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	})
}