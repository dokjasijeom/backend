package main

import (
	"github.com/dokjasijeom/backend/controller"
	repository "github.com/dokjasijeom/backend/repository/impl"
	service "github.com/dokjasijeom/backend/service/impl"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/dokjasijeom/backend/configuration"
)

func main() {
	app := fiber.New()

	database := configuration.ConnectDatabase()

	// repository
	userRepository := repository.NewUserRepositoryImpl(database)

	// service
	userService := service.NewUserServiceImpl(&userRepository)

	// controller
	userController := controller.NewUserController(&userService)

	app.Use(helmet.New())
	app.Use(csrf.New())
	app.Use(limiter.New())
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	userController.Route(app)

	//app.Get("/", func(c *fiber.Ctx) error {
	//	return c.SendString("Hello, World!")
	//})

	app.Listen(":3000")
}
