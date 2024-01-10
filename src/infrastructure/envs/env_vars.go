package envs

import (
	"log"

	"github.com/caarlos0/env/v9"
)

type envVars struct {
	AppName    string `env:"APP_NAME" envDefault:"cargo-api"`
	ServerPort string `env:"SERVER_PORT,notEmpty"`

	Database dataBaseEnv
}

var globalEnvs *envVars

type dataBaseEnv struct {
	Name string `env:"DATABASE_NAME,notEmpty"`
}

// New envs follow the singleton pattern when New is called verifies if has a global instance and return the global if has one, if not creates a new one.
func New() *envVars {
	if globalEnvs != nil {
		return globalEnvs
	}

	envs := new(envVars)
	if err := env.Parse(envs); err != nil {
		log.Panicf("Could not load envconfig: %s.", err)
	}

	globalEnvs = envs

	return globalEnvs
}
