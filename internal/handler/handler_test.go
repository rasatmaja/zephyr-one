package handler

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestHandler(t *testing.T) {
	// setup fiber app
	app := fiber.New()

	// setup handler
	handler := New()

	// run test case
	t.Run("Base/success", func(t *testing.T) {
		app.Get("/base", handler.Base)
		app.Post("/base", handler.Base)
		resp, err := app.Test(httptest.NewRequest("GET", "/base", nil))
		resp, err = app.Test(httptest.NewRequest("POST", "/base", nil))

		// begin assert response from http test
		if err != nil {
			t.Error("error should be nil")
			t.Fail()
		}

		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("status code should be %d, but got %d", fiber.StatusOK, resp.StatusCode)
			t.Fail()
		}
	})

	t.Run("hello/success", func(t *testing.T) {
		app.Get("/hello", handler.HelloWorld)
		resp, err := app.Test(httptest.NewRequest("GET", "/hello", nil))

		// begin assert response from http test
		if err != nil {
			t.Error("error should be nil")
			t.Fail()
		}

		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("status code should be %d, but got %d", fiber.StatusOK, resp.StatusCode)
			t.Fail()
		}
	})

	t.Run("not-found/success", func(t *testing.T) {
		app.Use(handler.PageNotfound)
		resp, err := app.Test(httptest.NewRequest("GET", "/unknown", nil))

		// begin assert response from http test
		if err != nil {
			t.Error("error should be nil")
			t.Fail()
		}

		if resp.StatusCode != fiber.StatusNotFound {
			t.Errorf("status code should be %d, but got %d", fiber.StatusNotFound, resp.StatusCode)
			t.Fail()
		}
	})

}
