package presentation

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setUpMiddlewares(fiber *fiber.App) {
	fiber.Use(logger.New())
}
