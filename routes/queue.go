package routes

import (
	"fmt"
	"slices"

	"github.com/angelnext/tasks/database"
	"github.com/angelnext/tasks/models"
	"github.com/gofiber/fiber/v2"
)

func GetQueues(c *fiber.Ctx) error {
	db := database.DB

	var queues []models.Queue
	db.Find(&queues)
	return c.JSON(queues)
}

func GetQueue(c *fiber.Ctx) error {
	db := database.DB

	var queue models.Queue
	db.Where(&models.Queue{MessageID: c.Params("id")}).First(&queue)

	if queue.MessageID != "" {
		return c.JSON(queue)
	}
	return c.Status(404).JSON(fiber.Map{})
}

func CreateQueue(c *fiber.Ctx) error {
	db := database.DB

	var queue models.Queue

	if getBodyReqErr := c.BodyParser(&queue); getBodyReqErr != nil {
		return c.Status(400).SendString("Invalid JSON Request Body")
	}

	db.Create(&queue)

	c.Set("Content-Location", fmt.Sprintf("/api/users/%s", queue.MessageID))

	return c.Status(201).JSON(queue)
}

func DeleteQueue(c *fiber.Ctx) error {
	db := database.DB

	db.Delete(&models.Queue{MessageID: c.Params("id")})
	return c.SendStatus(204)
}

func AddTestersToQueue(c *fiber.Ctx) error {
	db := database.DB

	id := c.Params("id")
	tester := c.Params("tester")

	var queue models.Queue

	db.Where(&models.Queue{MessageID: id}).First(&queue)

	if slices.Contains(queue.Testers, tester) {
		return c.Status(409).SendString(fmt.Sprintf("Tester with id \"%s\" already exists", tester))
	}

	queue.Testers = append(queue.Testers, tester)

	db.Save(&queue)

	return c.JSON(queue)
}

func AddMembersToQueue(c *fiber.Ctx) error {
	db := database.DB

	id := c.Params("id")
	member := c.Params("member")

	var queue models.Queue

	db.Where(&models.Queue{MessageID: id}).First(&queue)

	if slices.Contains(queue.Members, member) {
		return c.Status(409).SendString(fmt.Sprintf("Member with id \"%s\" already exists", member))
	}

	queue.Members = append(queue.Members, member)

	db.Save(&queue)

	return c.JSON(queue)
}

func RemoveTestersFromQueue(c *fiber.Ctx) error {
	db := database.DB

	id := c.Params("id")
	tester := c.Params("tester")

	var queue models.Queue

	db.Where(&models.Queue{MessageID: id}).First(&queue)

	i := slices.Index(queue.Testers, tester)

	if i == -1 {
		return c.Status(404).SendString(fmt.Sprintf("Tester with id \"%s\" isn't in queue", tester))
	}

	queue.Testers = slices.Delete(queue.Testers, i, i+1)

	db.Save(&queue)

	return c.SendStatus(204)
}

func RemoveMembersFromQueue(c *fiber.Ctx) error {
	db := database.DB

	id := c.Params("id")
	member := c.Params("member")

	var queue models.Queue

	db.Where(&models.Queue{MessageID: id}).First(&queue)

	i := slices.Index(queue.Members, member)

	if i == -1 {
		return c.Status(404).SendString(fmt.Sprintf("Member with id \"%s\" isn't in queue", member))
	}

	queue.Members = slices.Delete(queue.Members, i, i+1)

	db.Save(&queue)

	return c.SendStatus(204)
}
