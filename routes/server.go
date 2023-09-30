package routes

import (
	"fmt"

	"github.com/angelnext/tasks/database"
	"github.com/angelnext/tasks/models"
	"github.com/gofiber/fiber/v2"
)

func GetServers(c *fiber.Ctx) error {
	db := database.DB

	var servers []models.Server
	db.Find(&servers)
	return c.JSON(servers)
}

func GetServer(c *fiber.Ctx) error {
	db := database.DB

	var server models.Server
	db.Where(&models.Server{ID: c.Params("id")}).First(&server)
	if server.ID != "" {
		return c.JSON(server)
	}
	return c.Status(404).JSON(fiber.Map{})
}

func CreateServer(c *fiber.Ctx) error {
	db := database.DB

	var server models.Server

	if getBodyReqErr := c.BodyParser(&server); getBodyReqErr != nil {
		return c.Status(400).SendString("Invalid JSON Request Body")
	}

	db.Create(&server)

	c.Set("Content-Location", fmt.Sprintf("/api/servers/%s", server.ID))

	return c.Status(201).JSON(server)
}

func DeleteServer(c *fiber.Ctx) error {
	db := database.DB

	db.Delete(&models.Server{}, "id = ?", c.Params("id"))
	return c.SendStatus(204)
}

func UpdateServer(c *fiber.Ctx) error {
	db := database.DB

	id := c.Params("id")

	var parsedServer models.Server

	if getReqBodyErr := c.BodyParser(&parsedServer); getReqBodyErr != nil {
		return c.Status(400).SendString("Invalid JSON Request Body")
	}

	var server models.Server

	db.Where(&models.Server{ID: id}).First(&server)

	c.Set("Content-Location", fmt.Sprintf("/api/servers/%s", id))

	if parsedServer.Cooldown != 0 {
		server.Cooldown = parsedServer.Cooldown
	}

	if parsedServer.LT3PlusCooldown != 0 {
		server.LT3PlusCooldown = parsedServer.LT3PlusCooldown
	}

	if parsedServer.TicketAutoclose != 0 {
		server.TicketAutoclose = parsedServer.TicketAutoclose
	}

	if parsedServer.QueueAutoclose != 0 {
		server.QueueAutoclose = parsedServer.QueueAutoclose
	}

	if parsedServer.QueueChannelID != "" {
		server.QueueChannelID = parsedServer.QueueChannelID
	}

	if parsedServer.LogsChannelID != "" {
		server.LogsChannelID = parsedServer.LogsChannelID
	}

	if parsedServer.EURoleID != "" {
		server.EURoleID = parsedServer.EURoleID
	}

	if parsedServer.NARoleID != "" {
		server.NARoleID = parsedServer.NARoleID
	}

	if parsedServer.HT1RoleID != "" {
		server.HT1RoleID = parsedServer.HT1RoleID
	}

	if parsedServer.LT1RoleID != "" {
		server.LT1RoleID = parsedServer.LT1RoleID
	}

	if parsedServer.HT2RoleID != "" {
		server.HT2RoleID = parsedServer.HT2RoleID
	}

	if parsedServer.LT2RoleID != "" {
		server.LT2RoleID = parsedServer.LT2RoleID
	}

	if parsedServer.HT3RoleID != "" {
		server.HT3RoleID = parsedServer.HT3RoleID
	}

	if parsedServer.LT3RoleID != "" {
		server.LT3RoleID = parsedServer.LT3RoleID
	}

	if parsedServer.HT4RoleID != "" {
		server.HT4RoleID = parsedServer.HT4RoleID
	}

	if parsedServer.LT4RoleID != "" {
		server.LT4RoleID = parsedServer.LT4RoleID
	}

	if parsedServer.HT5RoleID != "" {
		server.HT5RoleID = parsedServer.HT5RoleID
	}

	if parsedServer.LT5RoleID != "" {
		server.LT5RoleID = parsedServer.LT5RoleID
	}

	if parsedServer.EUTicketsCategoryID != "" {
		server.EUTicketsCategoryID = parsedServer.EUTicketsCategoryID
	}

	if parsedServer.NATicketsCategoryID != "" {
		server.NATicketsCategoryID = parsedServer.NATicketsCategoryID
	}

	if parsedServer.HT3TicketsCategoryID != "" {
		server.HT3TicketsCategoryID = parsedServer.HT3TicketsCategoryID
	}

	if server.ID != "" {
		db.Save(&server)

		return c.JSON(server)
	}

	server.ID = id
	db.Create(&server)

	return c.Status(201).JSON(server)
}
