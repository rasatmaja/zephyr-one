package server

// InitializeRoute is a function to define routes and it's handlers
func (a *App) InitializeRoute() {
	defer a.server.Use(a.handler.PageNotfound)

	a.server.Get("/", a.handler.Base)

	a.server.Post("/", a.handler.Base)

	a.server.Get("/hello", a.handler.HelloWorld)

	a.server.Post("/register", a.handler.Regitration)

	a.server.Post("/login", a.handler.Auth)
}
