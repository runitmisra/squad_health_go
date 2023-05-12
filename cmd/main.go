package main

import (
	"squad_health_go/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	app.Use(logger.New())

	setupRoutes(app)

	app.Listen(":3000")
}
