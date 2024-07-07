package backoffice

import (
	"github.com/dokjasijeom/backend/configuration"
	"github.com/dokjasijeom/backend/model"
	"github.com/dokjasijeom/backend/service"
	"github.com/gofiber/fiber/v2"
)

func NewBackofficePublishDayController(publishDayService *service.PublishDayService, config configuration.Config) *BackofficePublishDayController {
	return &BackofficePublishDayController{PublishDayService: *publishDayService, Config: config}
}

type BackofficePublishDayController struct {
	service.PublishDayService
	configuration.Config
}

func (controller BackofficePublishDayController) Route(app fiber.Router) {
	publishDay := app.Group("/publish-days")
	publishDay.Post("/", controller.CreatePublishDay)
	publishDay.Get("/", controller.GetAllPublishDay)
	publishDay.Get("/:id", controller.GetPublishDayById)
	publishDay.Put("/:id", controller.UpdatePublishDayById)
	publishDay.Delete("/:id", controller.DeletePublishDayById)
}

// Create Publish Day
func (controller BackofficePublishDayController) CreatePublishDay(ctx *fiber.Ctx) error {
	var request = model.PublishDayRequestModel{}
	err := ctx.BodyParser(&request)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	result, err := controller.PublishDayService.CreatePublishDay(ctx.Context(), request.Day, request.DisplayDay, request.DisplayOrder)
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

// Get All Publish Day
func (controller BackofficePublishDayController) GetAllPublishDay(ctx *fiber.Ctx) error {
	result, err := controller.PublishDayService.GetAllPublishDay(ctx.Context())
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

// Get Publish Day By Id
func (controller BackofficePublishDayController) GetPublishDayById(ctx *fiber.Ctx) error {
	publishDayId, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	result, err := controller.PublishDayService.GetPublishDayById(ctx.Context(), uint(publishDayId))
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

// Update Publish Day By Id
func (controller BackofficePublishDayController) UpdatePublishDayById(ctx *fiber.Ctx) error {
	publishDayId, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var request = model.PublishDayRequestModel{}
	err = ctx.BodyParser(&request)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = controller.PublishDayService.UpdatePublishDay(ctx.Context(), uint(publishDayId), request.Day, request.DisplayDay, request.DisplayOrder)
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

// Delete Publish Day By Id
func (controller BackofficePublishDayController) DeletePublishDayById(ctx *fiber.Ctx) error {
	publishDayId, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = controller.PublishDayService.DeletePublishDay(ctx.Context(), uint(publishDayId))
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
