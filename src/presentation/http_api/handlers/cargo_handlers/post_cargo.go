package cargo_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gpbPiazza/cargo-api/src/presentation/dtos"
)

// Legal agora eu preciso do aplicativo que faz a criação de uma carga
func Post(c *fiber.Ctx) error {

	body := new(dtos.PostCargoRequest)
	if err := c.BodyParser(&body); err != nil {
		// TODO BUILD err Handler
		return err
	}

	return nil
}
