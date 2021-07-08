package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/zephyr-one/internal/response"
)

// AuthReq define a login request
type AuthReq struct {
	Username   string `json:"username"`
	Passphrase string `json:"passphrase"`
}

// Auth is a handler to authentications process
func (e *Endpoint) Auth(c *fiber.Ctx) error {
	fLog := e.log.With().Str("func", "Auth").Logger()

	// build response
	res := response.Factory()

	// parse body request from JSON to struct
	req := &AuthReq{}
	err := c.BodyParser(req)
	if err != nil {
		fLog.Error().Msgf("BodyParser error, got: %v", err)
		return res.BadRequest("unable processing request")
	}

	return res.Success()
}
