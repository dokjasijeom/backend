package main

import (
	"github.com/dokjasijeom/backend/controller"
	"github.com/dokjasijeom/backend/exception"
	repository "github.com/dokjasijeom/backend/repository/impl"
	service "github.com/dokjasijeom/backend/service/impl"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"os"

	"github.com/dokjasijeom/backend/configuration"
)

func main() {
	config := configuration.New()

	database := configuration.ConnectDatabase()

	// repository
	userRepository := repository.NewUserRepositoryImpl(database)
	userRoleRepository := repository.NewUserRoleRepositoryImpl(database)
	//roleRepository := repository.NewRoleRepositoryImpl(database)

	// service
	userService := service.NewUserServiceImpl(&userRepository, &userRoleRepository)
	//roleService := service.NewRoleServiceImpl(&roleRepository)

	// controller
	userController := controller.NewUserController(&userService, config)
	//roleController := controller.NewRoleController(&roleService, config)
	testController := controller.NewTestController(config)

	// setup fiber
	app := fiber.New(configuration.NewFiberConfiguration())
	app.Use(helmet.New())
	//app.Use(csrf.New())
	app.Use(limiter.New())
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	userController.Route(app)
	//roleController.Route(app)
	testController.Route(app)

	//app.Get("/", func(c *fiber.Ctx) error {
	//	return c.SendString("Hello, World!")
	//}

	serverHost := func() string {
		port := os.Getenv("PORT")
		if port == "" {
			port = "3000"
		}
		if serverHost := os.Getenv("HOST"); serverHost != "" {
			return serverHost + ":" + port
		} else {
			return ":" + port
		}
	}()

	err := app.Listen(serverHost)
	exception.PanicLogging(err)
}
