package usecases

import "github.com/google/uuid"

type TokenizerService interface {
	Token(userID uuid.UUID, expirationAccessTime int) (string, error)
}
