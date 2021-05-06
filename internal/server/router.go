package server

import "github.com/gofiber/fiber/v2"

// InitializeRoute ...
func (a *App) InitializeRoute() {
	a.server.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from zephyr one")
	})

	a.server.Get("/hello", a.handler.HelloWorld)
}
