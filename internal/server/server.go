package server

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/zephyr-one/internal/config"
	"github.com/rasatmaja/zephyr-one/internal/handler"
	"github.com/rasatmaja/zephyr-one/internal/logger"
	"github.com/rasatmaja/zephyr-one/internal/middleware"
)

// App is struct that define server, repo, and all component app needs
type App struct {
	server     *fiber.App
	handler    *handler.Endpoint
	logger     *logger.Logger
	env        *config.ENV
	middleware *middleware.App
}

// New is a function to initialize sever and its component
func New() *App {

	// setup logger
	log := logger.New()

	// setup config
	env := config.LoadENV()

	// setup handler
	handler := handler.New()

	// setup server
	svr := fiber.New(
		fiber.Config{
			ReadTimeout:  time.Duration(env.ServerReadTO) * time.Second,
			WriteTimeout: time.Duration(env.ServerWriteTO) * time.Second,
			IdleTimeout:  time.Duration(env.ServerIdleTO) * time.Second,
		},
	)

	// setup middleware
	mdlwre := middleware.New(svr)

	return &App{
		server:     svr,
		handler:    handler,
		logger:     log,
		env:        env,
		middleware: mdlwre,
	}
}

// Start is a function to start server
func (a *App) Start() {

	a.InitializeMiddleware()
	a.InitializeRoute()

	a.InitializeShutdownSequence()
	defer fmt.Println("[ SRVR ] Server Shutdown...")

	a.ServerListen()

}

// ServerListen is a function to initialize server listen
func (a *App) ServerListen() {

	var err error
	host := fmt.Sprintf("%s:%d", a.env.ServerHost, a.env.ServerPort)

	if a.env.TLS {
		cert := a.GenerateSelfSignedCertificates()
		fmt.Println("[ SRVR ] Server using self-signed certificate")
		err = a.server.ListenTLS(host, cert.CertPath, cert.KeyPath)
	} else {
		err = a.server.Listen(host)
	}

	if err != nil {
		panic(err)
	}
}

// InitializeShutdownSequence is a function initialize shutdown sequence
func (a *App) InitializeShutdownSequence() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println(" :: OS signal received")
		fmt.Println("[ SRVR ] Gracefully shutting down...")
		fmt.Println("[ SRVR ] Running cleanup tasks...")

		os.Remove("cert.pem")
		os.Remove("key.pem")

		err := a.server.Shutdown()
		if err != nil {
			panic(err)
		}
	}()

}

// InitializeMiddleware is a function to start middleware
func (a *App) InitializeMiddleware() {
	a.middleware.Initialize()
}
