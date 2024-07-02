package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/dokjasijeom/backend/common"
	"github.com/dokjasijeom/backend/configuration"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/exception"
	"github.com/dokjasijeom/backend/middleware"
	"github.com/dokjasijeom/backend/model"
	"github.com/dokjasijeom/backend/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/image/webp"
	"image"
	_ "image/gif"
	"image/jpeg"
	"image/png"
	_ "image/png"
	"io"
	"log"
	"net/smtp"
	"os"
	"strings"
)

func NewUserController(userService *service.UserService, seriesService *service.SeriesService, episodeService *service.EpisodeService, providerService *service.ProviderService, userRecordSeriesService *service.UserRecordSeriesService, userRecordSeriesEpisodeService *service.UserRecordSeriesEpisodeService, config configuration.Config) *UserController {
	return &UserController{UserService: *userService, SeriesService: *seriesService, EpisodeService: *episodeService, ProviderService: *providerService, UserRecordSeriesService: *userRecordSeriesService, UserRecordSeriesEpisodeService: *userRecordSeriesEpisodeService, Config: config}
}

type UserController struct {
	service.UserService
	service.SeriesService
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
	app.Patch("/user", middleware.AuthenticateJWT("ANY", controller.Config), controller.UpdateUser)
	app.Delete("/user", middleware.AuthenticateJWT("ANY", controller.Config), controller.DeleteUser)
	app.Post("/user/forgot", controller.ForgotPassword)
	app.Get("/user/reset-password", controller.ResetPassword)
	app.Patch("/user/provider", middleware.AuthenticateJWT("ANY", controller.Config), controller.UpdateUserProvider)
	app.Delete("/user/avatar", middleware.AuthenticateJWT("ANY", controller.Config), controller.DeleteUserAvatar)
	app.Get("/user/series/complete-records", middleware.AuthenticateJWT("ANY", controller.Config), controller.GetUserCompleteRecords)
	app.Post("/user/series/record", middleware.AuthenticateJWT("ANY", controller.Config), controller.CreateUserRecordSeriesEpisode)
	app.Delete("/user/series/record", middleware.AuthenticateJWT("ANY", controller.Config), controller.DeleteUserRecordSeriesEpisode)
	app.Patch("/user/series/record/:id", middleware.AuthenticateJWT("ANY", controller.Config), controller.UpdateUserRecordSeries)
	app.Get("/user/series/:id", middleware.AuthenticateJWT("ANY", controller.Config), controller.GetUserRecordSeries)

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
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
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
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
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
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    fiber.StatusInternalServerError,
			Message: createErr.Error(),
			Data:    nil,
		})
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
	result := controller.UserService.GetUserByEmailAndSeries(ctx.Context(), userEmail)
	return ctx.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "success",
		Data:    result,
	})
}

// Update user
// Path: PATCH /user
// @Description Update user
// @Summary Update user
// @Tags User
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body UserUpdateRequestModel true "Request Body"
// @Success 200 {object} model.GeneralResponse
// @Router /user [patch]
func (controller UserController) UpdateUser(ctx *fiber.Ctx) error {
	var request = model.UserUpdateRequestModel{}
	form, err := ctx.MultipartForm()

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userEmail := claims["email"].(string)

	userEntity := controller.UserService.GetUserByEmail(ctx.Context(), userEmail)
	if userEntity.Email == "" {
		return ctx.Status(fiber.StatusNotFound).JSON(model.GeneralResponse{
			Code:    fiber.StatusNotFound,
			Message: "user not found",
			Data:    nil,
		})
	}

	// update user avatar
	fileheader, err := ctx.FormFile("image")
	if fileheader != nil {
		file, err := fileheader.Open()
		if err != nil {
			exception.PanicLogging(err)
		}

		buffer, err := io.ReadAll(file)
		if err != nil {
			exception.PanicLogging(err)
		}

		ext := fileheader.Filename[strings.LastIndex(fileheader.Filename, "."):]

		filename, err := imageProcessing(ctx.Context(), ext, buffer, 50)
		if err != nil {
			exception.PanicLogging(err)
		}
		request.Avatar = filename

		if userEntity.Profile.Avatar != "" {
			err = removeImage(ctx.Context(), userEntity.Profile.Avatar)
			if err != nil {
				exception.PanicLogging(err)
			}
		}
	}

	// update user name
	if form.Value["username"] != nil {
		request.Username = form.Value["username"][0]
	}

	if form.Value["password"] != nil {
		request.Password = form.Value["password"][0]
	}

	if form.Value["passwordConfirm"] != nil {
		request.PasswordConfirm = form.Value["passwordConfirm"][0]
	}

	// update user
	result, err := controller.UserService.UpdateUserProfile(ctx.Context(), userEntity.Id, request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if request.Password != "" && request.PasswordConfirm != "" && request.Password == request.PasswordConfirm {
		_, err = controller.UserService.UpdateUserPassword(ctx.Context(), userEntity.Id, request)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
				Code:    fiber.StatusInternalServerError,
				Message: err.Error(),
				Data:    nil,
			})
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "success",
		Data:    result,
	})
}

// Delete user
func (controller UserController) DeleteUser(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userEmail := claims["email"].(string)

	userEntity := controller.UserService.GetUserByEmail(ctx.Context(), userEmail)
	if userEntity.Email == "" {
		return ctx.Status(fiber.StatusNotFound).JSON(model.GeneralResponse{
			Code:    fiber.StatusNotFound,
			Message: "user not found",
			Data:    nil,
		})
	}

	_, err := controller.UserService.DeleteUser(ctx.Context(), userEntity.Id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.Status(fiber.StatusNoContent).JSON(model.GeneralResponse{
		Code:    fiber.StatusNoContent,
		Message: "success",
		Data:    nil,
	})
}

// Update user provider
// Path: PATCH /user/provider
// @Description Update user provider
// @Summary Update user provider
// @Tags User
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body UserProviderRequestModel true "Request Body"
// @Success 200 {object} model.GeneralResponse
// @Router /user/provider [patch]
func (controller UserController) UpdateUserProvider(ctx *fiber.Ctx) error {
	var request struct {
		Providers []string `json:"providers"`
	}

	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userEmail := claims["email"].(string)

	userEntity := controller.UserService.GetUserByEmail(ctx.Context(), userEmail)
	if userEntity.Email == "" {
		return ctx.Status(fiber.StatusNotFound).JSON(model.GeneralResponse{
			Code:    fiber.StatusNotFound,
			Message: "user not found",
			Data:    nil,
		})
	}

	providerEntities, err := controller.ProviderService.GetProviderByHashIds(ctx.Context(), request.Providers)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	providerIds := make([]uint, 0)
	for _, provider := range providerEntities {
		providerIds = append(providerIds, provider.Id)
	}

	// update user provider
	result, err := controller.UserService.UpdateUserProviders(ctx.Context(), userEntity.Id, providerIds)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "success",
		Data:    result,
	})
}

// Delete user avatar
// Path: DELETE /user/avater
// @Description Delete user avatar
// @Summary Delete user avatar
// @Tags User
// @Accept json
// @Produce json
// @Security Bearer
// @Success 204 {object} model.GeneralResponse
// @Router /user/avater [delete]
func (controller UserController) DeleteUserAvatar(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userEmail := claims["email"].(string)

	userEntity := controller.UserService.GetUserByEmail(ctx.Context(), userEmail)
	if userEntity.Email == "" {
		return ctx.Status(fiber.StatusNotFound).JSON(model.GeneralResponse{
			Code:    fiber.StatusNotFound,
			Message: "user not found",
			Data:    nil,
		})
	}

	if userEntity.Profile.Avatar != "" {
		err := removeImage(ctx.Context(), userEntity.Profile.Avatar)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
				Code:    fiber.StatusInternalServerError,
				Message: err.Error(),
				Data:    nil,
			})
		}
		_, err = controller.UserService.DeleteUserAvatar(ctx.Context(), userEntity.Id)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
				Code:    fiber.StatusInternalServerError,
				Message: err.Error(),
				Data:    nil,
			})
		}
	}

	return ctx.Status(fiber.StatusNoContent).JSON(model.GeneralResponse{
		Code:    fiber.StatusNoContent,
		Message: "success",
		Data:    nil,
	})

}

func (controller UserController) GetUserCompleteRecords(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userEmail := claims["email"].(string)

	userEntity := controller.UserService.GetUserByEmail(ctx.Context(), userEmail)
	if userEntity.Email == "" {
		return ctx.Status(fiber.StatusNotFound).JSON(model.GeneralResponse{
			Code:    fiber.StatusNotFound,
			Message: "user not found",
			Data:    nil,
		})
	}

	result, err := controller.UserRecordSeriesService.GetUserCompleteRecords(ctx.Context(), userEntity.Id)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

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
	var request model.UserRecordSeriesEpisodeRequestModel
	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userEmail := claims["email"].(string)

	userEntity := controller.UserService.GetUserByEmail(ctx.Context(), userEmail)

	recordEntity, err := controller.UserRecordSeriesService.GetUserRecordSeriesByUserIdAndId(ctx.Context(), userEntity.Id, request.UserRecordSeriesId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	providerEntity, err := controller.ProviderService.GetProviderByName(ctx.Context(), request.ProviderName)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	isBulkCreate := false
	// request.From is not omitted than isBulkCreate is true
	if request.From != 0 {
		isBulkCreate = true
	}

	seriesEntity, err := controller.SeriesService.GetSeriesById(ctx.Context(), recordEntity.SeriesId)

	var seriesEpisodes []entity.Episode
	// get episodes by series id
	seriesEpisodes = seriesEntity.Episodes

	if isBulkCreate {
		// 내 서재에 등록한 작품에 다중 회차를 기록할 때
		var episodes []entity.UserRecordSeriesEpisode
		// create bulk user record series episode
		if recordEntity.SeriesId != 0 && recordEntity.Title != "" {
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
			for i := request.From; i <= request.To; i++ {
				episodes = append(episodes, entity.UserRecordSeriesEpisode{
					UserRecordSeriesId: request.UserRecordSeriesId,
					EpisodeId:          0,
					EpisodeNumber:      i,
					Watched:            true,
					ProviderId:         providerEntity.Id,
					ProviderName:       providerEntity.Name,
				})
			}

			result, _ := controller.UserRecordSeriesEpisodeService.CreateBulkUserRecordSeriesEpisode(ctx.Context(), episodes)
			return ctx.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
				Code:    fiber.StatusCreated,
				Message: "success",
				Data:    result,
			})
		}
	} else {
		// 내 서재에 등록한 작품에 단일 회차를 기록할 때
		if recordEntity.SeriesId != 0 && recordEntity.Title != "" {
			currentEpisode, ok := lo.Find(seriesEpisodes, func(episode entity.Episode) bool {
				return episode.EpisodeNumber == request.To
			})
			if ok {
				result, err := controller.UserRecordSeriesEpisodeService.CreateUserRecordSeriesEpisode(ctx.Context(), entity.UserRecordSeriesEpisode{
					UserRecordSeriesId: request.UserRecordSeriesId,
					EpisodeId:          currentEpisode.Id,
					EpisodeNumber:      currentEpisode.EpisodeNumber,
					Watched:            true,
					ProviderId:         providerEntity.Id,
					ProviderName:       providerEntity.Name,
				})

				if err != nil {
					return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
						Code:    fiber.StatusBadRequest,
						Message: err.Error(),
						Data:    nil,
					})
				}

				// result to array
				var results []entity.UserRecordSeriesEpisode
				results = append(results, result)

				return ctx.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
					Code:    fiber.StatusCreated,
					Message: "success",
					Data:    results,
				})
			}
		} else {
			result, err := controller.UserRecordSeriesEpisodeService.CreateUserRecordSeriesEpisode(ctx.Context(), entity.UserRecordSeriesEpisode{
				UserRecordSeriesId: request.UserRecordSeriesId,
				EpisodeId:          0,
				EpisodeNumber:      request.To,
				Watched:            true,
				ProviderId:         providerEntity.Id,
				ProviderName:       providerEntity.Name,
			})

			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
					Code:    fiber.StatusBadRequest,
					Message: err.Error(),
					Data:    nil,
				})
			}

			// result to array
			var results []entity.UserRecordSeriesEpisode
			results = append(results, result)

			return ctx.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
				Code:    fiber.StatusCreated,
				Message: "success",
				Data:    results,
			})
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "success",
		Data:    nil,
	})
}

// Delete user record series episode
// Path: DELETE /user/series/record
// @Description Delete user record series episode
// @Summary Delete user record series episode
// @Tags User
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body UserRecordSeriesEpisodeDeleteRequestModel true "Request Body"
// @Success 204 {object} model.GeneralResponse
// @Router /user/series/record [delete]
func (controller UserController) DeleteUserRecordSeriesEpisode(ctx *fiber.Ctx) error {
	var request model.UserRecordSeriesEpisodeDeleteRequestModel
	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if request.UserRecordSeriesId == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "user record series id is required",
			Data:    nil,
		})
	}

	if len(request.RecordIds) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "record ids is required",
			Data:    nil,
		})
	}

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userEmail := claims["email"].(string)

	userEntity := controller.UserService.GetUserByEmail(ctx.Context(), userEmail)
	recordEntity, err := controller.UserRecordSeriesService.GetUserRecordSeriesByUserIdAndId(ctx.Context(), userEntity.Id, request.UserRecordSeriesId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if recordEntity.Id == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(model.GeneralResponse{
			Code:    fiber.StatusNotFound,
			Message: "record not found",
			Data:    nil,
		})
	}

	// delete user record series episode
	err = controller.UserRecordSeriesEpisodeService.DeleteUserRecordSeriesEpisodeByUserRecordSeriesEpisodeIds(ctx.Context(), request.RecordIds)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.Status(fiber.StatusNoContent).JSON(model.GeneralResponse{
		Code:    fiber.StatusNoContent,
		Message: "success",
		Data:    nil,
	})
}

func (controller UserController) UpdateUserRecordSeries(ctx *fiber.Ctx) error {
	var request model.UserRecordSeriesUpdateModel

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

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userEmail := claims["email"].(string)

	userEntity := controller.UserService.GetUserByEmail(ctx.Context(), userEmail)
	if userEntity.Email == "" {
		return ctx.Status(fiber.StatusNotFound).JSON(model.GeneralResponse{
			Code:    fiber.StatusNotFound,
			Message: "user not found",
			Data:    nil,
		})
	}

	recordEntity, err := controller.UserRecordSeriesService.GetUserRecordSeriesByUserIdAndId(ctx.Context(), userEntity.Id, uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if recordEntity.Id == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(model.GeneralResponse{
			Code:    fiber.StatusNotFound,
			Message: "record not found",
			Data:    nil,
		})
	}

	if request.Title != "" {
		recordEntity.Title = request.Title
	}
	if request.Author != "" {
		recordEntity.Author = request.Author
	}
	if request.Genre != "" {
		recordEntity.Genre = request.Genre
	}
	if request.SeriesType != "" {
		recordEntity.SeriesType = request.SeriesType
	}
	recordEntity.ReadCompleted = request.ReadCompleted

	// update user record series
	result, err := controller.UserRecordSeriesService.UpdateUserRecordSeries(ctx.Context(), userEntity.Id, recordEntity.Id, recordEntity)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "success",
		Data:    result,
	})
}

// Forgot password
func (controller UserController) ForgotPassword(ctx *fiber.Ctx) error {
	var request struct {
		Email string `json:"email"`
	}
	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	userEntity := controller.UserService.GetUserByEmail(ctx.Context(), request.Email)
	if userEntity.Email == "" {
		return ctx.Status(fiber.StatusNotFound).JSON(model.GeneralResponse{
			Code:    fiber.StatusNotFound,
			Message: "user not found",
			Data:    nil,
		})
	}

	// make token
	token, err := controller.UserService.MakePasswordResetToken(ctx.Context(), request.Email)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	mailtrap_username := os.Getenv("MAILTRAP_USERNAME")
	mailtrap_password := os.Getenv("MAILTRAP_PASSWORD")
	mailtrap_host := os.Getenv("MAILTRAP_HOST")

	mailtrap_auth := smtp.PlainAuth("", mailtrap_username, mailtrap_password, mailtrap_host)

	// mailtrap from
	from := os.Getenv("MAILTRAP_FROM")
	// mailtrap to
	to := []string{request.Email}

	message := []byte("To: " + request.Email + "\r\n" +
		"From: " + from + "\r\n" +
		"Subject: [독자시점] 로그인 매직 링크\r\n" +
		"\r\n" +
		"다음의 링크를 주소창에 붙여넣기 하여 로그인을 완료한 후 내 정보 페이지에서 비밀번호를 수정해주세요.: \r\n\r\n" + os.Getenv("BACKEND_HOST") + "/user/reset-password?token=" + token + "\r\n\r\n" + "감사합니다.\r\n")

	smtpUrl := mailtrap_host + ":587"
	err = smtp.SendMail(smtpUrl, mailtrap_auth, from, to, message)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "success",
		Data:    nil,
	})
}

// Reset Password
func (controller UserController) ResetPassword(ctx *fiber.Ctx) error {
	token := ctx.Query("token")
	if token == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "token is required",
			Data:    nil,
		})
	}

	// get email
	email, err := controller.UserService.GetPasswordResetToken(ctx.Context(), token)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// delete token
	err = controller.UserService.DeletePasswordResetToken(ctx.Context(), token)

	userEntity := controller.UserService.GetUserByEmail(ctx.Context(), email)
	if userEntity.Email == "" {
		return ctx.Status(fiber.StatusNotFound).JSON(model.GeneralResponse{
			Code:    fiber.StatusNotFound,
			Message: "user not found",
			Data:    nil,
		})
	}

	result, _ := controller.UserService.AuthenticateOnlyEmail(ctx.Context(), email)
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
		"code":    200,
		"message": "success",
		"data": map[string]interface{}{
			"token": tokenJwtResult,
			"email": result.Email,
			"role":  userRoles,
		},
	}

	out, err := json.Marshal(removeNils(resultWithToken))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	ctx.Cookie(&fiber.Cookie{Name: "DS_AUT", Value: tokenJwtResult, Path: "/", Domain: os.Getenv("COOKIE_DOMAIN")})
	ctx.Cookie(&fiber.Cookie{Name: "DS_USER", Value: string(out), Path: "/", Domain: os.Getenv("COOKIE_DOMAIN")})
	ctx.Cookie(&fiber.Cookie{Name: "isForgotPassword", Value: "true", Path: "/", Domain: os.Getenv("COOKIE_DOMAIN")})

	// redirect main page
	return ctx.Redirect(os.Getenv("FRONTEND_HOST") + "/my/profile")
}

// Get user record series
// Path: GET /user/series/:id
// @Description Get user record series
// @Summary Get user record series
// @Tags User
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "User Record Series Id"
// @Success 200 {object} model.GeneralResponse
// @Router /user/series/{id} [get]
func (controller UserController) GetUserRecordSeries(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", 0)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "id is required",
			Data:    nil,
		})
	}

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userEmail := claims["email"].(string)

	userEntity := controller.UserService.GetUserByEmail(ctx.Context(), userEmail)
	recordEntity, err := controller.UserRecordSeriesService.GetUserRecordSeriesByUserIdAndId(ctx.Context(), userEntity.Id, uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if recordEntity.Id == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(model.GeneralResponse{
			Code:    fiber.StatusNotFound,
			Message: "record not found",
			Data:    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "success",
		Data:    recordEntity,
	})
}

func imageProcessing(ctx context.Context, fileExt string, buffer []byte, quality int) (string, error) {
	filename := strings.Replace(uuid.New().String(), "-", "", -1)
	newFileExt := ".jpg"
	var img image.Image
	var err error

	if fileExt == ".webp" {
		img, err = webp.Decode(bytes.NewReader(buffer))
	} else if fileExt == ".png" {
		img, err = png.Decode(bytes.NewReader(buffer))
		newFileExt = ".png"
	} else {
		img, _, err = image.Decode(bytes.NewReader(buffer))
	}

	if err != nil {
		exception.PanicLogging(err)
	}
	var jpegBuffer bytes.Buffer
	if fileExt == ".png" {
		err = png.Encode(&jpegBuffer, img)
	} else {
		err = jpeg.Encode(&jpegBuffer, img, &jpeg.Options{Quality: quality})
	}
	if err != nil {
		exception.PanicLogging(err)
	}

	f, err := os.CreateTemp("./tmp/", "tempfile-")
	if err != nil {
		exception.PanicLogging(err)
	}

	if _, err := f.Write(jpegBuffer.Bytes()); err != nil {
		exception.PanicLogging(err)
	}

	cld, err := configuration.NewCloudinaryConfigruation()
	if err != nil {
		exception.PanicLogging(err)
	}

	filename = filename + newFileExt

	filenameWithoutExt := strings.TrimSuffix(filename, newFileExt)

	uploadResult, err := cld.Upload.Upload(ctx, f.Name(), uploader.UploadParams{
		ResourceType: "image",
		PublicID:     filenameWithoutExt,
		Folder:       "avatar",
	})
	if err != nil {
		exception.PanicLogging(err)
	}

	defer f.Close()
	defer os.Remove(f.Name())

	log.Println(uploadResult)

	return "avatar/" + filename, err
}

func removeImage(ctx context.Context, filePath string) error {
	cld, err := configuration.NewCloudinaryConfigruation()
	if err != nil {
		exception.PanicLogging(err)
	}

	filename := strings.TrimPrefix(filePath, "avatar/")

	_, err = cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: filename,
	})
	if err != nil {
		exception.PanicLogging(err)
	}

	return nil
}

func removeNils(initialMap map[string]interface{}) map[string]interface{} {
	withoutNils := map[string]interface{}{}
	for key, value := range initialMap {
		_, ok := value.(map[string]interface{})
		if ok {
			value = removeNils(value.(map[string]interface{}))
			withoutNils[key] = value
			continue
		}
		if value != nil {
			withoutNils[key] = value
		}
	}
	return withoutNils
}
