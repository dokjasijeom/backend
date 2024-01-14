package controller

import (
	"github.com/dokjasijeom/backend/configuration"
	"github.com/dokjasijeom/backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func NewBackofficeController(config configuration.Config) *BackofficeController {
	return &BackofficeController{Config: config}
}

type BackofficeController struct {
	configuration.Config
}

func (controller BackofficeController) Route(app *fiber.App) {
	backoffice := app.Group("/backoffice", middleware.AuthenticateJWT("ADMIN", controller.Config))

}
