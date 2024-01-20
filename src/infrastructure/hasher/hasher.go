package hasher

import "golang.org/x/crypto/bcrypt"

const cost = 14

type hasherService struct {
}

func NewHasherService() hasherService {
	return hasherService{}
}

func (hs hasherService) Hash(content string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(content), cost)
	return string(result), err
}

func (hs hasherService) ValidateHash(hash string, content string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(content))
}
