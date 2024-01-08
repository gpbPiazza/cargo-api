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

type tagRule string

const ()

func (tr tagRule) string() string {
	return string(tr)
}

func (tr tagRule) customValidator() func(fl validator.FieldLevel) bool {
	validatorsMap := newCustomValidators()

	customValidator, ok := validatorsMap[tr]
	if !ok {
		customValidator = defaultValidatorAlwaysReturnErr
	}

	return customValidator
}

func newCustomValidators() map[tagRule]func(fl validator.FieldLevel) bool {
	return map[tagRule]func(fl validator.FieldLevel) bool{}
}

// func isValidContractStatus(fl validator.FieldLevel) bool {
// 	value := fl.Field().Interface()
// 	status := structs.ContractStatuses(fmt.Sprintf("%s", value))
// 	return status == "" || status == structs.OPEN_CS || status == structs.CANCELED_CS
// }

func defaultValidatorAlwaysReturnErr(fl validator.FieldLevel) bool {
	return false
}
