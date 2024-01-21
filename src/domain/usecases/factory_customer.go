package usecases

import "github.com/gpbPiazza/cargo-api/src/domain/models"

type CustomerFactory interface {
	Make(params SignupParams, password string) models.Customer
}
