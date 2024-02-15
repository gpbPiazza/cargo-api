package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	router := app.Group("/api/v1")

	setCargoRoutes(router)
}
