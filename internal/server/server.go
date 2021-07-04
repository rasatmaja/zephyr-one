package server

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/zephyr-one/internal/config"
	"github.com/rasatmaja/zephyr-one/internal/handler"
	"github.com/rasatmaja/zephyr-one/internal/logger"
	"github.com/rasatmaja/zephyr-one/internal/middleware"
	"github.com/rasatmaja/zephyr-one/internal/utils"
)

// App is struct that define server, repo, and all component app needs
type App struct {
	server     *fiber.App
	handler    *handler.Endpoint
	logger     *logger.Logger
	env        *config.ENV
	middleware *middleware.App
	utils      *utils.Registry
}

// New is a function to initialize sever and its component
func New() *App {

	// setup logger
	log := logger.New()

	// setup config
	env := config.LoadENV()

	// setup handler
	handler := handler.New()

	// setup utils
	utils := utils.New()

	// setup server
	svr := fiber.New(
		fiber.Config{
			ReadTimeout:  time.Duration(env.ServerReadTO) * time.Second,
			WriteTimeout: time.Duration(env.ServerWriteTO) * time.Second,
			IdleTimeout:  time.Duration(env.ServerIdleTO) * time.Second,

			// Implement default error
			ErrorHandler: handler.Error,
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
		utils:      utils,
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
		cert, _ := a.utils.Cert.GenerateSelfSignedCertificates()

		// Register certificate to asset registry for cleanup
		a.utils.Assets.Register(utils.Asset{Path: cert.CertPath, Type: utils.AssetFile})
		a.utils.Assets.Register(utils.Asset{Path: cert.KeyPath, Type: utils.AssetFile})

		fmt.Println("[ SRVR ] Server using self-signed certificate")
		err = a.server.ListenTLS(host, cert.CertPath, cert.KeyPath)
	} else {
		err = a.server.Listen(host)
	}

	if err != nil {
		panic(err)
	}
}

// InitializeMiddleware is a function to start middleware
func (a *App) InitializeMiddleware() {
	a.middleware.Initialize()
}
