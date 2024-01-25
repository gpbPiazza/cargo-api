package repositories

import (
	"context"
	goSQL "database/sql"
	"errors"

	"github.com/gpbPiazza/cargo-api/src/domain/models"
	"github.com/gpbPiazza/cargo-api/src/infrastructure/databases/sql"
	"github.com/gpbPiazza/cargo-api/src/infrastructure/errs"
)

type customerRepository struct{}

func NewCustomerRepository() customerRepository {
	return customerRepository{}
}

func (cr *customerRepository) FindByTaxID(ctx context.Context, taxID string) (models.Customer, error) {
	conn := sql.ConnFromCtx(ctx)
	defer conn.Close()
	result := new(models.Customer)

	row := conn.QueryRowContext(ctx, "SELECT * FROM customers where taxID = $1", taxID)
	if row.Err() != nil {
		return models.Customer{}, errs.NewUnexpected(row.Err())
	}

	err := row.Scan(result)
	if errors.Is(goSQL.ErrNoRows, err) {
		return *result, errs.ErrNotFound
	}

	if err != nil {
		return *result, errs.NewUnexpected(err)
	}

	return *result, nil
}

func (cr *customerRepository) Create(ctx context.Context, customer models.Customer) error {
	return nil
}
