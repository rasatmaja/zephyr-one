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
	"github.com/rasatmaja/zephyr-one/internal/token/contract"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuth(t *testing.T) {
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
	app.Post("/login", handler.Auth)
	defer app.Shutdown()

	t.Run("error-parse-request-body", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/login", nil)
		resp, err := app.Test(req)

		// begin assert response from http test
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("error-get-auth", func(t *testing.T) {
		// start repo mock
		repo := &repository.Mock{}
		repo.On("Auth", mock.Anything, mock.Anything).Return(&models.Auth{}, fmt.Errorf("error get auth"))
		handler.repo = repo

		// build body request
		body := &AuthReq{
			Username:   "test",
			Passphrase: "test",
		}

		sbody, _ := json.Marshal(body)
		req := httptest.NewRequest("POST", "/login", bytes.NewReader(sbody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		// begin assert response from http test
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("error-compare-password", func(t *testing.T) {
		// start repo mock
		repo := &repository.Mock{}
		repo.On("Auth", mock.Anything, mock.Anything).Return(&models.Auth{}, nil)
		handler.repo = repo

		// start password mock
		pwd := &password.Mock{}
		pwd.On("Compare", mock.Anything, mock.Anything).Return(false, fmt.Errorf("failed to compare")).Once()
		handler.password = pwd

		// build body request
		body := &AuthReq{
			Username:   "test",
			Passphrase: "test",
		}

		sbody, _ := json.Marshal(body)
		req := httptest.NewRequest("POST", "/login", bytes.NewReader(sbody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		// begin assert response from http test
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("error-phassprase-not-match", func(t *testing.T) {
		// start repo mock
		repo := &repository.Mock{}
		repo.On("Auth", mock.Anything, mock.Anything).Return(&models.Auth{}, nil)
		handler.repo = repo

		// start password mock
		pwd := &password.Mock{}
		pwd.On("Compare", mock.Anything, mock.Anything).Return(false, nil).Once()
		handler.password = pwd

		// build body request
		body := &AuthReq{
			Username:   "test",
			Passphrase: "test",
		}

		sbody, _ := json.Marshal(body)
		req := httptest.NewRequest("POST", "/login", bytes.NewReader(sbody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		// begin assert response from http test
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("error-create-token", func(t *testing.T) {
		// start repo mock
		repo := &repository.Mock{}
		repo.On("Auth", mock.Anything, mock.Anything).Return(&models.Auth{}, nil)
		handler.repo = repo

		// start password mock
		pwd := &password.Mock{}
		pwd.On("Compare", mock.Anything, mock.Anything).Return(true, nil).Once()
		handler.password = pwd

		// start token mock
		tkn := &contract.Mock{}
		tkn.On("Sign", mock.Anything, mock.Anything).Return("", fmt.Errorf("error build token"))
		handler.token = tkn

		// build body request
		body := &AuthReq{
			Username:   "test",
			Passphrase: "test",
		}

		sbody, _ := json.Marshal(body)
		req := httptest.NewRequest("POST", "/login", bytes.NewReader(sbody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		// begin assert response from http test
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("error-create-token", func(t *testing.T) {
		// start repo mock
		repo := &repository.Mock{}
		repo.On("Auth", mock.Anything, mock.Anything).Return(&models.Auth{}, nil)
		handler.repo = repo

		// start password mock
		pwd := &password.Mock{}
		pwd.On("Compare", mock.Anything, mock.Anything).Return(true, nil).Once()
		handler.password = pwd

		// start token mock
		tkn := &contract.Mock{}
		tkn.On("Sign", mock.Anything, mock.Anything).Return("", nil)
		handler.token = tkn

		// build body request
		body := &AuthReq{
			Username:   "test",
			Passphrase: "test",
		}

		sbody, _ := json.Marshal(body)
		req := httptest.NewRequest("POST", "/login", bytes.NewReader(sbody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		// begin assert response from http test
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})
}
