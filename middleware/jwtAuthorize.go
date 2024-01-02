package middleware

import (
	"github.com/dokjasijeom/backend/configuration"
	"github.com/dokjasijeom/backend/model"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthenticateJWT(role string, config configuration.Config) func(*fiber.Ctx) error {
	jwtSecret := config.Get("JWT_SECRET_KEY")
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(jwtSecret)},
		SuccessHandler: func(ctx *fiber.Ctx) error {
			user := ctx.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			roles := claims["roles"].([]interface{})

			for _, roleInterface := range roles {
				if role == "ANY" {
					return ctx.Next()
				}
				if roleInterface == role {
					return ctx.Next()
				}
			}

			return ctx.Status(fiber.StatusUnauthorized).JSON(model.GeneralResponse{Code: 401, Message: "Unauthorized", Data: "Invalid Role"})
		},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{Code: 400, Message: "Bad Request", Data: "Missing or malformed JWT"})
			} else {
				return ctx.Status(fiber.StatusUnauthorized).JSON(model.GeneralResponse{Code: 401, Message: "Unauthorized", Data: "Invalid or expired JWT"})
			}
		},
	})
}
