package routes

import (
	"fmt"
	"golang.org/x/exp/slices"

	"github.com/angelnext/tasks/database"
	"github.com/angelnext/tasks/models"
	"github.com/gofiber/fiber/v2"
)

func GetTickets(c *fiber.Ctx) error {
	db := database.DB

	var tickets []models.Ticket
	db.Find(&tickets)
	return c.JSON(tickets)
}

func GetTicket(c *fiber.Ctx) error {
	db := database.DB

	var ticket models.Ticket
	db.Where(&models.Ticket{ChannelID: c.Params("id")}).First(&ticket)
	if ticket.ChannelID != "" {
		return c.JSON(ticket)
	}
	return c.Status(404).JSON(fiber.Map{})
}

func CreateTicket(c *fiber.Ctx) error {
	db := database.DB

	var ticket models.Ticket

	if getBodyReqErr := c.BodyParser(&ticket); getBodyReqErr != nil {
		return c.Status(400).SendString("Invalid JSON Request Body")
	}

	db.Create(&ticket)

	c.Set("Content-Location", fmt.Sprintf("/api/tickets/%s", ticket.ChannelID))

	return c.Status(201).JSON(ticket)
}

func DeleteTicket(c *fiber.Ctx) error {
	db := database.DB

	db.Delete(&models.Ticket{ChannelID: c.Params("id")})
	return c.SendStatus(204)
}

func AddTestersToTicket(c *fiber.Ctx) error {
	db := database.DB

	id := c.Params("id")
	tester := c.Params("tester")

	var ticket models.Ticket

	db.Where(&models.Ticket{ChannelID: id}).First(&ticket)

	if slices.Contains(ticket.Testers, tester) {
		return c.Status(409).SendString(fmt.Sprintf("Tester with id \"%s\" already exists", tester))
	}

	ticket.Testers = append(ticket.Testers, tester)

	db.Save(&ticket)

	c.Set("Content-Location", fmt.Sprintf("/api/tickets/%s", id))

	return c.JSON(ticket)
}

func RemoveTestersFromTicket(c *fiber.Ctx) error {
	db := database.DB

	id := c.Params("id")
	tester := c.Params("tester")

	var ticket models.Ticket

	db.Where(&models.Ticket{ChannelID: id}).First(&ticket)

	i := slices.Index(ticket.Testers, tester)

	if i == -1 {
		return c.Status(404).SendString(fmt.Sprintf("Tester with id \"%s\" isn't in queue", tester))
	}

	ticket.Testers = slices.Delete(ticket.Testers, i, i+1)

	db.Save(&ticket)

	return c.SendStatus(204)
}
