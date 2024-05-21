package controller

import (
	"github.com/dokjasijeom/backend/common"
	"github.com/dokjasijeom/backend/configuration"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/middleware"
	"github.com/dokjasijeom/backend/model"
	"github.com/dokjasijeom/backend/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/samber/lo"
)

func NewUserController(userService *service.UserService, episodeService *service.EpisodeService, providerService *service.ProviderService, userRecordSeriesService *service.UserRecordSeriesService, userRecordSeriesEpisodeService *service.UserRecordSeriesEpisodeService, config configuration.Config) *UserController {
	return &UserController{UserService: *userService, EpisodeService: *episodeService, ProviderService: *providerService, UserRecordSeriesService: *userRecordSeriesService, UserRecordSeriesEpisodeService: *userRecordSeriesEpisodeService, Config: config}
}

type UserController struct {
	service.UserService
	service.EpisodeService
	service.ProviderService
	service.UserRecordSeriesService
	service.UserRecordSeriesEpisodeService
	configuration.Config
}

func (controller UserController) Route(app *fiber.App) {
	app.Post("/user/auth", controller.AuthenticateUser)
	app.Post("/users", controller.CreateUser)
	app.Get("/user", middleware.AuthenticateJWT("ANY", controller.Config), controller.GetUser)
	app.Post("/user/series/record", middleware.AuthenticateJWT("ANY", controller.Config), controller.CreateUserRecordSeriesEpisode)
}

// Authenticate user
// Path: POST /user/auth
// @Description Authenticate user
// @Summary Authenticate user
// @Tags User
// @Accept json
// @Produce json
// @Param request body Email Password true "Request Body"
// @Success 200 {object} model.GeneralResponse
// @Router /user/auth [post]
func (controller UserController) AuthenticateUser(ctx *fiber.Ctx) error {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := ctx.BodyParser(&request)
	if err != nil {
		return err
	}

	result, _ := controller.UserService.AuthenticateUser(ctx.Context(), request.Email, request.Password)
	//var userRoles []map[string]interface{}
	//for _, userRole := range result.UserRoles {
	//	userRoles = append(userRoles, map[string]interface{}{
	//		"role": userRole,
	//	})
	//}
	// spread result.Roles key Role to array

	var userRoles []string
	for _, role := range result.Roles {
		userRoles = append(userRoles, role.Role)
	}

	tokenJwtResult := common.GenerateToken(result.Email, result.Roles, controller.Config)
	resultWithToken := map[string]interface{}{
		"token": tokenJwtResult,
		"email": result.Email,
		"role":  userRoles,
	}
	return ctx.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "success",
		Data:    resultWithToken,
	})
}

// Create user
// Path: POST /users
// @Description Create user
// @Summary Create user
// @Tags User
// @Accept json
// @Produce json
// @Param request body Email Password ComparePassword true "Request Body"
// @Success 201 {object} model.GeneralResponse
// @Router /users [post]
func (controller UserController) CreateUser(ctx *fiber.Ctx) error {
	var request struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ComparePassword string `json:"compare_password"`
	}
	err := ctx.BodyParser(&request)
	if err != nil {
		return err
	}

	existUser := controller.UserService.GetUserByEmail(ctx.Context(), request.Email)
	if existUser.Email == request.Email {
		return ctx.Status(fiber.StatusConflict).JSON(model.GeneralResponse{
			Code:    fiber.StatusConflict,
			Message: "user is already exist",
			Data:    nil,
		})
	}

	_, createErr := controller.UserService.CreateUser(ctx.Context(), request.Email, request.Password, request.ComparePassword)
	if createErr != nil {
		return createErr
	}
	result, _ := controller.UserService.AuthenticateUser(ctx.Context(), request.Email, request.Password)

	var userRoles []string
	for _, role := range result.Roles {
		userRoles = append(userRoles, role.Role)
	}

	tokenJwtResult := common.GenerateToken(result.Email, result.Roles, controller.Config)
	resultWithToken := map[string]interface{}{
		"token": tokenJwtResult,
		"email": result.Email,
		"role":  userRoles,
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    fiber.StatusCreated,
		Message: "success",
		Data:    resultWithToken,
	})
}

// Get user
// Path: GET /user
// @Description Get user
// @Summary Get user
// @Tags User
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} model.GeneralResponse
// @Router /user [get]
func (controller UserController) GetUser(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userEmail := claims["email"].(string)
	result := controller.UserService.GetUserByEmail(ctx.Context(), userEmail)
	return ctx.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "success",
		Data:    result,
	})
}

// Create user record series episode
// Path: POST /user/series/record
// @Description Create user record series episode
// @Summary Create user record series episode
// @Tags User
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body UserRecordSeriesEpisodeRequestModel true "Request Body"
// @Success 201 {object} model.GeneralResponse
// @Router /user/series/record [post]
func (controller UserController) CreateUserRecordSeriesEpisode(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userEmail := claims["email"].(string)

	userEntity := controller.UserService.GetUserByEmail(ctx.Context(), userEmail)

	var request model.UserRecordSeriesEpisodeRequestModel
	err := ctx.BodyParser(&request)
	if err != nil {
		return err
	}

	recordEntity, err := controller.UserRecordSeriesService.GetUserRecordSeriesByUserIdAndId(ctx.Context(), userEntity.Id, request.UserRecordSeriesId)
	if err != nil {
		return err
	}

	providerEntity, err := controller.ProviderService.GetProviderByName(ctx.Context(), request.ProviderName)
	if err != nil {
		return err
	}

	isBulkCreate := false
	// request.From is not omitted than isBulkCreate is true
	if request.From != 0 {
		isBulkCreate = true
	}

	var seriesEpisodes []entity.Episode
	// get episodes by series id
	seriesEpisodes, _ = controller.EpisodeService.GetEpisodesBySeriesId(ctx.Context(), recordEntity.SeriesId)

	if isBulkCreate {
		// 내 서재에 등록한 작품에 다중 회차를 기록할 때
		var episodes []entity.UserRecordSeriesEpisode
		// create bulk user record series episode
		for i := request.From; i <= request.To; i++ {
			currentEpisode, ok := lo.Find(seriesEpisodes, func(episode entity.Episode) bool {
				return episode.EpisodeNumber == i
			})
			if ok {
				episodes = append(episodes, entity.UserRecordSeriesEpisode{
					UserRecordSeriesId: request.UserRecordSeriesId,
					EpisodeId:          currentEpisode.Id,
					EpisodeNumber:      currentEpisode.EpisodeNumber,
					Watched:            true,
					ProviderId:         providerEntity.Id,
					ProviderName:       providerEntity.Name,
				})
			}
		}

		result, _ := controller.UserRecordSeriesEpisodeService.CreateBulkUserRecordSeriesEpisode(ctx.Context(), episodes)
		return ctx.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
			Code:    fiber.StatusCreated,
			Message: "success",
			Data:    result,
		})
	} else {
		// 내 서재에 등록한 작품에 단일 회차를 기록할 때
		currentEpisode, ok := lo.Find(seriesEpisodes, func(episode entity.Episode) bool {
			return episode.EpisodeNumber == request.To
		})
		if ok {
			result, _ := controller.UserRecordSeriesEpisodeService.CreateUserRecordSeriesEpisode(ctx.Context(), entity.UserRecordSeriesEpisode{
				UserRecordSeriesId: request.UserRecordSeriesId,
				EpisodeId:          currentEpisode.Id,
				EpisodeNumber:      currentEpisode.EpisodeNumber,
				Watched:            true,
				ProviderId:         providerEntity.Id,
				ProviderName:       providerEntity.Name,
			})
			return ctx.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
				Code:    fiber.StatusCreated,
				Message: "success",
				Data:    result,
			})
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "success",
		Data:    nil,
	})
}
