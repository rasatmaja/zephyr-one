package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/zephyr-one/internal/response"
	"github.com/rasatmaja/zephyr-one/internal/token/contract"
)

// AuthReq define a login request
type AuthReq struct {
	Username   string `json:"username"`
	Passphrase string `json:"passphrase"`
}

// AuthRes define login response
type AuthRes struct {
	Token string `json:"token"`
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

	// get auth data
	auth, err := e.repo.Auth(c.Context(), req.Username)
	if err != nil {
		fLog.Error().Msgf("Auth error, got: %v", err)
		return res.InternalServerError("unable get auth data from database")
	}

	// compare passphrase
	match, err := e.password.Compare(req.Passphrase, auth.Passphrase)
	if err != nil {
		fLog.Error().Msgf("Auth error, got: %v", err)
		return res.InternalServerError("unable get compare passphrase with data from database")
	}

	if !match {
		fLog.Error().Msg("Passphrase not match")
		return res.Unauthorized("Passphrase not match")
	}

	// build jwt payloads
	payload := &contract.Payload{
		Subject: auth.Username,
		JWTID:   auth.ID,
	}
	token, err := e.token.Sign(c.Context(), payload)
	if err != nil {
		fLog.Error().Msgf("token build error, got: %v", err)
		return res.InternalServerError("failed to build JWT token")
	}

	return res.Success("User successfully authenticate").SetData(&AuthRes{Token: token})
}
