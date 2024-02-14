package envs

type dataBaseEnv struct {
	Name         string `env:"DATABASE_NAME,notEmpty"`
	Username     string `env:"DATABASE_USERNAME,notEmpty"`
	Password     string `env:"DATABASE_PASSWORD,notEmpty"`
	Host         string `env:"DATABASE_HOST,notEmpty"`
	Port         string `env:"DATABASE_PORT,notEmpty"`
	MaxOpenConns int    `env:"DATABASE_MAX_OPENS_CONNS,notEmpty"`
	MaxIdleConns int    `env:"DATABASE_MAX_IDLE_CONNS,notEmpty"`
}
