package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/zephyr-one/internal/constant"
	"github.com/rasatmaja/zephyr-one/internal/response"
)

// Info is a handler to get user general information
func (e *Endpoint) Info(c *fiber.Ctx) error {
	fLog := e.log.With().Str("func", "AllContacts").Logger()

	// build response
	res := response.Factory()

	var authID string
	if c.Locals(constant.AuthIDContext) == nil {
		return res.Unauthorized("user id empty")
	}
	authID = c.Locals(constant.AuthIDContext).(string)

	user, err := e.repo.Account(c.Context(), authID)
	if err != nil {
		fLog.Error().Msgf("unable retrive account info [%s], got %s", authID, err)
		return res.InternalServerError("unable get account info")
	}

	return res.Success().SetData(user)
}
