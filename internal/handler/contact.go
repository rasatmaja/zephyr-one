package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/zephyr-one/internal/constant"
	"github.com/rasatmaja/zephyr-one/internal/database/models"
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

	var authID string
	if c.Locals(constant.AuthIDContext) != nil {
		authID = c.Locals(constant.AuthIDContext).(string)
	}
	contact := &models.Contact{
		AuthID:        authID,
		ContactTypeID: "1", // TODO: should replace baseed on request body
		Contact:       req.Contact,
	}

	err = e.repo.CreateContact(c.Context(), contact)
	if err != nil {
		fLog.Error().Msgf("unable adding contact [%s] on [%s] , got: %v", req.Contact, authID, err)
		return res.InternalServerError("unable adding contact")
	}

	return res.Success("successfully add contact")
}