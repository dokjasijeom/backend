package backoffice

import (
	"github.com/dokjasijeom/backend/configuration"
	"github.com/dokjasijeom/backend/model"
	"github.com/dokjasijeom/backend/service"
	"github.com/gofiber/fiber/v2"
)

func NewBackofficeGenreController(genreService *service.GenreService, config configuration.Config) *BackofficeGenreController {
	return &BackofficeGenreController{GenreService: *genreService, Config: config}
}

type BackofficeGenreController struct {
	service.GenreService
	configuration.Config
}

func (controller BackofficeGenreController) Route(app fiber.Router) {
	genre := app.Group("/genres")
	genre.Post("/", controller.CreateGenre)
	genre.Get("/", controller.GetAllMainGenre)
	genre.Get("/:id/sub", controller.GetAllSubGenre)
	genre.Delete("/:id", controller.DeleteGenreById)
	genre.Put("/:id", controller.UpdateGenreById)
}

// Create Genre
func (controller BackofficeGenreController) CreateGenre(ctx *fiber.Ctx) error {
	var request = model.GenreRequestModel{}
	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	result, err := controller.GenreService.CreateGenre(ctx.Context(), request.Name, request.GenreType, request.ParentGenreId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
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

// Get All Main Genre
func (controller BackofficeGenreController) GetAllMainGenre(ctx *fiber.Ctx) error {
	result, err := controller.GenreService.GetAllMainGenre(ctx.Context())
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

// Get All Sub Genre
func (controller BackofficeGenreController) GetAllSubGenre(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", 0)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	result, err := controller.GenreService.GetAllSubGenre(ctx.Context(), uint(id))
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

// Delete Genre By Id
func (controller BackofficeGenreController) DeleteGenreById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", 0)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = controller.GenreService.DeleteGenre(ctx.Context(), uint(id))
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

// Update Genre By Id
func (controller BackofficeGenreController) UpdateGenreById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", 0)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var request = model.GenreRequestModel{}
	err = ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = controller.GenreService.UpdateGenre(ctx.Context(), uint(id), request.Name)
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
