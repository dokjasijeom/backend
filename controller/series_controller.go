package controller

import (
	"github.com/dokjasijeom/backend/configuration"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/middleware"
	"github.com/dokjasijeom/backend/model"
	"github.com/dokjasijeom/backend/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func NewSeriesController(seriesService *service.SeriesService, seriesDailyViewService *service.SeriesDailyViewService, userService *service.UserService, userRecordSeriesService *service.UserRecordSeriesService, userRecordSeriesEpisodeService *service.UserRecordSeriesEpisodeService, config configuration.Config) *SeriesController {
	return &SeriesController{SeriesService: *seriesService, SeriesDailyViewService: *seriesDailyViewService, UserService: *userService, UserRecordSeriesService: *userRecordSeriesService, UserRecordSeriesEpisodeService: *userRecordSeriesEpisodeService, Config: config}
}

type SeriesController struct {
	service.SeriesService
	service.SeriesDailyViewService
	service.UserService
	service.UserRecordSeriesService
	service.UserRecordSeriesEpisodeService
	configuration.Config
}

func (controller SeriesController) Route(app fiber.Router) {
	series := app.Group("/series")
	series.Get("/", controller.GetAllSeries)
	series.Get("/new", controller.GetNewEpisodeUpdateProviderSeries)
	series.Get("/:hashId", controller.GetSeriesByHashId)
	series.Post("/:hashId/like", middleware.AuthenticateJWT("ANY", controller.Config), controller.LikeSeries)
	series.Delete("/:hashId/like", middleware.AuthenticateJWT("ANY", controller.Config), controller.UnlikeSeries)
	series.Post("/:hashId/record", middleware.AuthenticateJWT("ANY", controller.Config), controller.CreateUserRecordSeries)
	series.Delete("/:hashId/record", middleware.AuthenticateJWT("ANY", controller.Config), controller.DeleteUserRecordSeries)
	series.Post("/non-exist/record", middleware.AuthenticateJWT("ANY", controller.Config), controller.CreateUserRecordEmptySeries)
	series.Delete("/non-exist/record", middleware.AuthenticateJWT("ANY", controller.Config), controller.DeleteUserRecordEmptySeries)
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

func (controller SeriesController) GetNewEpisodeUpdateProviderSeries(ctx *fiber.Ctx) error {
	provider := ctx.Query("provider")
	seriesType := ctx.Query("seriesType")

	// validate provider
	if provider == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid provider",
			Data:    nil,
		})
	}

	// validate seriesType
	if seriesType == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid seriesType",
			Data:    nil,
		})
	}

	result, err := controller.SeriesService.GetNewEpisodeUpdateProviderSeries(ctx.Context(), provider, seriesType)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
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

	now := time.Now()
	currentDate, _ := time.Parse("2006-01-02", now.Format("2006-01-02"))
	err = controller.SeriesDailyViewService.UpsertSeriesDailyView(ctx.Context(), result.Id, currentDate)

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

	hasLike, err := controller.SeriesService.HasLikeSeries(ctx.Context(), userEntity.Id, series.Id)
	if hasLike == true {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "User has already liked this series",
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

// unlike series
func (controller SeriesController) UnlikeSeries(ctx *fiber.Ctx) error {
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

	hasLike, err := controller.SeriesService.HasLikeSeries(ctx.Context(), userEntity.Id, series.Id)
	if hasLike == false {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "User has not liked this series",
			Data:    nil,
		})
	}

	err = controller.SeriesService.UnlikeSeries(ctx.Context(), userEntity.Id, series.Id)
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

// create user record series
func (controller SeriesController) CreateUserRecordSeries(ctx *fiber.Ctx) error {
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

	userRecordSeries := entity.UserRecordSeries{
		UserId:       userEntity.Id,
		SeriesId:     series.Id,
		TotalEpisode: series.TotalEpisode,
	}

	record, err := controller.UserRecordSeriesService.CreateUserRecordSeries(ctx.Context(), userRecordSeries)
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
		Data:    record,
	})
}

// delete user record series
func (controller SeriesController) DeleteUserRecordSeries(ctx *fiber.Ctx) error {
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

	recordEntity, err := controller.UserRecordSeriesService.GetUserRecordSeriesByUserIdAndSeriesId(ctx.Context(), userEntity.Id, series.Id)

	err = controller.UserRecordSeriesService.DeleteUserRecordSeriesByUserIdAndSeriesId(ctx.Context(), userEntity.Id, series.Id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}
	err = controller.UserRecordSeriesEpisodeService.DeleteUserRecordSeriesEpisodeByUserRecordSeriesId(ctx.Context(), recordEntity.Id)
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

// create user record empty series
func (controller SeriesController) CreateUserRecordEmptySeries(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userEmail := claims["email"].(string)
	userEntity := controller.UserService.GetUserByEmail(ctx.Context(), userEmail)

	var request = model.UserRecordSeriesEmptyModel{}
	err := ctx.BodyParser(&request)

	userRecordSeries := entity.UserRecordSeries{
		UserId:       userEntity.Id,
		SeriesId:     0,
		Title:        request.Title,
		Author:       request.Author,
		Genre:        request.Genre,
		TotalEpisode: request.TotalEpisode,
	}

	record, err := controller.UserRecordSeriesService.CreateUserRecordSeries(ctx.Context(), userRecordSeries)
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
		Data:    record,
	})
}

// delete user record empty series
func (controller SeriesController) DeleteUserRecordEmptySeries(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userEmail := claims["email"].(string)
	userEntity := controller.UserService.GetUserByEmail(ctx.Context(), userEmail)

	var request = model.UserRecordSeriesEmptyModel{}
	err := ctx.BodyParser(&request)

	if request.Id == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid id",
			Data:    nil,
		})
	}

	recordEntity, err := controller.UserRecordSeriesService.GetUserRecordSeriesByUserIdAndSeriesId(ctx.Context(), userEntity.Id, request.Id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = controller.UserRecordSeriesService.DeleteUserRecordSeriesByUserIdAndId(ctx.Context(), userEntity.Id, request.Id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}
	err = controller.UserRecordSeriesEpisodeService.DeleteUserRecordSeriesEpisodeByUserRecordSeriesId(ctx.Context(), recordEntity.Id)
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
