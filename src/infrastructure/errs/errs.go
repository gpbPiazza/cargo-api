package errs

import (
	"errors"
	"fmt"
)

type ErrType error

var (
	ErrNotFound     ErrType = errors.New("not_found_error")
	ErrValidation   ErrType = errors.New("validation_error")
	ErrUnauthorized ErrType = errors.New("unauthorized_error")
	ErrUnexpected   ErrType = errors.New("unexpected_error")
	ErrUntyped      ErrType = errors.New("untyped_error")
)

func New(errType ErrType, message string) error {
	err := fmt.Errorf("%w: %s", errType, message)
	return err
}

// READY MORE ABOUT: https://cloud.google.com/apis/design/errors
// https://www.baeldung.com/rest-api-error-handling-best-practices
// https://developers.facebook.com/docs/whatsapp/cloud-api/support/error-codes/
// https://developer.mastercard.com/card-management/documentation/code-and-formats/error-codes/
// https://developers.docusign.com/docs/esign-rest-api/esign101/error-codes/

// todo: Implement http parser
// todo: implement more builders.
