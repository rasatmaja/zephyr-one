package server

import "github.com/gofiber/fiber/v2"

// Start is a function to start server
func Start() {
	app := fiber.New()
	app.Listen(":3090")
}
