package controller

import (
	"github.com/dokjasijeom/backend/configuration"
	"github.com/dokjasijeom/backend/exception"
	"github.com/dokjasijeom/backend/service"
	"github.com/gofiber/fiber/v2"
)

func NewRoleController(roleService *service.RoleService, config configuration.Config) *RoleController {
	return &RoleController{RoleService: *roleService, Config: config}
}

type RoleController struct {
	service.RoleService
	configuration.Config
}

func (controller RoleController) Route(app *fiber.App) {
	app.Post("/roles", controller.CreateRole)
}

func (controller RoleController) CreateRole(ctx *fiber.Ctx) error {
	var request struct {
		Name string `json:"name" validate:"required"`
	}
	err := ctx.BodyParser(&request)
	if err != nil {
		exception.PanicLogging(err)
		return err
	}

	result, _ := controller.RoleService.CreateRole(ctx.Context(), request.Name)

	return ctx.Status(fiber.StatusOK).JSON(result)
}
