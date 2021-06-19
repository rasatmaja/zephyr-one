package middleware

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/zephyr-one/internal/config"
)

func TestMiddleware(t *testing.T) {
	// setup fiber app
	app := fiber.New()

	// load env
	env := config.LoadENV()
	env.LogLevel = "disable"

	// setup middleware
	mdlwr := New(app)
	mdlwr.Initialize()

	// run test case
	t.Run("request-id/success", func(t *testing.T) {

		// begin test
		app.Get("/test", func(c *fiber.Ctx) error {
			return c.SendStatus(200)
		})
		resp, err := app.Test(httptest.NewRequest("GET", "/test", nil))

		// begin assert response from http test
		if instance == nil {
			t.Error("instance shouldn't be nil")
			t.Fail()
		}

		if err != nil {
			t.Error("error should be nil")
			t.Fail()
		}

		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("status code should be %d", fiber.StatusOK)
			t.Fail()
		}

		if len(resp.Header["X-Request-Id"]) == 0 {
			t.Error("header X-Request-Id should be present")
			t.Fail()
		}
	})

	t.Run("transaction-id/success", func(t *testing.T) {

		// begin test
		app.Post("/test-post", func(c *fiber.Ctx) error {
			return c.SendStatus(200)
		})
		resp, err := app.Test(httptest.NewRequest("POST", "/test-post", nil))

		// begin assert response from http test
		if instance == nil {
			t.Error("instance shouldn't be nil")
			t.Fail()
		}

		if err != nil {
			t.Error("error should be nil")
			t.Fail()
		}

		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("status code should be %d", fiber.StatusOK)
			t.Fail()
		}

		if len(resp.Header["X-Transaction-Id"]) == 0 {
			t.Error("header X-Transaction-Id should be present")
			t.Fail()
		}
	})

	t.Run("debug-pprof/success", func(t *testing.T) {
		resp, err := app.Test(httptest.NewRequest("GET", "/debug", nil))
		resp, err = app.Test(httptest.NewRequest("GET", "/debug/", nil))

		// begin assert response from http test
		if instance == nil {
			t.Error("instance shouldn't be nil")
			t.Fail()
		}

		if err != nil {
			t.Error("error should be nil")
			t.Fail()
		}

		if resp.StatusCode != fiber.StatusFound {
			t.Errorf("status code should be %d, but got %d", fiber.StatusFound, resp.StatusCode)
			t.Fail()
		}
	})

	t.Run("swagger/success", func(t *testing.T) {
		resp, err := app.Test(httptest.NewRequest("GET", "/docs", nil))
		resp, err = app.Test(httptest.NewRequest("GET", "/docs/", nil))

		// begin assert response from http test
		if instance == nil {
			t.Error("instance shouldn't be nil")
			t.Fail()
		}

		if err != nil {
			t.Error("error should be nil")
			t.Fail()
		}

		if resp.StatusCode != fiber.StatusFound {
			t.Errorf("status code should be %d, but got %d", fiber.StatusFound, resp.StatusCode)
			t.Fail()
		}
	})

}
