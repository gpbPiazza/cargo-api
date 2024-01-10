package application

import (
	"context"

	"github.com/gpbPiazza/cargo-api/src/domain/models"
	"github.com/gpbPiazza/cargo-api/src/infrastructure/validator"
)

type SignUpApp struct {
}

type SignUpParams struct {
	TaxID string              `json:"tax_id" validate:"required,validate-tax-id"`
	Name  string              `json:"name" validate:"required"`
	Role  models.CustomerRole `json:"role" validate:"required,oneof=SHIPPER RECEIVER"`
	Phone string              `json:"phone" validate:"required"`
	Email string              `json:"email" validate:"required,email"`
}

func (sc *SignUpApp) SignUp(ctx context.Context, customerParams SignUpParams) error {
	// TODO: call sanitizer.SanitizeInput

	if err := validator.Validate(ctx, customerParams, validator.ValidateTaxID); err != nil {
		return err
	}

	// chama o serviço de criar um costumer

	// return

	return nil
}
