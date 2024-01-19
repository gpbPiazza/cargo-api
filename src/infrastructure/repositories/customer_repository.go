package repositories

import (
	"context"

	"github.com/gpbPiazza/cargo-api/src/domain/models"
	"github.com/gpbPiazza/cargo-api/src/infrastructure/database"
)

type customerRepository struct{}

func NewCustomerRepository() customerRepository {
	return customerRepository{}
}

func (cr *customerRepository) FindByTaxID(ctx context.Context, taxID string) (models.Customer, error) {
	conn := database.ConnFromCtx(ctx)
	defer conn.Close()

	row := conn.QueryRowContext(ctx, "SELECT * FROM customers where taxID = %s", taxID)
	if row.Err() != nil {
		// TODO:
	}
	return models.Customer{}, nil
	// return row.Scan(&models.Customer{})
}
