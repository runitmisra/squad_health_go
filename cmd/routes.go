package main

import (
	"squad_health_go/handlers"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/users", handlers.ListUsers)
	app.Post("/user", handlers.CreateUser)
}
