package presentation

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/gpbPiazza/cargo-api/src/infrastructure/env"
	"github.com/gpbPiazza/cargo-api/src/presentation/routes"
)

func NewApp() app {
	fiberApp := fiber.New()
	env.Init()

	setUpMiddlewares(fiberApp)
	routes.SetUp(fiberApp)

	return app{
		fiber: fiberApp,
		envs:  env.Envs(),
	}
}

type app struct {
	fiber *fiber.App
	envs  env.EnvVars
}

func (a app) Start() {
	a.fiber.Listen(fmt.Sprintf(":%s", a.envs.ServerPort))
}
