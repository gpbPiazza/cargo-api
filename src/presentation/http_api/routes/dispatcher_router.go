package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const (
	internalPrefix = "private"
	clientPrefix   = "api"
)

type Dispatcher struct {
	InternalRoute fiber.Router
	ClientRoute   fiber.Router
}

func NewDispatcher(app *fiber.App) Dispatcher {
	internalAPI := app.Group(fmt.Sprintf("/%s/v1", internalPrefix))
	clientAPI := app.Group(fmt.Sprintf("/%s/v1", clientPrefix))

	return Dispatcher{
		InternalRoute: internalAPI,
		ClientRoute:   clientAPI,
	}
}
