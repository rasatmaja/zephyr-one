package server

import (
	"github.com/gofiber/fiber/v2"
)

// InitializeRoute ...
func (a *App) InitializeRoute() {
	fLog := a.logger.With().Str("func", "InitializeRoute").Str("go", "router").Logger()
	defer a.server.Use(a.handler.PageNotfound)

	a.server.Get("/", func(c *fiber.Ctx) error {
		fLog.Trace().Msg("Base URL Hit")
		return c.SendString("Hello from zephyr one")
	})

	a.server.Post("/", func(c *fiber.Ctx) error {
		fLog.Trace().Msg("Base URL POST Hit")
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello from Zephyr One",
		})
	})

	a.server.Get("/hello", a.handler.HelloWorld)
}
