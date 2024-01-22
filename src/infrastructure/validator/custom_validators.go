package validator

import (
	"fmt"

	"github.com/go-playground/validator"
)

type validatorTag string

const (
	tagValidateTaxID validatorTag = "validate-tax-id"
)

var (
	customValidatorsTags = map[validatorTag]func(fl validator.FieldLevel) bool{
		tagValidateTaxID: validateTaxID,
	}
)

func (tr validatorTag) string() string {
	return string(tr)
}

func (tr validatorTag) customValidator() func(fl validator.FieldLevel) bool {
	customValidator, ok := customValidatorsTags[tr]
	if !ok {
		customValidator = defaultValidatorAlwaysReturnErr
	}

	return customValidator
}

func defaultValidatorAlwaysReturnErr(fl validator.FieldLevel) bool {
	return false
}

func validateTaxID(fl validator.FieldLevel) bool {
	value := fl.Field().Interface()
	tax := fmt.Sprintf("%s", value)
	return isValidTaxID(tax)
}
