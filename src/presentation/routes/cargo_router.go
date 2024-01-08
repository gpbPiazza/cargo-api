package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gpbPiazza/cargo-api/src/presentation/handlers/cargo_handlers"
)

func setCargoRoutes(api fiber.Router) {
	api.Post("/cargos", cargo_handlers.Post)
}