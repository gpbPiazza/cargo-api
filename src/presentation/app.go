package presentation

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gpbPiazza/cargo-api/src/infrastructure"
	"github.com/gpbPiazza/cargo-api/src/presentation/routes"
)

func NewApp() app {
	fiberApp := fiber.New()

	setUpMiddlewares(fiberApp)
	routes.SetUp(fiberApp)

	envVars := infrastructure.NewEnvVars()

	return app{
		fiber: fiberApp,
		envs:  envVars,
	}
}

type app struct {
	fiber *fiber.App
	envs  infrastructure.EnvVars
}

func (a app) Start() {
	a.fiber.Listen(fmt.Sprintf(":%s", a.envs.AppPort))
}
