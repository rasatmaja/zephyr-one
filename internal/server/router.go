package server

import "github.com/gofiber/fiber/v2"

// InitializeRoute is a function to define routes and it's handlers
func (a *App) InitializeRoute() {
	defer a.server.Use(a.handler.PageNotfound)

	a.server.Get("/", a.handler.Base)

	a.server.Post("/", a.handler.Base)

	a.server.Get("/hello", a.handler.HelloWorld)

	api := a.server.Group("api")
	a.v1(api)
}

func (a *App) v1(router fiber.Router) {
	v1 := router.Group("v1")
	v1.Post("/register", a.handler.Regitration)
	v1.Post("/login", a.handler.Auth)

	user := v1.Group("user")
	contact := user.Group("contact").Use(a.middleware.VerifyToken)
	contact.Post("/", a.handler.AddContact)
	contact.Get("/:type?", a.handler.Contact)
	contact.Put("/:contact/primary", a.handler.SetPrimaryContact)
	contact.Delete("/:contact", a.handler.RemoveContact)

	info := user.Group("info").Use(a.middleware.VerifyToken)
	info.Get("/", a.handler.Info)
}
