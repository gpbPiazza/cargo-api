package services

import (
	"context"
	"errors"

	"github.com/gpbPiazza/cargo-api/src/domain/models"
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
	finderCustomerRepository  usecases.FinderCustomerRepository
	hasherService             usecases.Hasher
	factory                   usecases.CustomerFactory
	creatorCustomerRepository usecases.CreatorCustomerRepository
}

func NewSignupService(
	findCustomer usecases.FinderCustomerRepository,
	hasherService usecases.Hasher,
	customerFactory usecases.CustomerFactory,
	creatorCustomer usecases.CreatorCustomerRepository) *signupService {
	return &signupService{
		finderCustomerRepository:  findCustomer,
		hasherService:             hasherService,
		factory:                   customerFactory,
		creatorCustomerRepository: creatorCustomer,
	}
}

func (ss *signupService) Register(ctx context.Context, params usecases.SignupParams) (models.Customer, error) {
	emptyCustomer := models.Customer{}
	if err := validator.Validate(ctx, params); err != nil {
		return emptyCustomer, err
	}

	bytesPassword := []byte(params.Password)
	if len(bytesPassword) > maxPasswordBytes {
		return emptyCustomer, ErrPasswordTooLong
	}

	customerFound, err := ss.finderCustomerRepository.FindByTaxID(ctx, params.TaxID)
	if err != nil && !errors.Is(err, errs.ErrNotFound) {
		return emptyCustomer, err
	}

	if customerFound.TaxID == params.TaxID {
		return emptyCustomer, ErrCustomerAlreadyRegistered
	}

	hashedPassword, err := ss.hasherService.Hash(params.Password)
	if err != nil {
		return emptyCustomer, err
	}

	customer := ss.factory.Make(params, hashedPassword)

	if err := ss.creatorCustomerRepository.Create(ctx, customer); err != nil {
		return emptyCustomer, err
	}

	return customer, nil
}
