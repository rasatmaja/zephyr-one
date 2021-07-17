package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/zephyr-one/internal/response"
)

// AddContactReq define request toadd contact
type AddContactReq struct {
	Contact string `json:"contact"`
	Type    string `json:"type"`
}

// AddContact is a handler to add contact
func (e *Endpoint) AddContact(c *fiber.Ctx) error {
	fLog := e.log.With().Str("func", "AddContact").Logger()

	// build response
	res := response.Factory()

	// parse body request from JSON to struct
	req := &AddContactReq{}
	err := c.BodyParser(req)
	if err != nil {
		fLog.Error().Msgf("BodyParser error, got: %v", err)
		return res.BadRequest("unable processing request")
	}

	return res.Success("User add contact").SetData(req)
}
