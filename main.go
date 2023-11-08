package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/dokjasijeom/backend/configuration"
)

func main() {
	app := fiber.New()

	configuration.ConnectDatabase()

	app.Use(helmet.New())
	app.Use(csrf.New())
	app.Use(limiter.New())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		configuration.TestDataBase()
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
