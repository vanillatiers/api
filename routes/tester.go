package routes

import (
	"fmt"

	"github.com/angelnext/tasks/database"
	"github.com/angelnext/tasks/models"
	"github.com/gofiber/fiber/v2"
)

func GetTesters(c *fiber.Ctx) error {
	db := database.DB

	var testers []models.Tester
	db.Find(&testers)
	return c.JSON(testers)
}

func GetTester(c *fiber.Ctx) error {
	db := database.DB

	var tester models.Tester
	db.Where(&models.Tester{UserID: c.Params("id")}).First(&tester)
	if tester.UserID != "" {
		return c.JSON(tester)
	}
	return c.Status(404).JSON(fiber.Map{})
}

func CreateTester(c *fiber.Ctx) error {
	db := database.DB

	var tester models.Tester

	if getBodyReqErr := c.BodyParser(&tester); getBodyReqErr != nil {
		return c.Status(400).SendString("Invalid JSON Request Body")
	}

	db.Create(&tester)

	c.Set("Content-Location", fmt.Sprintf("/api/testers/%s", tester.UserID))

	return c.Status(201).JSON(tester)
}

func DeleteTester(c *fiber.Ctx) error {
	db := database.DB

	db.Delete(&models.Tester{UserID: c.Params("id")})
	return c.SendStatus(204)
}

func UpdateTester(c *fiber.Ctx) error {
	db := database.DB

	id := c.Params("id")

	var tester models.Tester

	db.Where(&models.Tester{UserID: id}).First(&tester)

	c.Set("Content-Location", fmt.Sprintf("/api/testers/%s", id))

	if tester.UserID != "" {
		tester.TotalTests++

		db.Save(&tester)

		return c.JSON(tester)
	}

	tester.UserID = id
	tester.TotalTests = 1

	db.Create(&tester)

	return c.Status(201).JSON(tester)
}
