package repositories

import (
	"context"
	"errors"

	"github.com/gpbPiazza/cargo-api/src/domain/models"
	"github.com/gpbPiazza/cargo-api/src/infrastructure/databases/sql"
	"github.com/gpbPiazza/cargo-api/src/infrastructure/errs"

	"github.com/jackc/pgx"
)

type customerRepository struct{}

func NewCustomerRepository() customerRepository {
	return customerRepository{}
}

func (cr *customerRepository) FindByTaxID(ctx context.Context, taxID string) (models.Customer, error) {
	conn := sql.Conn(ctx)
	result := new(models.Customer)

	row, err := conn.Query(ctx, "SELECT * FROM customers where taxID = $1", taxID)
	if err != nil {
		return models.Customer{}, errs.NewUnexpected(err)
	}
	defer row.Close()

	err = row.Scan(result.ID, result.TaxID)

	if errors.Is(pgx.ErrNoRows, err) {
		return *result, errs.ErrNotFound
	}

	if err != nil {
		return models.Customer{}, errs.NewUnexpected(err)
	}

	return *result, nil
}

func (cr *customerRepository) Create(ctx context.Context, customer models.Customer) error {
	return nil
}
