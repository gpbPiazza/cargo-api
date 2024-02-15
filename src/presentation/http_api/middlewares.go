package http_api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setMiddlewares(fiber *fiber.App) {
	fiber.Use(logger.New())
}
