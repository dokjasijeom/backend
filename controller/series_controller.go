package controller

import (
	"github.com/dokjasijeom/backend/configuration"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/model"
	"github.com/dokjasijeom/backend/service"
	"github.com/gofiber/fiber/v2"
)

func NewSeriesController(seriesService *service.SeriesService, config configuration.Config) *SeriesController {
	return &SeriesController{SeriesService: *seriesService, Config: config}
}

type SeriesController struct {
	service.SeriesService
	configuration.Config
}

func (controller SeriesController) Route(app fiber.Router) {
	series := app.Group("/series")
	series.Get("/", controller.GetAllSeries)
	series.Get("/:id", controller.GetSeriesById)
}

func (controller SeriesController) GetAllSeries(ctx *fiber.Ctx) error {
	publishDay := ctx.Query("publishDay")
	seriesType := ctx.Query("seriesType")

	var result []entity.Series
	var err error

	if publishDay != "" && seriesType != "" {
		result, err = controller.SeriesService.GetSeriesByPublishDayAndSeriesType(ctx.Context(), publishDay, seriesType)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
				Code:    fiber.StatusInternalServerError,
				Message: err.Error(),
				Data:    nil,
			})
		}

	} else {
		result, err = controller.SeriesService.GetAllSeries(ctx.Context())
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
				Code:    fiber.StatusInternalServerError,
				Message: err.Error(),
				Data:    nil,
			})
		}
	}

	// result for and i want to change Thumbnail Variable value
	for i, v := range result {
		result[i].Thumbnail = controller.Config.Get("CLOUDINARY_URL") + v.Thumbnail
		// series 결과 목록에서 Id 필드값을 제거
		result[i].Id = 0
	}

	return ctx.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

// get series by id
func (controller SeriesController) GetSeriesById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}

	result, err := controller.SeriesService.GetSeriesById(ctx.Context(), uint(id))
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}
