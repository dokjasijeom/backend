package controller

import (
	"github.com/dokjasijeom/backend/configuration"
	"github.com/dokjasijeom/backend/model"
	"github.com/dokjasijeom/backend/service"
	"github.com/gofiber/fiber/v2"
)

func NewBackofficeProviderController(providerService *service.ProviderService, config configuration.Config) *BackofficeProviderController {
	return &BackofficeProviderController{ProviderService: *providerService, Config: config}
}

type BackofficeProviderController struct {
	service.ProviderService
	configuration.Config
}

func (controller BackofficeProviderController) Route(app fiber.Router) {
	provider := app.Group("/providers")
	provider.Post("/", controller.CreateProvider)
	provider.Get("/", controller.GetAllProvider)
	provider.Get("/:id", controller.GetProviderById)
	provider.Put("/:id", controller.UpdateProviderById)
	provider.Delete("/:id", controller.DeleteProviderById)
}

// Create Provider
func (controller BackofficeProviderController) CreateProvider(ctx *fiber.Ctx) error {
	var request = model.ProviderRequestModel{}
	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	result, err := controller.ProviderService.CreateProvider(ctx.Context(), request.Name, request.DisplayName, request.Description, request.HomepageUrl)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    fiber.StatusCreated,
		Message: "Success",
		Data:    result,
	})
}

// Get All Provider
func (controller BackofficeProviderController) GetAllProvider(ctx *fiber.Ctx) error {
	result, err := controller.ProviderService.GetAllProvider(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

// Get Provider by Id
func (controller BackofficeProviderController) GetProviderById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}

	result, err := controller.ProviderService.GetProviderById(ctx.Context(), uint(id))
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

// Update Provider by Id
func (controller BackofficeProviderController) UpdateProviderById(ctx *fiber.Ctx) error {
	var request = model.ProviderRequestModel{}
	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}

	err = controller.ProviderService.UpdateProvider(ctx.Context(), uint(id), request.Name, request.DisplayName, request.Description, request.HomepageUrl)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    nil,
	})
}

// Delete Provider by Id
func (controller BackofficeProviderController) DeleteProviderById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", 0)
	if err != nil {
		return err
	}

	err = controller.ProviderService.DeleteProvider(ctx.Context(), uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.Status(fiber.StatusNoContent).JSON(model.GeneralResponse{
		Code:    fiber.StatusNoContent,
		Message: "Success",
		Data:    nil,
	})
}
