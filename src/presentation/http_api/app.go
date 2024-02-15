package http_api

import (
	"github.com/gofiber/fiber/v2"
)

func NewApp() *fiber.App {
	app := fiber.New()
	setMiddlewares(app)
	setRoutes(app)
	return app
}
