package usecases

type Hasher interface {
	Hash(content string) (string, error)
	ValidateHash(hash string, content string) error
}
