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
func New(log *logger.Logger) *Endpoint {
	return &Endpoint{
		log: log,
	}
}

// HelloWorld ...
func (e *Endpoint) HelloWorld(c *fiber.Ctx) error {
	e.log.Log.Info("Endpoint Hello World Hit")
	return c.SendString("Hello World")
}
