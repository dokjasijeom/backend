package controller

import (
	"github.com/dokjasijeom/backend/configuration"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/model"
	"github.com/dokjasijeom/backend/service"
	"github.com/gofiber/fiber/v2"
	"log"
)

func NewCategoryController(seriesService *service.SeriesService, config configuration.Config) *CategoryController {
	return &CategoryController{SeriesService: *seriesService, Config: config}
}

type CategoryController struct {
	service.SeriesService
	configuration.Config
}

func (controller CategoryController) Route(app fiber.Router) {
	category := app.Group("/categories")
	category.Get("/", controller.GetAllCategorySeries)
}

func (controller CategoryController) GetAllCategorySeries(ctx *fiber.Ctx) error {
	type CategoryParameter struct {
		SeriesType entity.SeriesType `query:"seriesType" default:""`
		Genre      string            `query:"genre" default:""`
		Providers  []string          `query:"providers" default:""`
		Page       int               `query:"page" default:"1"`
		PageSize   int               `query:"pageSize" default:"20"`
	}

	var param CategoryParameter
	if err := ctx.QueryParser(&param); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if param.Page < 1 {
		param.Page = 1
	}
	if param.PageSize < 1 {
		param.PageSize = 20
	}

	result, err := controller.SeriesService.GetAllCategorySeries(ctx.Context(), param.SeriesType, param.Genre, param.Providers, param.Page, param.PageSize)
	if err != nil {
		log.Println("여기")
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	for i, v := range result.Series {
		result.Series[i].Thumbnail = controller.Config.Get("CLOUDINARY_URL") + v.Thumbnail
		// series 결과 목록에서 Id 필드값을 제거
		result.Series[i].Id = 0

		// authors for
		for j, _ := range v.Authors {
			result.Series[i].Authors[j].Id = 0
		}
		// publishers for
		for j, _ := range v.Publishers {
			result.Series[i].Publishers[j].Id = 0
		}
		// genres for
		for j, _ := range v.Genres {
			result.Series[i].Genres[j].Id = 0
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}
