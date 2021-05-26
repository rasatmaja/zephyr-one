package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// InitializeRoute ...
func (a *App) InitializeRoute() {
	a.server.Get("/", func(c *fiber.Ctx) error {
		fLog := a.logger.With().Str("func", "InitializeRoute").Str("go", "router").Logger()
		fLog.Trace().Msg("Base URL Hit")
		return c.SendString("Hello from zephyr one")
	})

	a.server.Get("/hello", a.handler.HelloWorld)

	a.server.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(
			fiber.Map{
				"message": fmt.Sprintf("URL [%s] not found", c.Path()),
			},
		)
	})
}
