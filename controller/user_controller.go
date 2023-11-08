package controller

import (
	"github.com/dokjasijeom/backend/model"
	"github.com/dokjasijeom/backend/service"
	"github.com/gofiber/fiber/v2"
)

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{UserService: *userService}
}

type UserController struct {
	service.UserService
}

func (controller UserController) Route(app *fiber.App) {
	app.Post("/user/auth", controller.AuthenticateUser)
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
	return ctx.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "success",
		Data:    result,
	})
}
