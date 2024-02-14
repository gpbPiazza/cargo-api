package services

import (
	"context"
	"time"

	"github.com/gpbPiazza/cargo-api/src/domain/models"
	"github.com/gpbPiazza/cargo-api/src/domain/usecases"
	"github.com/gpbPiazza/cargo-api/src/infrastructure/errs"
)

const (
	accessTokenExperitionTime = 3 * time.Hour
)

var (
	ErrMismatchedHashAndPassword = errs.New(errs.ErrUnauthorized, "hashedPassword is not the hash of the given password")
)

type authenticatorService struct {
	hasherService    usecases.Hasher
	tokenizerService usecases.TokenizerService
}

func NewAuthenticatorService(hasherService usecases.Hasher, tokenizerService usecases.TokenizerService) *authenticatorService {
	return &authenticatorService{
		hasherService:    hasherService,
		tokenizerService: tokenizerService,
	}
}

func (as *authenticatorService) Authenticate(ctx context.Context, customer models.Customer, password string) (string, error) {
	if err := as.hasherService.CompareHash(customer.Password, password); err != nil {
		return "", ErrMismatchedHashAndPassword
	}

	return as.tokenizerService.Token(customer.ID, accessTokenExperitionTime.Seconds())
}
