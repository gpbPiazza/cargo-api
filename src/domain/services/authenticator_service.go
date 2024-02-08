package services

import (
	"context"

	"github.com/gpbPiazza/cargo-api/src/domain/models"
	"github.com/gpbPiazza/cargo-api/src/domain/usecases"
	"github.com/gpbPiazza/cargo-api/src/infrastructure/errs"
)

var (
	ErrMismatchedHashAndPassword = errs.New(errs.ErrUnauthorized, "hashedPassword is not the hash of the given password")
)

type authenticatorService struct {
	hasherService usecases.Hasher
}

func NewAuthenticatorService(hasherService usecases.Hasher) *authenticatorService {
	return &authenticatorService{
		hasherService: hasherService,
	}
}

func (as *authenticatorService) Authenticate(ctx context.Context, customer models.Customer, password string) (string, error) {
	if err := as.hasherService.CompareHash(customer.Password, password); err != nil {
		return "", ErrMismatchedHashAndPassword
	}

	// TODO call tokenizer.AccessToken(ctx, customer.ID, experation)
	// TODO call save updateAccessToken(ctx, accessToken)

	return "", nil
}
