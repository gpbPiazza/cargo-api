package usecases

import (
	"context"

	"github.com/gpbPiazza/cargo-api/src/domain/models"
)

type CreatorCustomerRepository interface {
	Create(ctx context.Context, customer models.Customer) error
}
