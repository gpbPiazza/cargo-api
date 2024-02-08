package usecases

import (
	"context"

	"github.com/gpbPiazza/cargo-api/src/domain/models"
)

type AuthenticatorService interface {
	Authenticate(ctx context.Context, customer models.Customer, password string) (string, error)
}
