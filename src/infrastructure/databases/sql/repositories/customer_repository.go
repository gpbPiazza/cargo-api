package repositories

import (
	"context"
	"errors"

	"github.com/gpbPiazza/cargo-api/src/domain/models"
	"github.com/gpbPiazza/cargo-api/src/infrastructure/errs"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type customerRepository struct {
	dbConn *pgxpool.Pool
}

func NewCustomerRepository(poolConextion *pgxpool.Pool) *customerRepository {
	return &customerRepository{
		dbConn: poolConextion,
	}
}

func (cr *customerRepository) FindByTaxID(ctx context.Context, taxID string) (models.Customer, error) {
	result := new(models.Customer)

	row, err := cr.dbConn.Query(ctx, "SELECT * FROM customers where taxID = $1", taxID)
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
