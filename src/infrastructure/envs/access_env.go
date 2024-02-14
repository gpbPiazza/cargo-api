package envs

type accessEnv struct {
	SignedKey string `env:"SIGNED_KEY,notEmpty"`
}
