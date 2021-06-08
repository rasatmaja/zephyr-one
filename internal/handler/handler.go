package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/zephyr-one/internal/logger"
)

// Endpoint ...
type Endpoint struct {
	log *logger.Logger
}

// New is afunction to create handle instance
func New() *Endpoint {
	log := logger.New()
	return &Endpoint{
		log: log,
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
	return c.Status(fiber.StatusNotFound).JSON(
		fiber.Map{
			"message": fmt.Sprintf("URL [%s] not found", c.Path()),
		},
	)
}

// Base is a handler to handle base url "/"
func (e *Endpoint) Base(c *fiber.Ctx) error {
	fLog := e.log.With().Str("func", "Base").Logger()
	if c.Method() == fiber.MethodPost {
		fLog.Trace().Msg("Base URL POST Hit")
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello from Zephyr One",
		})
	}
	fLog.Trace().Msg("Base URL Hit")
	return c.SendString("Hello from zephyr one")
}
