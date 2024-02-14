package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/gpbPiazza/cargo-api/src/infrastructure/envs"
)

type jwtService struct {
	signedKey []byte
}

func NewJwtService() *jwtService {
	accessEnv := envs.New().AccessEnv
	return &jwtService{
		signedKey: []byte(accessEnv.SignedKey),
	}
}

func (js *jwtService) Token(userID uuid.UUID, expirationAccessTime float64) (string, error) {
	accessExpiration := time.Now().Add(time.Duration(expirationAccessTime) * time.Second).Unix()

	claims := jwt.MapClaims{
		"user_id": userID.String(),
		"exp":     accessExpiration,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	return token.SignedString(js.signedKey)
}
