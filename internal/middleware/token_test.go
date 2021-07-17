package middleware

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/zephyr-one/internal/config"
	"github.com/rasatmaja/zephyr-one/internal/handler"
	"github.com/rasatmaja/zephyr-one/internal/token/contract"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestVerifyToken(t *testing.T) {
	// setup handler
	handler := &handler.Endpoint{}

	// setup fiber app
	app := fiber.New(
		fiber.Config{
			ErrorHandler: handler.Error,
		},
	)

	// load env
	env := config.LoadENV()
	env.LogLevel = "disable"

	// setup middleware
	mdlwr := New(app)
	mdlwr.Initialize()

	app.Use(mdlwr.VerifyToken)
	app.Post("/test", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	// run test case
	t.Run("no-auth-header", func(t *testing.T) {

		// begin test
		req := httptest.NewRequest("POST", "/test", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		// begin assert response from http test
		assert.NotNil(t, instance)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("no-bearer-in-auth-header", func(t *testing.T) {

		// begin test
		req := httptest.NewRequest("POST", "/test", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ6ZXBoZXJ5LW9uZSIsInN1YiI6InJhc2lvMjkiLCJleHAiOjE2MjkxMTMwNzYsIm5iZiI6MTYyNjUyMTA3NiwiaWF0IjoxNjI2NTIxMDc2LCJqdGkiOiI3ZDk5MWY1My1jZjJkLTRmZWQtYmQwZC0zMDQ0M2JiNmJiYzEifQ.i8Uj4IimiwoHh0Abq2pg56NKSwqDZKKyJLgztN7hRXs")
		resp, err := app.Test(req)

		// begin assert response from http test
		assert.NotNil(t, instance)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("error-verify-token", func(t *testing.T) {

		// start token mock
		tkn := &contract.Mock{}
		tkn.On("Verify", mock.Anything, mock.Anything).Return(&contract.Payload{}, fmt.Errorf("error verify token"))
		mdlwr.token = tkn

		// begin test
		req := httptest.NewRequest("POST", "/test", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ6ZXBoZXJ5LW9uZSIsInN1YiI6InJhc2lvMjkiLCJleHAiOjE2MjkxMTMwNzYsIm5iZiI6MTYyNjUyMTA3NiwiaWF0IjoxNjI2NTIxMDc2LCJqdGkiOiI3ZDk5MWY1My1jZjJkLTRmZWQtYmQwZC0zMDQ0M2JiNmJiYzEifQ.i8Uj4IimiwoHh0Abq2pg56NKSwqDZKKyJLgztN7hRXs")
		resp, err := app.Test(req)

		// begin assert response from http test
		assert.NotNil(t, instance)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("success", func(t *testing.T) {

		// start token mock
		tkn := &contract.Mock{}
		tkn.On("Verify", mock.Anything, mock.Anything).Return(&contract.Payload{}, nil)
		mdlwr.token = tkn

		// begin test
		req := httptest.NewRequest("POST", "/test", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ6ZXBoZXJ5LW9uZSIsInN1YiI6InJhc2lvMjkiLCJleHAiOjE2MjkxMTMwNzYsIm5iZiI6MTYyNjUyMTA3NiwiaWF0IjoxNjI2NTIxMDc2LCJqdGkiOiI3ZDk5MWY1My1jZjJkLTRmZWQtYmQwZC0zMDQ0M2JiNmJiYzEifQ.i8Uj4IimiwoHh0Abq2pg56NKSwqDZKKyJLgztN7hRXs")
		resp, err := app.Test(req)

		// begin assert response from http test
		assert.NotNil(t, instance)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})
}
