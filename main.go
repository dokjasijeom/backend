package main

import (
	"github.com/dokjasijeom/backend/controller"
	"github.com/dokjasijeom/backend/controller/backoffice"
	"github.com/dokjasijeom/backend/exception"
	"github.com/dokjasijeom/backend/middleware"
	repository "github.com/dokjasijeom/backend/repository/impl"
	service "github.com/dokjasijeom/backend/service/impl"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/helmet"
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
	seriesRepository := repository.NewSeriesRepositoryImpl(database)
	genreRepository := repository.NewGenreRepositoryImpl(database)
	providerRepository := repository.NewProviderRepositoryImpl(database)
	publisherRepoistory := repository.NewPublisherRepositoryImpl(database)
	publishDayRepository := repository.NewPublishDayRepositoryImpl(database)
	personRepository := repository.NewPersonRepositoryImpl(database)
	episodeRepository := repository.NewEpisodeRepositoryImpl(database)
	seriesDailyViewRepository := repository.NewSeriesDailyViewRepositoryImpl(database)
	userRecordSeriesRepository := repository.NewUserRecordSeriesRepositoryImpl(database)
	userRecordSeriesEpisodeRepository := repository.NewUserRecordSeriesEpisodeRepositoryImpl(database)
	userProfileRepository := repository.NewUserProfileRepositoryImpl(database)

	// service
	userService := service.NewUserServiceImpl(&userRepository, &userRoleRepository, &userProfileRepository)
	//roleService := service.NewRoleServiceImpl(&roleRepository)
	seriesService := service.NewSeriesServiceImpl(&seriesRepository)
	genreService := service.NewGenreServiceImpl(&genreRepository)
	providerService := service.NewProviderServiceImpl(&providerRepository)
	publisherService := service.NewPublisherServiceImpl(&publisherRepoistory)
	publishDayService := service.NewPublishDayServiceImpl(&publishDayRepository)
	personService := service.NewPersonServiceImpl(&personRepository)
	episodeService := service.NewEpisodeServiceImpl(&episodeRepository)
	seriesDailyViewService := service.NewSeriesDailyViewServiceImpl(&seriesDailyViewRepository)
	userRecordSeriesService := service.NewUserRecordSeriesServiceImpl(&userRecordSeriesRepository)
	userRecordSeriesEpisodeService := service.NewUserRecordSeriesEpisodeServiceImpl(&userRecordSeriesEpisodeRepository)

	// controller
	userController := controller.NewUserController(&userService, &seriesService, &episodeService, &providerService, &userRecordSeriesService, &userRecordSeriesEpisodeService, config)
	//roleController := controller.NewRoleController(&roleService, config)
	testController := controller.NewTestController(config)
	seriesController := controller.NewSeriesController(&seriesService, &seriesDailyViewService, &userService, &userRecordSeriesService, &userRecordSeriesEpisodeService, config)
	searchController := controller.NewSearchController(&seriesService, config)
	genreController := controller.NewGenreController(&genreService, config)
	categoryController := controller.NewCategoryController(&seriesService, config)
	providerController := controller.NewProviderController(&providerService, config)

	backofficeSeriesController := backoffice.NewBackofficeSeriesController(&seriesService, &episodeService, config)
	backofficeGenreController := backoffice.NewBackofficeGenreController(&genreService, config)
	backofficeProviderController := backoffice.NewBackofficeProviderController(&providerService, config)
	backofficePublisherController := backoffice.NewBackofficePublisherController(&publisherService, config)
	backofficePublishDayController := backoffice.NewBackofficePublishDayController(&publishDayService, config)
	backofficePersonController := backoffice.NewBackofficePersonController(&personService, config)

	// setup fiber
	app := fiber.New(configuration.NewFiberConfiguration())
	app.Use(helmet.New())
	//app.Use(csrf.New())
	//app.Use(limiter.New())
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(favicon.New(favicon.Config{
		File: "./favicon.ico",
		URL:  "/favicon.ico",
	}))

	userController.Route(app)
	//roleController.Route(app)
	testController.Route(app)
	seriesController.Route(app)
	searchController.Route(app)
	genreController.Route(app)
	categoryController.Route(app)
	providerController.Route(app)

	backoffice := app.Group("/backoffice", middleware.AuthenticateJWT("ADMIN", config))
	backofficeSeriesController.Route(backoffice)
	backofficeGenreController.Route(backoffice)
	backofficeProviderController.Route(backoffice)
	backofficePublisherController.Route(backoffice)
	backofficePublishDayController.Route(backoffice)
	backofficePersonController.Route(backoffice)

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
