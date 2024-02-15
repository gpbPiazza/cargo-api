package http_api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/gpbPiazza/cargo-api/src/infrastructure/envs"
	"github.com/gpbPiazza/cargo-api/src/presentation/http_api/routes"
)

func NewApp() app {
	fiberApp := fiber.New()

	setUpMiddlewares(fiberApp)
	routes.SetUp(fiberApp)

	return app{
		fiber: fiberApp,
	}
}

type app struct {
	fiber *fiber.App
}

func (a app) Start() {
	envsVar := envs.New()
	a.fiber.Listen(fmt.Sprintf(":%s", envsVar.ServerPort))
}
