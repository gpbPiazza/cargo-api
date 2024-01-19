package services

import (
	"context"
	"errors"

	"github.com/gpbPiazza/cargo-api/src/domain/usecases"
	"github.com/gpbPiazza/cargo-api/src/infrastructure/errs"
	"github.com/gpbPiazza/cargo-api/src/infrastructure/validator"
)

var (
	ErrCustomerAlreadyRegistered = errs.New(errs.ErrValidation, "customer already registered")
)

type signupService struct {
	findCustomer usecases.FindCustomerRepository
}

func NewSignupService(findCustomer usecases.FindCustomerRepository) *signupService {
	return &signupService{
		findCustomer: findCustomer,
	}
}

func (ss *signupService) Register(ctx context.Context, params usecases.SignupParams) error {
	if err := validator.Validate(ctx, params, validator.ValidateTaxID); err != nil {
		return err
	}

	customer, err := ss.findCustomer.FindByTaxID(ctx, params.TaxID)
	if err != nil && !errors.Is(err, errs.ErrNotFound) {
		return err
	}

	if customer.TaxID == params.TaxID {
		return ErrCustomerAlreadyRegistered
	}

	// TODO: CREATE CUSTOMER_MODEL by "usecases.SignupParams"
	// TODO: CALL REPOSITORY TO SAVE CUSTOMER MODEL

	return nil
}
