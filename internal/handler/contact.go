package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/zephyr-one/internal/constant"
	"github.com/rasatmaja/zephyr-one/internal/database/models"
	zosql "github.com/rasatmaja/zephyr-one/internal/database/sql"
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
		if err == zosql.ErrDataDuplicate {
			return res.BadRequest("contact already exist")
		}
		return res.InternalServerError("unable adding contact")
	}

	return res.Success("successfully add contact")
}

// Contact is a handler to get all user contacts
func (e *Endpoint) Contact(c *fiber.Ctx) error {
	fLog := e.log.With().Str("func", "AllContacts").Logger()

	// build response
	res := response.Factory()

	var authID string
	if c.Locals(constant.AuthIDContext) == nil {
		return res.Unauthorized("user id empty")
	}
	authID = c.Locals(constant.AuthIDContext).(string)

	//get type params
	types := c.Params("type")

	contacs, err := e.repo.Contacts(c.Context(), authID, types)
	if err != nil {
		fLog.Error().Msgf("unable retrive contact [%s], got %s", authID, err)
		return res.InternalServerError("unable get contact")
	}

	return res.Success().SetData(contacs)
}

// SetPrimaryContact is a handler to update user primary contact
func (e *Endpoint) SetPrimaryContact(c *fiber.Ctx) error {
	fLog := e.log.With().Str("func", "SetPrimaryContact").Logger()

	// build response
	res := response.Factory()

	var authID string
	if c.Locals(constant.AuthIDContext) == nil {
		return res.Unauthorized("user id empty")
	}
	authID = c.Locals(constant.AuthIDContext).(string)

	//get contact params
	contact := c.Params("contact")

	err := e.repo.SetPrimaryContact(c.Context(), authID, contact)
	if err != nil {
		fLog.Error().Msgf("unable update contact [%s], got %s", authID, err)
		return res.InternalServerError("unable get contact")
	}

	return res.Success()
}
