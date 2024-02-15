package apps

import (
	"context"

	"github.com/gpbPiazza/cargo-api/src/domain/factories"
	"github.com/gpbPiazza/cargo-api/src/domain/models"
	"github.com/gpbPiazza/cargo-api/src/domain/services"
	"github.com/gpbPiazza/cargo-api/src/domain/usecases"
	"github.com/gpbPiazza/cargo-api/src/infrastructure/databases/sql/repositories"
	"github.com/gpbPiazza/cargo-api/src/infrastructure/hasher"
	"github.com/gpbPiazza/cargo-api/src/infrastructure/tokenizer"
)

type signupApp struct {
	signupService        usecases.SignupService
	authenticatorService usecases.AuthenticatorService
}

func NewSignupApp() *signupApp {
	customerRepository := repositories.NewCustomerRepository(nil)
	hasherService := hasher.NewService()
	tokenizerService := tokenizer.NewService()
	customerFactory := factories.NewCustomerFactory(nil)

	signupService := services.NewSignupService(customerRepository, hasherService, customerFactory, customerRepository)
	authenticatorService := services.NewAuthenticatorService(hasherService, tokenizerService)

	return &signupApp{
		signupService:        signupService,
		authenticatorService: authenticatorService,
	}
}

func (sa *signupApp) Signup(ctx context.Context, params usecases.SignupParams) (models.Customer, string, error) {
	// TODO: call sanitizer.SanitizeInput params

	customer, err := sa.signupService.Register(ctx, params)
	if err != nil {
		return models.Customer{}, "", err
	}

	accessToken, err := sa.authenticatorService.Authenticate(ctx, customer, params.Password)
	if err != nil {
		return models.Customer{}, "", err
	}

	return customer, accessToken, nil
}
