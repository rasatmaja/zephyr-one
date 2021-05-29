package middleware

import (
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/zephyr-one/pkg/helper"
)

var instance *App
var singleton sync.Once

// App is a struct to define a middleware fiber app
type App struct{ *fiber.App }

// New is a function to initialize middleware
func New(app *fiber.App) *App {
	singleton.Do(func() {
		fmt.Println("[ MDWR ] Starting middleware core ...")
		instance = &App{app}
	})
	return instance
}

// InitializeMiddleware is a function to register and initialize middleware func
func (mdlwr *App) InitializeMiddleware() {
	mdlwr.RequestID()
	//mdlwr.PageNotfound()
}

// PageNotfound is a function to initialize 404 page not found page as a midleware
func (mdlwr *App) PageNotfound() {
	fmt.Println("[ MDWR ] Initialize 404 page notfound middleware")
	mdlwr.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(
			fiber.Map{
				"message": fmt.Sprintf("URL [%s] not found", c.Path()),
			},
		)
	})
}

// RequestID is a function to initialize request id for http header as a midleware
func (mdlwr *App) RequestID() {
	fmt.Println("[ MDWR ] Initialize RequestID middleware")
	mdlwr.Use(func(c *fiber.Ctx) error {
		reqID := c.Get("X-Request-Id")
		if len(reqID) == 0 {
			reqID, _ = helper.GenerateRandomString(8)
			c.Set("X-Request-Id", reqID)
		}
		return c.Next()
	})
}
