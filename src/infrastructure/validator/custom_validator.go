package validator

import (
	"github.com/go-playground/validator"
)

type tagRule string

var (
	customeValidatorsTagFunc = map[tagRule]func(fl validator.FieldLevel) bool{}
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
