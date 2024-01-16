package repositories

import (
	"context"

	"github.com/google/uuid"
)

type CustomerRepository interface {
	FindByID(ctx context.Context, ID uuid.UUID)
	FindByTaxID(ctx context.Context, taxID string)
	FindByName(ctx context.Context, name string)
	FindByCargoTrackingID(ctx context.Context, ID string)
}

type customerRepository struct{}

func NewCustomerRepository() CustomerRepository {
	return &customerRepository{}
}

func (cr *customerRepository) FindByTaxID(ctx context.Context, taxID string) {

}

func (cr *customerRepository) FindByID(ctx context.Context, ID uuid.UUID) {
	panic("implement me")
}

func (cr *customerRepository) FindByName(ctx context.Context, name string) {
	panic("implement me")
}

func (cr *customerRepository) FindByCargoTrackingID(ctx context.Context, ID string) {
	panic("implement me")
}
