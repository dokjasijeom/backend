package controller

import (
	"github.com/dokjasijeom/backend/configuration"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/middleware"
	"github.com/dokjasijeom/backend/model"
	"github.com/dokjasijeom/backend/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func NewSeriesController(seriesService *service.SeriesService, userService *service.UserService, config configuration.Config) *SeriesController {
	return &SeriesController{SeriesService: *seriesService, UserService: *userService, Config: config}
}

type SeriesController struct {
	service.SeriesService
	service.UserService
	configuration.Config
}

func (controller SeriesController) Route(app fiber.Router) {
	series := app.Group("/series")
	series.Get("/", controller.GetAllSeries)
	series.Get("/:hashId", controller.GetSeriesByHashId)
	series.Post("/:hashId/like", middleware.AuthenticateJWT("ANY", controller.Config), controller.LikeSeries)
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

		// authors for
		for j, _ := range v.Authors {
			result[i].Authors[j].Id = 0
		}
		// publishers for
		for j, _ := range v.Publishers {
			result[i].Publishers[j].Id = 0
		}
		// genres for
		for j, _ := range v.Genres {
			result[i].Genres[j].Id = 0
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

// get series by hash id
func (controller SeriesController) GetSeriesByHashId(ctx *fiber.Ctx) error {
	hashId := ctx.Params("hashId")
	if hashId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid hashId",
			Data:    nil,
		})
	}

	result, err := controller.SeriesService.GetSeriesByHashId(ctx.Context(), hashId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	result.Id = 0
	result.Thumbnail = controller.Config.Get("CLOUDINARY_URL") + result.Thumbnail

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

	result.Id = 0
	result.Thumbnail = controller.Config.Get("CLOUDINARY_URL") + result.Thumbnail
	// authors for
	for j, _ := range result.Authors {
		result.Authors[j].Id = 0
	}
	// publishers for
	for j, _ := range result.Publishers {
		result.Publishers[j].Id = 0
	}
	// genres for
	for j, _ := range result.Genres {
		result.Genres[j].Id = 0
	}

	return ctx.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

// like series
func (controller SeriesController) LikeSeries(ctx *fiber.Ctx) error {
	hashId := ctx.Params("hashId")
	if hashId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid hashId",
			Data:    nil,
		})
	}

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userEmail := claims["email"].(string)
	userEntity := controller.UserService.GetUserByEmail(ctx.Context(), userEmail)

	series, err := controller.SeriesService.GetSeriesByHashId(ctx.Context(), hashId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = controller.SeriesService.LikeSeries(ctx.Context(), userEntity.Id, series.Id)
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
