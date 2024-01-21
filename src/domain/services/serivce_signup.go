package services

import (
	"context"
	"errors"

	"github.com/gpbPiazza/cargo-api/src/domain/usecases"
	"github.com/gpbPiazza/cargo-api/src/infrastructure/errs"
	"github.com/gpbPiazza/cargo-api/src/infrastructure/validator"
)

const (
	maxPasswordBytes = 72
)

var (
	ErrCustomerAlreadyRegistered = errs.New(errs.ErrValidation, "customer already registered")
	ErrPasswordTooLong           = errs.New(errs.ErrValidation, "password length exceeds 72 bytes")
)

type signupService struct {
	findCustomer  usecases.FindCustomerRepository
	hasherService usecases.Hasher
}

func NewSignupService(findCustomer usecases.FindCustomerRepository, hasherService usecases.Hasher) *signupService {
	return &signupService{
		findCustomer:  findCustomer,
		hasherService: hasherService,
	}
}

func (ss *signupService) Register(ctx context.Context, params usecases.SignupParams) error {
	if err := validator.Validate(ctx, params, validator.ValidateTaxID); err != nil {
		return err
	}

	bytesPassword := []byte(params.Password)
	if len(bytesPassword) > maxPasswordBytes {
		return ErrPasswordTooLong
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
