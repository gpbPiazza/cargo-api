package presentation

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gpbPiazza/cargo-api/src/application"
)

func adaptRoute(controller application.Controller) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// request := http.HttpRequest{
		// 	Ctx:        c.UserContext(),
		// 	Body:       c.Body(),
		// 	Params:     c.AllParams(),
		// 	FiberCtx:   c,
		// 	PathParam:  c.Params("id", uuid.Nil.String()),
		// 	Pagination: pagination,
		// }

		// response := handler.Handle(request)

		return c.SendStatus(http.StatusOK)
	}
}

func setUpEventLoggingApp(api fiber.Router) {
	api.Post("/event-logging", adaptRoute(application.NewEventLoggingController()))
}

func setUpRoutes(app *fiber.App) {
	router := app.Group("/api/v1")

	setUpEventLoggingApp(router)
}
