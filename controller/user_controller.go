package controller

import (
	"github.com/dokjasijeom/backend/common"
	"github.com/dokjasijeom/backend/configuration"
	"github.com/dokjasijeom/backend/model"
	"github.com/dokjasijeom/backend/service"
	"github.com/gofiber/fiber/v2"
)

func NewUserController(userService *service.UserService, config configuration.Config) *UserController {
	return &UserController{UserService: *userService, Config: config}
}

type UserController struct {
	service.UserService
	configuration.Config
}

func (controller UserController) Route(app *fiber.App) {
	app.Post("/user/auth", controller.AuthenticateUser)
	app.Post("/users", controller.CreateUser)
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
	var userRoles []map[string]interface{}
	for _, userRole := range result.UserRoles {
		userRoles = append(userRoles, map[string]interface{}{
			"role": userRole.Roles,
		})
	}
	tokenJwtResult := common.GenerateToken(result.Email, userRoles, controller.Config)
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

	err = controller.UserService.CreateUser(ctx.Context(), request.Email, request.Password, request.ComparePassword)
	if err != nil {
		return err
	}
	result, _ := controller.UserService.AuthenticateUser(ctx.Context(), request.Email, request.Password)
	var userRoles []map[string]interface{}
	for _, userRole := range result.UserRoles {
		userRoles = append(userRoles, map[string]interface{}{
			"role": userRole.Roles,
		})
	}
	tokenJwtResult := common.GenerateToken(result.Email, userRoles, controller.Config)
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
