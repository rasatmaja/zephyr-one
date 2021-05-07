package server

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/zephyr-one/internal/handler"
	"github.com/rasatmaja/zephyr-one/internal/logger"
)

// App is struct that define server, repo, and all component app needs
type App struct {
	server  *fiber.App
	handler *handler.Endpoint
	logger  *logger.Logger
}

// New is a function to initialize sever and its component
func New() *App {

	log := logger.New()
	return &App{
		server:  fiber.New(),
		handler: handler.New(log),
		logger:  log,
	}
}

// Start is a function to start server
func (a *App) Start() {

	a.InitializeRoute()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println(" <- OS signal received")
		fmt.Println("Gracefully shutting down...")

		defer func() {
			a.logger.Sync()
		}()

		err := a.server.Shutdown()
		if err != nil {
			panic(err)
		}
	}()

	err := a.server.Listen(":3090")
	if err != nil {
		panic(err)
	}
	fmt.Println("Running cleanup tasks...")
	fmt.Println("Server Shutdown...")
}
