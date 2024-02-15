package http_api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gpbPiazza/cargo-api/src/presentation/http_api/routes"
)

func setRoutes(app *fiber.App) {
	dispatcher := routes.NewDispatcher(app)
	routes.CustomerRoutes(dispatcher)
}
