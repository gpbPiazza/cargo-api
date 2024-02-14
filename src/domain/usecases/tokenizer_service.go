package usecases

import "github.com/google/uuid"

type TokenizerService interface {
	// Token generates a new JWT token with userID and experionTime properties.
	// The expirationAccessTime has seconds unit of time.
	Token(userID uuid.UUID, expirationAccessTime float64) (string, error)
}
