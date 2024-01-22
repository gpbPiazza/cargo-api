package usecases

import (
	"context"

	"github.com/gpbPiazza/cargo-api/src/domain/models"
)

type SignupParams struct {
	TaxID    string              `json:"tax_id" validate:"required,validate-tax-id"`
	Name     string              `json:"name" validate:"required"`
	Role     models.CustomerRole `json:"role" validate:"required,oneof=SHIPPER RECEIVER"`
	Phone    string              `json:"phone" validate:"required"`
	Email    string              `json:"email" validate:"required,email"`
	Password string              `json:"password" validate:"required,lte=72,gte=8"` // TODO: IMPROVE PASSWORD VALIDATION require especial characters, min numbers, min capital letters
}

type SignupService interface {
	Register(ctx context.Context, params SignupParams) error
}
