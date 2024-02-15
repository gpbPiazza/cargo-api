package usecases

import (
	"context"

	"github.com/gpbPiazza/cargo-api/src/domain/models"
	"github.com/gpbPiazza/cargo-api/src/domain/usecases"
)

type SignupApp interface {
	Signup(ctx context.Context, params usecases.SignupParams) (models.Customer, string, error)
}
