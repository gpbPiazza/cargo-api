package validator

import (
	"context"

	"github.com/go-playground/validator"
)

func Validate(ctx context.Context, object interface{}) error {
	validate := validator.New()

	for tag := range customValidatorsTags {
		if err := validate.RegisterValidation(tag.string(), tag.customValidator()); err != nil {
			return err
		}
	}

	return validate.StructCtx(ctx, object)
}
