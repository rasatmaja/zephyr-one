package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/zephyr-one/internal/response"
)

// Regitration is a handler to registration process
func (e *Endpoint) Regitration(c *fiber.Ctx) error {
	res := response.Factory()
	res.Created("successfully registered")
	return c.Status(res.Code).JSON(res)
}
