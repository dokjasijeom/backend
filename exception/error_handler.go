package exception

import (
	"github.com/dokjasijeom/backend/model"
	"github.com/gofiber/fiber/v2"
	"log"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	log.Println(err)
	return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
		Code:    fiber.StatusInternalServerError,
		Message: "General Error",
		Data:    err.Error(),
	})
}
