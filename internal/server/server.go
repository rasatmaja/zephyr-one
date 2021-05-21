package server

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/zephyr-one/internal/config"
	"github.com/rasatmaja/zephyr-one/internal/handler"
	"github.com/rasatmaja/zephyr-one/internal/logger"
)

// App is struct that define server, repo, and all component app needs
type App struct {
	server  *fiber.App
	handler *handler.Endpoint
	logger  *logger.Logger
	env     *config.ENV
}

// New is a function to initialize sever and its component
func New() *App {

	// setup logger
	log := logger.New()

	// setup config
	env := config.LoadENV()

	// setup handler
	handler := handler.New()

	return &App{
		server:  fiber.New(),
		handler: handler,
		logger:  log,
		env:     env,
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

		err := a.server.Shutdown()
		if err != nil {
			panic(err)
		}
	}()

	err := a.server.Listen(fmt.Sprintf("%s:%d", a.env.ServerHost, a.env.ServerPort))
	if err != nil {
		panic(err)
	}
	fmt.Println("Running cleanup tasks...")
	fmt.Println("Server Shutdown...")
}
