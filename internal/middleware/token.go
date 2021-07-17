package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/zephyr-one/internal/response"
)

// VerifyToken define middleware to verify Authorization header
func (mdlwr *App) VerifyToken(c *fiber.Ctx) error {

	// build response
	res := response.Factory()

	authHeader := c.Get("Authorization")
	if len(authHeader) == 0 {
		return res.Unauthorized("Authorization header empty")
	}

	// get token bearer prefix
	tokenType := strings.ToLower(strings.TrimSpace(authHeader[:6]))
	if tokenType != "bearer" {
		return res.Unauthorized("Authorization header missing `bearer` prefix")
	}

	// get token string
	token := strings.TrimSpace(authHeader[7:])

	payload, err := mdlwr.token.Verify(c.Context(), token)

	if err != nil {
		return res.Unauthorized(err.Error())
	}

	mdlwr.log.Info().Msgf("%v/n", payload)
	return c.Next()

}
