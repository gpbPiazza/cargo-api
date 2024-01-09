package validator

import (
	"fmt"

	"github.com/go-playground/validator"
)

type tagRule string

const (
	ValidateTaxID tagRule = "validate-tax-id"
)

var (
	customeValidatorsTagFunc = map[tagRule]func(fl validator.FieldLevel) bool{
		ValidateTaxID: validateTaxID,
	}
)

func (tr tagRule) string() string {
	return string(tr)
}

func (tr tagRule) customValidator() func(fl validator.FieldLevel) bool {
	customValidator, ok := customeValidatorsTagFunc[tr]
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
