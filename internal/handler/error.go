package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/zephyr-one/internal/response"
)

func (e *Endpoint) Error(c *fiber.Ctx, err error) error {
	if e, ok := err.(*response.Response); ok {
		return c.Status(e.Code).JSON(e)
	}
	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Error{
		Code:    fiber.StatusInternalServerError,
		Message: err.Error(),
	})
}
