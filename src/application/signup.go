package application

import (
	"context"

	"github.com/gpbPiazza/cargo-api/src/domain/models"
	"github.com/gpbPiazza/cargo-api/src/infrastructure/validator"
)

type SignupController struct {
}

type SignupCustomerParams struct {
	TaxID string              `json:"tax_id"`
	Name  string              `json:"name"`
	Role  models.CustomerRole `json:"role"`
	Phone string              `json:"phone"`
	Email string              `json:"email"`
}

func (sc *SignupController) Control(ctx context.Context, customerParams SignupCustomerParams) error {
	// call sanitizer.SanitizeTaxID
	// call sanitizer.SanitizePhone

	if err := validator.Validate(ctx, customerParams); err != nil {
		return err
	}

	// validação básica do userParams

	// chama o serviço de criar um costumer

	// return

	return nil
}
