package hasher

import (
	"golang.org/x/crypto/bcrypt"
)

const cost = 14

type hasherService struct {
}

func NewService() hasherService {
	return hasherService{}
}

func (hs hasherService) Hash(password string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(result), err
}

func (hs hasherService) CompareHash(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
