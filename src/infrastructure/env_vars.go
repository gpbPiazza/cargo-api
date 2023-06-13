package infrastructure

import "os"

type EnvVars struct {
	SERVER_PORT string
}

func NewEnvVars() EnvVars {
	var envs EnvVars

	os.Setenv("SERVER_PORT", "8080")
	envs.SERVER_PORT = "8080"

	return envs
}
