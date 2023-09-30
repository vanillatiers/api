package main

import (
	"fmt"
	"log"

	"github.com/angelnext/tasks/database"
	"github.com/angelnext/tasks/models"
	"github.com/angelnext/tasks/routes"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	var dbOpenErr error
	if database.DB, dbOpenErr = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DB")))); dbOpenErr != nil {
		log.Fatalln(dbOpenErr)
	}

	if databaseMigrateErr := database.DB.AutoMigrate(&models.Ticket{}, &models.Tester{}, &models.User{}, &models.Queue{}, &models.Server{}, &models.APIKey{}); databaseMigrateErr != nil {
		log.Fatalln("Error while migrating database")
	}

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		AppName:       "Vanilla Tiers Network API",
		Prefork:       true,
	})

  app.Use(helmet.New())

	app.Use("*", func(c *fiber.Ctx) error {
		c.Set("Cache-Control", "max-age=604800")

		db := database.DB

		key := c.Query("apikey")

		var api models.APIKey

		db.Where(&models.APIKey{Key: key}).First(&api)

		if api.Key == "" || api.Perms == 0 {
			return c.Status(401).SendString("You don't have access to this method")
		} else if api.Perms == 1 {
			if c.Method() != "GET" {
				return c.Status(401).SendString("You don't have access to this method")
			}
			c.Next()
			return nil
		} else {
			c.Next()
			return nil
		}
	})

	app.Get("/api/testers", routes.GetTesters)
	app.Get("/api/testers/:id", routes.GetTester)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Get("/api/servers", routes.GetServers)
	app.Get("/api/servers/:id", routes.GetServer)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Get("/api/queues", routes.GetQueues)
	app.Get("/api/queues/:id", routes.GetQueue)
	app.Get("/api/tickets", routes.GetTickets)
	app.Get("/api/tickets/:id", routes.GetTicket)

	app.Post("/api/testers", routes.CreateTester)
	app.Post("/api/users", routes.CreateUser)
	app.Post("/api/servers", routes.CreateServer)
	app.Post("/api/queues", routes.CreateQueue)
	app.Post("/api/tickets", routes.CreateTicket)

	app.Delete("/api/testers/:id", routes.DeleteTester)
	app.Delete("/api/users/:id", routes.DeleteUser)
	app.Delete("/api/servers/:id", routes.DeleteServer)
	app.Delete("/api/queues/:id", routes.DeleteQueue)
	app.Delete("/api/queues/:id/testers/:tester", routes.RemoveTestersFromQueue)
	app.Delete("/api/queues/:id/members/:member", routes.RemoveMembersFromQueue)
	app.Delete("/api/tickets/:id", routes.DeleteTicket)
	app.Delete("/api/tickets/:id/testers/:tester", routes.RemoveTestersFromTicket)

	app.Put("/api/testers/:id", routes.UpdateTester)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Put("/api/servers/:id", routes.UpdateServer)
	app.Put("/api/queues/:id/testers/:tester", routes.AddTestersToQueue)
	app.Put("/api/queues/:id/members/:member", routes.AddMembersToQueue)
	app.Put("/api/tickets/:id/testers/:tester", routes.AddTestersToTicket)

	log.Fatalln(app.Listen(":4000"))
}
