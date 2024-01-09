package validator

import (
	"context"
	"fmt"

	"github.com/go-playground/validator"
)

func Validate(ctx context.Context, object interface{}, tagRules ...tagRule) error {
	var err error
	validate := validator.New()

	for _, tagRule := range tagRules {
		err = validate.RegisterValidation(tagRule.string(), tagRule.customValidator())
		if err != nil {
			return err
		}
	}

	err = validate.StructCtx(ctx, object)
	if err == nil {
		return nil
	}

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}

	for _, validationError := range validationErrors {
		// TODO: IMPLEMENT ERR CONSTRUCTION
		fmt.Println(validationError)
	}

	return err
}
