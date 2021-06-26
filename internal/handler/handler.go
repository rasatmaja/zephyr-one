package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/zephyr-one/internal/database"
	"github.com/rasatmaja/zephyr-one/internal/database/repository"
	"github.com/rasatmaja/zephyr-one/internal/logger"
	"github.com/rasatmaja/zephyr-one/internal/password"
	"github.com/rasatmaja/zephyr-one/internal/response"
)

// Endpoint ...
type Endpoint struct {
	log      *logger.Logger
	repo     repository.IRepository
	password password.IPassword
}

// New is afunction to create handle instance
func New() *Endpoint {
	return &Endpoint{
		log:      logger.New(),
		repo:     database.Factory(),
		password: password.Factory(),
	}
}

// HelloWorld ...
func (e *Endpoint) HelloWorld(c *fiber.Ctx) error {
	fLog := e.log.With().Str("func", "HelloWorld").Logger()
	fLog.Trace().Msg("Endpoint Hello World Hit")
	return c.SendString("Hello World")
}

// PageNotfound is a handler to handle undefined route
func (e *Endpoint) PageNotfound(c *fiber.Ctx) error {
	res := response.Factory()
	res.NotFound(fmt.Sprintf("URL [%s] not found", c.Path()))
	return c.Status(res.Code).JSON(res)
}

// Base is a handler to handle base url "/"
func (e *Endpoint) Base(c *fiber.Ctx) error {
	fLog := e.log.With().Str("func", "Base").Logger()
	if c.Method() == fiber.MethodPost {
		fLog.Trace().Msg("Base URL POST Hit")
		res := response.Factory()
		res.Success("Hello from Zephyr One")
		return c.Status(res.Code).JSON(res)
	}
	fLog.Trace().Msg("Base URL Hit")
	return c.SendString("Hello from zephyr one")
}
