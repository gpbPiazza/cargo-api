package validator

import (
	"context"
	"reflect"
	"strings"

	"github.com/go-playground/validator"
)

var validate *validator.Validate

func init() {
	if validate != nil {
		return
	}

	validate = validator.New()

	validate.RegisterTagNameFunc(setJSONTagsIntoErrMessages)

	for tag := range customValidatorsTags {
		_ = validate.RegisterValidation(tag.string(), tag.customValidator())
		// TODO: DO SOMETHING HERE LOG? FATAL? IGNORE? DONT KNOW RIGHT NOW
	}
}

func setJSONTagsIntoErrMessages(field reflect.StructField) string {
	name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
	if name == "-" {
		return ""
	}
	return name
}

func Validate(ctx context.Context, object interface{}) error {
	return validate.StructCtx(ctx, object)
}
