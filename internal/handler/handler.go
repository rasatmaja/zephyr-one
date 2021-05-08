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
	return &Endpoint{}
}

// SetLogger is a function to set logger on handler
func (e *Endpoint) SetLogger(log *logger.Logger) { e.log = log }

// HelloWorld ...
func (e *Endpoint) HelloWorld(c *fiber.Ctx) error {
	e.log.Info("Endpoint Hello World Hit")
	return c.SendString("Hello World")
}
