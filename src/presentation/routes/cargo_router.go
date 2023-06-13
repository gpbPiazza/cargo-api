package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gpbPiazza/cargo-api/src/application"
)

func setCargoRoutes(api fiber.Router) {
	api.Post("/cargos/:id/event-logging", adaptRoute(application.NewEventLoggingController()))

	api.Post("/cargos", adaptRoute(application.NewBookingController()))
}
