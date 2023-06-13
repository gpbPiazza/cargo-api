package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gpbPiazza/cargo-api/src/application"
)

func adaptRoute(controller application.Controller) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		request := application.HttpRequest{
			Ctx: c.UserContext(),
			// Body: c.BodyParser,
		}

		response := controller.Control(request)

		return c.SendStatus(response.StatusCode)
	}
}
