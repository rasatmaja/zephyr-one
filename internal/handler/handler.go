package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/zephyr-one/internal/logger"
)

// Endpoint ...
type Endpoint struct {
	log *logger.Logger
}

// New is afunction to create handle instance
func New() *Endpoint {
	logger := logger.New()
	return &Endpoint{
		log: logger,
	}
}

// HelloWorld ...
func (e *Endpoint) HelloWorld(c *fiber.Ctx) error {
	fLog := e.log.With().Str("go", "handle").Str("func", "HelloWorld").Logger()
	fLog.Trace().Msg("Endpoint Hello World Hit")
	return c.SendString("Hello World")
}
