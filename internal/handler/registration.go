package handler

import (
	"github.com/gofiber/fiber/v2"
	zosql "github.com/rasatmaja/zephyr-one/internal/database/sql"
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

	// build response
	res := response.Factory()

	// parse body request from JSON to struct
	req := &RegistrationRes{}
	err := c.BodyParser(req)
	if err != nil {
		fLog.Error().Msgf("BodyParser error, got: %v", err)
		return res.BadRequest("unable processing request")
	}

	//generate hashed password
	hashedpwd, err := e.password.Hash(req.Passphrase)
	if err != nil {
		fLog.Error().Msgf("Hashing password error, got: %v", err)
		return res.InternalServerError("unable hashing passphrase")
	}

	// begin transaction
	repo, trx, err := e.repo.BeginTX(c.Context())
	if err != nil {
		fLog.Error().Msgf("unable begin transaction, got: %v", err)
		return res.InternalServerError("failed to create record in database")
	}

	auth, err := repo.CreateAuth(c.Context(), req.Username, hashedpwd)
	if err != nil {
		trx.Rollback()
		fLog.Error().Msgf("unable insert to auth table error, got: %v", err)
		if err == zosql.ErrDataDuplicate {
			return res.BadRequest("Username already exist")
		}
		return res.InternalServerError("failed to create record in database")
	}

	_, err = repo.CreateAccountInfo(c.Context(), auth.ID, req.Name)
	if err != nil {
		trx.Rollback()
		fLog.Error().Msgf("unable insert to auth table error, got: %v", err)
		return res.InternalServerError("failed to create record in database")
	}
	trx.Commit()

	res.Created("successfully registered")
	return c.Status(res.Code).JSON(res)
}
