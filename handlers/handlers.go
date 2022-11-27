package handlers

import (
	"squad_health_go/database"
	"squad_health_go/models"

	"github.com/gofiber/fiber/v2"
)

func ListUsers(c *fiber.Ctx) error {
	facts := []models.User{}
	database.DB.Db.Find(&facts)

	return c.Status(200).JSON(facts)
}

func CreateUser(c *fiber.Ctx) error {
	fact := new(models.User)
	err := c.BodyParser(fact)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}
