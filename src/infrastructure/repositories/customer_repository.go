package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/gpbPiazza/cargo-api/src/infrastructure/database"
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
	conn := database.ConnFromCtx(ctx)
	defer conn.Close()

	row := conn.QueryRowContext(ctx, "SELECT * FROM customers where taxID = %s", taxID)
	if row.Err() != nil {
		// TODO:
	}

	// return row.Scan(&models.Customer{})
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
