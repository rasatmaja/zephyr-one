package server

import "github.com/gofiber/fiber/v2"

// Start is a function to start server
func Start() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from zephyr one")
	})
	app.Listen(":3090")
}
