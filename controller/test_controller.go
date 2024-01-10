package controller

import (
	"fmt"
	"github.com/dokjasijeom/backend/configuration"
	"github.com/dokjasijeom/backend/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func NewTestController(config configuration.Config) *TestController {
	return &TestController{Config: config}
}

type TestController struct {
	configuration.Config
}

func (controller TestController) Route(app *fiber.App) {
	app.Get("/test", middleware.AuthenticateJWT("ANY", controller.Config), controller.restricted)
	app.Get("/ping", controller.ping)
}

func (controller TestController) restricted(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	return ctx.SendString(fmt.Sprintf("Welcome %s!", email))
}

func (controller TestController) ping(ctx *fiber.Ctx) error {
	return ctx.SendString("pong")
}
