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
	finderCustomer usecases.FinderCustomerRepository
	hasher         usecases.Hasher
	factory        usecases.CustomerFactory
}

func NewSignupService(
	findCustomer usecases.FinderCustomerRepository,
	hasherService usecases.Hasher,
	customerFactory usecases.CustomerFactory) *signupService {
	return &signupService{
		finderCustomer: findCustomer,
		hasher:         hasherService,
		factory:        customerFactory,
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

	customerFound, err := ss.finderCustomer.FindByTaxID(ctx, params.TaxID)
	if err != nil && !errors.Is(err, errs.ErrNotFound) {
		return err
	}

	if customerFound.TaxID == params.TaxID {
		return ErrCustomerAlreadyRegistered
	}

	hashedPassword, err := ss.hasher.Hash(params.Password)
	if err != nil {
		return err
	}

	_ = ss.factory.Make(params, hashedPassword)

	// TODO: CALL REPOSITORY TO SAVE CUSTOMER MODEL

	return nil
}
