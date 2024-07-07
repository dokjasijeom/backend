package controller

import (
	"github.com/dokjasijeom/backend/configuration"
	"github.com/dokjasijeom/backend/model"
	"github.com/dokjasijeom/backend/service"
	"github.com/gofiber/fiber/v2"
)

func NewProviderController(providerSerivce *service.ProviderService, config configuration.Config) *ProviderController {
	return &ProviderController{ProviderService: *providerSerivce, Config: config}
}

type ProviderController struct {
	service.ProviderService
	configuration.Config
}

func (controller ProviderController) Route(app fiber.Router) {
	provider := app.Group("/providers")
	provider.Get("/", controller.GetAllProviders)
}

func (controller ProviderController) GetAllProviders(ctx *fiber.Ctx) error {
	result, err := controller.ProviderService.GetAllProvider(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	for i, _ := range result {
		result[i].Id = 0
	}

	return ctx.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}
