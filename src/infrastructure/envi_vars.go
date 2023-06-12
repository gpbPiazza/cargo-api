package infrastructure

type EnvVars struct {
	AppPort string `env:"PORT"`
}

func NewEnvVars() EnvVars {
	var envs EnvVars

	envs.AppPort = "8080"

	return envs
}
