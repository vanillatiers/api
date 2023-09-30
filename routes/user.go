package routes

import (
	"fmt"

	"github.com/angelnext/tasks/database"
	"github.com/angelnext/tasks/models"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	db := database.DB

	var users []models.User
	db.Find(&users)
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	db := database.DB

	var user models.User
	db.Where(&models.User{ID: c.Params("id")}).First(&user)
	if user.ID != "" {
		return c.JSON(user)
	}
	return c.Status(404).JSON(fiber.Map{})
}

func CreateUser(c *fiber.Ctx) error {
	db := database.DB

	var user models.User

	if getBodyReqErr := c.BodyParser(&user); getBodyReqErr != nil {
		return c.Status(400).SendString("Invalid JSON Request Body")
	}

	db.Create(&user)

	c.Set("Content-Location", fmt.Sprintf("/api/users/%s", user.ID))

	return c.Status(201).JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	db := database.DB

	db.Delete(&models.User{}, "id = ?", c.Params("id"))
	return c.SendStatus(204)
}

func UpdateUser(c *fiber.Ctx) error {
	db := database.DB

	id := c.Params("id")

	var parsedUser models.User

	if getReqBodyErr := c.BodyParser(&parsedUser); getReqBodyErr != nil {
		return c.Status(400).SendString("Invalid JSON Request Body")
	}

	var user models.User

	db.Where(&models.User{ID: id}).First(&user)

	c.Set("Content-Location", fmt.Sprintf("/api/users/%s", id))

	if parsedUser.Username != "" {
		user.Username = parsedUser.Username
	}

	if parsedUser.PreferredServer != "" {
		user.PreferredServer = parsedUser.PreferredServer
	}

	if parsedUser.Rank != "" {
		user.Rank = parsedUser.Rank
	}

	if parsedUser.Region != "" {
		user.Region = parsedUser.Region
	}

	if user.ID != "" {
		db.Save(&user)

		return c.JSON(user)
	}

	user.ID = id

	db.Create(&user)

	return c.Status(201).JSON(user)
}
