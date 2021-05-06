package handler

import "github.com/gofiber/fiber/v2"

// Endpoint ...
type Endpoint struct{}

// New is afunction to create handle instance
func New() *Endpoint {
	return &Endpoint{}
}

// HelloWorld ...
func (e *Endpoint) HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}
