package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/zephyr-one/internal/response"
)

// RegistrationRes define registration request body
type RegistrationRes struct {
	Username   string `json:"username"`
	Passphrase string `json:"passphrase"`
	Name       string `json:"name"`
}

// Regitration is a handler to registration process
func (e *Endpoint) Regitration(c *fiber.Ctx) error {
	fLog := e.log.With().Str("func", "Regitration").Logger()

	// parse body request from JSON to struct
	req := &RegistrationRes{}
	err := c.BodyParser(req)
	if err != nil {
		fLog.Error().Msgf("BodyParser error, got: %v", err)
		// build response
		res := response.Factory()
		res.BadRequest("unable processing request")
		return c.Status(res.Code).JSON(res)
	}

	//generate hashed password
	hashedpwd, err := e.password.Hash(req.Passphrase)
	if err != nil {
		fLog.Error().Msgf("Hashing password error, got: %v", err)
		// build response
		res := response.Factory()
		res.InternalServerError("unable hashing passphrase")
		return c.Status(res.Code).JSON(res)
	}

	repo, trx, err := e.repo.BeginTX(c.Context())
	if err != nil {
		fLog.Error().Msgf("unable begin transaction, got: %v", err)
		// build response
		res := response.Factory()
		res.InternalServerError("failed to create record in database")
		return c.Status(res.Code).JSON(res)
	}

	auth, err := repo.CreateAuth(c.Context(), req.Username, hashedpwd)
	if err != nil {
		trx.Rollback()
		fLog.Error().Msgf("unable insert to auth table error, got: %v", err)
		// build response
		res := response.Factory()
		res.InternalServerError("failed to create record in database")
		return c.Status(res.Code).JSON(res)
	}

	_, err = repo.CreateAccountInfo(c.Context(), auth.ID, req.Name)
	if err != nil {
		trx.Rollback()
		fLog.Error().Msgf("unable insert to auth table error, got: %v", err)
		// build response
		res := response.Factory()
		res.InternalServerError("failed to create record in database")
		return c.Status(res.Code).JSON(res)
	}
	trx.Commit()
	// build response
	res := response.Factory()
	res.Created("successfully registered")
	return c.Status(res.Code).JSON(res)
}
