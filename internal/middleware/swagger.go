package middleware

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed swagger/*
var swaggerDir embed.FS

// SwaggerUI is a middleware to server Swagger UI
func (mdlwr *App) SwaggerUI() {
	fmt.Println("[ MDWR ] Initialize Swagger UI")
	mdlwr.server.Use("/docs", filesystem.New(filesystem.Config{
		Root:       http.FS(swaggerDir),
		PathPrefix: "swagger",
		Browse:     true,
	}))

	mdlwr.server.Get("/docs", func(c *fiber.Ctx) error {
		slashPresent := c.Path()[len(c.Path())-1:] == "/"
		if slashPresent {
			return c.Redirect("index.html")
		}
		return c.Redirect("docs/index.html")
	})
}
