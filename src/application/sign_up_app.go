package application

import (
	"context"

	"github.com/gpbPiazza/cargo-api/src/domain/models"
	"github.com/gpbPiazza/cargo-api/src/infrastructure/validator"
)

type SignupApp struct {
}

type SignupParams struct {
	TaxID    string              `json:"tax_id" validate:"required,validate-tax-id"`
	Name     string              `json:"name" validate:"required"`
	Role     models.CustomerRole `json:"role" validate:"required,oneof=SHIPPER RECEIVER"`
	Phone    string              `json:"phone" validate:"required"`
	Email    string              `json:"email" validate:"required,email"`
	Password string              `json:"password" validate:"required,lte=72,gte=8"` // TODO: IMPROVE PASSWORD VALIDATION require especial characters, min numbers, min capital letters
}

func (sc *SignupApp) Signup(ctx context.Context, customerParams SignupParams) error {
	// TODO: call sanitizer.SanitizeInput

	if err := validator.Validate(ctx, customerParams, validator.ValidateTaxID); err != nil {
		return err
	}

	// chama o serviço de criar um costumer

	// return

	return nil
}
