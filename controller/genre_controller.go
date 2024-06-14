package controller

import (
	"github.com/dokjasijeom/backend/configuration"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/service"
	"github.com/gofiber/fiber/v2"
)

func NewGenreController(genreService *service.GenreService, config configuration.Config) *GenreController {
	return &GenreController{GenreService: *genreService, Config: config}
}

type GenreController struct {
	service.GenreService
	configuration.Config
}

func (controller GenreController) Route(app fiber.Router) {
	genre := app.Group("/genres")
	genre.Get("/", controller.GetAllMainGenre)
}

func (controller GenreController) GetAllMainGenre(ctx *fiber.Ctx) error {
	// seriesType의 기본값은 webnovel
	seriesType := ctx.Query("seriesType", "webnovel")

	result, err := controller.GenreService.GetAllMainGenre(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
	}

	// seriesType으로 필터링
	// GenreType이 common 이거나 seriesType과 같은 GenreType만 반환
	for i := 0; i < len(result); i++ {
		result[i].Id = 0
		if result[i].GenreType != entity.GenreType(seriesType) && result[i].GenreType != "common" {
			result = append(result[:i], result[i+1:]...)
			i--
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "success",
		"data":    result,
	})
}
