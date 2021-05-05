package server

import "github.com/gofiber/fiber/v2"

// App is struct that define server, repo, and all component app needs
type App struct {
	server *fiber.App
}

// New is a function to initialize sever and its component
func New() *App {
	server := fiber.New()
	return &App{
		server: server,
	}
}

// Start is a function to start server
func (a *App) Start() {
	a.server.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from zephyr one")
	})
	a.server.Listen(":3090")
}
