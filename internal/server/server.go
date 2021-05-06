package server

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/zephyr-one/internal/handler"
)

// App is struct that define server, repo, and all component app needs
type App struct {
	server  *fiber.App
	handler *handler.Endpoint
}

// New is a function to initialize sever and its component
func New() *App {
	return &App{
		server:  fiber.New(),
		handler: handler.New(),
	}
}

// Start is a function to start server
func (a *App) Start() {
	a.server.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from zephyr one")
	})

	a.server.Get("/hello", a.handler.HelloWorld)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println(" <- OS signal received")
		fmt.Println("Gracefully shutting down...")
		a.server.Shutdown()
	}()

	a.server.Listen(":3090")
	fmt.Println("Running cleanup tasks...")
	fmt.Println("Server Shutdown...")
}
