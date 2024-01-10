package application

import (
	"context"

	"github.com/gpbPiazza/cargo-api/src/domain/models"
	"github.com/gpbPiazza/cargo-api/src/infrastructure/validator"
)

type SignupApp struct {
}

type SignupParams struct {
	TaxID string              `json:"tax_id" validate:"required"`
	Name  string              `json:"name" validate:"required"`
	Role  models.CustomerRole `json:"role" validate:"required,oneof=SHIPPER RECEIVER"`
	Phone string              `json:"phone" validate:"required"`
	Email string              `json:"email" validate:"required"`
}

func (sc *SignupApp) Signup(ctx context.Context, customerParams SignupParams) error {
	// call sanitizer.SanitizeTaxID
	// call sanitizer.SanitizePhone
	// sanitizer.Sanitizer

	if err := validator.Validate(ctx, customerParams); err != nil {
		return err
	}

	// validação básica do userParams

	// chama o serviço de criar um costumer

	// return

	return nil
}
