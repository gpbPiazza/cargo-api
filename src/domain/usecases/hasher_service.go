package usecases

type Hasher interface {
	Hash(password string) (string, error)
	CompareHash(hash string, password string) error
}
