package server

import (
	"github.com/gofiber/fiber/v2"
)

// InitializeRoute ...
func (a *App) InitializeRoute() {
	defer a.server.Use(a.handler.PageNotfound)

	a.server.Get("/", func(c *fiber.Ctx) error {
		fLog := a.logger.With().Str("func", "InitializeRoute").Str("go", "router").Logger()
		fLog.Trace().Msg("Base URL Hit")
		return c.SendString("Hello from zephyr one")
	})

	a.server.Get("/hello", a.handler.HelloWorld)
}
