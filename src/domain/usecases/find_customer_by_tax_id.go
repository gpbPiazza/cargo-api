package usecases

import (
	"context"

	"github.com/gpbPiazza/cargo-api/src/domain/models"
)

type FindCustomerRepository interface {
	FindByTaxID(ctx context.Context, taxID string) (models.Customer, error)
}
