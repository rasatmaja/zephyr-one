package middleware

import (
	"fmt"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rasatmaja/zephyr-one/internal/config"
	"github.com/rasatmaja/zephyr-one/internal/constant"
	"github.com/rasatmaja/zephyr-one/internal/logger"
	"github.com/rasatmaja/zephyr-one/internal/token"
	"github.com/rasatmaja/zephyr-one/internal/token/contract"
	"github.com/rasatmaja/zephyr-one/pkg/helper"
)

var instance *App
var singleton sync.Once

// App is a struct to define a middleware fiber app
type App struct {
	server *fiber.App
	log    *logger.Logger
	env    *config.ENV
	token  contract.IToken
}

// New is a function to initialize middleware
func New(app *fiber.App) *App {
	singleton.Do(func() {
		fmt.Println("[ MDWR ] Starting middleware core ...")
		instance = &App{
			server: app,
			log:    logger.New(),
			env:    config.LoadENV(),
			token:  token.Factory(),
		}
	})
	return instance
}

// Initialize is a function to register and initialize middleware func
func (mdlwr *App) Initialize() {
	mdlwr.ResponseTime()
	mdlwr.RequestID()
	mdlwr.TransactionID()
	mdlwr.Recover()
	mdlwr.Profiling()
	mdlwr.SwaggerUI()
}

// RequestID is a function to initialize request id for http header as a midleware
func (mdlwr *App) RequestID() {
	fmt.Println("[ MDWR ] Initialize RequestID middleware")
	mdlwr.server.Use(func(c *fiber.Ctx) error {
		reqID := c.Get(constant.XRequestIDHTTPHeader)
		if len(reqID) == 0 {
			str := helper.NewStrings()
			reqID, _ = str.GenerateRandomString(8)
			c.Set(constant.XRequestIDHTTPHeader, reqID)
		}
		return c.Next()
	})
}

// Recover is a function to initialize recover as a midleware
func (mdlwr *App) Recover() {
	fmt.Println("[ MDWR ] Initialize Recover middleware")
	mdlwr.server.Use(recover.New())
}

// TransactionID is a function to initialize trasaction id for http header as a midleware
// this header only appear on method POST, DELETE, and PUT (except GET)
func (mdlwr *App) TransactionID() {
	fmt.Println("[ MDWR ] Initialize Transaction ID middleware")
	mdlwr.server.Use(func(c *fiber.Ctx) error {
		if c.Method() != "GET" {
			trxID := c.Get(constant.XtransactionIDHTTPHeader)
			if len(trxID) == 0 {
				str := helper.NewStrings()
				trxID, _ = str.GenerateRandomString(32)
				c.Set(constant.XtransactionIDHTTPHeader, trxID)
			}
		}
		return c.Next()
	})
}

// ResponseTime is a middleware to track how much time it takes to process a request
// This middlware only active if server not in production
func (mdlwr *App) ResponseTime() {
	if !mdlwr.env.ServerProduction {
		fmt.Println("[ MDWR ] Initialize Response Time middleware")
		mdlwr.server.Use(func(c *fiber.Ctx) error {
			start := time.Now()
			err := c.Next()
			dur := time.Since(start).Milliseconds()
			mdlwr.log.Trace().Msgf("Request [%s] done in %d ms", c.Path(), dur)
			return err
		})
	}
}

// Profiling is a midlleware to attach pprof into server
func (mdlwr *App) Profiling() {
	fmt.Println("[ MDWR ] Initialize Profilling middleware")
	if !mdlwr.env.ServerProduction {
		mdlwr.server.Use(pprof.New())

		// redirect into pprof
		mdlwr.server.Get("/debug", func(c *fiber.Ctx) error {
			slashPresent := c.Path()[len(c.Path())-1:] == "/"
			if slashPresent {
				return c.Redirect("pprof")
			}
			return c.Redirect("debug/pprof")
		})
	}
}
