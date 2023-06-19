package env

import (
	"log"

	"github.com/caarlos0/env"
)

type EnvVars struct {
	AppName    string `env:"APP_NAME" envDefault:"cargo-api"`
	ServerPort string `env:"SERVER_PORT"`

	Database DataBaseEnv
}

var globalEnvs EnvVars

type DataBaseEnv struct {
	Name string `env:"DATABASE_NAME"`
}

func Init() {
	var envs EnvVars

	if err := env.Parse(&envs); err != nil {
		log.Fatal(err)
	}

	globalEnvs = envs
}

func Envs() EnvVars {
	return globalEnvs
}

func Validate() error {
	return nil
}
