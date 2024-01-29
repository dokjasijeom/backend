package backoffice

import (
	"github.com/dokjasijeom/backend/configuration"
	"github.com/dokjasijeom/backend/model"
	"github.com/dokjasijeom/backend/service"
	"github.com/gofiber/fiber/v2"
)

func NewBackofficePublisherController(publisherService *service.PublisherService, config configuration.Config) *BackofficePublisherController {
	return &BackofficePublisherController{PublisherService: *publisherService, Config: config}
}

type BackofficePublisherController struct {
	service.PublisherService
	configuration.Config
}

func (controller BackofficePublisherController) Route(app fiber.Router) {
	publisher := app.Group("/publishers")
	publisher.Post("/", controller.CreatePublisher)
	publisher.Get("/", controller.GetAllPublisher)
	publisher.Get("/:id", controller.GetPublisherById)
	publisher.Put("/:id", controller.UpdatePublisherById)
	publisher.Delete("/:id", controller.DeletePublisherById)
}

// Create Publisher
func (controller BackofficePublisherController) CreatePublisher(ctx *fiber.Ctx) error {
	var request = model.PublisherRequestModel{}
	err := ctx.BodyParser(&request)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	result, err := controller.PublisherService.CreatePublisher(ctx.Context(), request.Name, request.Description, request.HomepageUrl)
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

// Get All Publisher
func (controller BackofficePublisherController) GetAllPublisher(ctx *fiber.Ctx) error {
	result, err := controller.PublisherService.GetAllPublisher(ctx.Context())
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

// Get Publisher by Id
func (controller BackofficePublisherController) GetPublisherById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	result, err := controller.PublisherService.GetPublisherById(ctx.Context(), uint(id))
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

// Update Publisher by Id
func (controller BackofficePublisherController) UpdatePublisherById(ctx *fiber.Ctx) error {
	var request = model.PublisherRequestModel{}
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
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	result, err := controller.PublisherService.UpdatePublisher(ctx.Context(), uint(id), request.Name, request.Description, request.HomepageUrl)
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

// Delete Publisher by Id
func (controller BackofficePublisherController) DeletePublisherById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = controller.PublisherService.DeletePublisherById(ctx.Context(), uint(id))
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
