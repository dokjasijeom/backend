package configuration

import (
	"github.com/dokjasijeom/backend/exception"
	"github.com/gofiber/fiber/v2"
)

func NewFiberConfiguration() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
		BodyLimit:    30 * 1024 * 1024,
	}
}
