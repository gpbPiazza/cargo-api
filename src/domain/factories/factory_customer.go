package customer

import (
	"github.com/gpbPiazza/cargo-api/src/domain/models"
	"github.com/gpbPiazza/cargo-api/src/domain/usecases"
)

type customerFactory struct {
	identifier usecases.TaxIDIdentifier
}

func NewCustomerFactory(identifier usecases.TaxIDIdentifier) *customerFactory {
	return &customerFactory{
		identifier: identifier,
	}
}

func (cf *customerFactory) Make(params usecases.SignupParams, password string) models.Customer {
	customerType := models.COMPANY_CT

	if cf.identifier.IdentifyPersonTaxID(params.TaxID) {
		customerType = models.PERSON_CT
	}

	result := models.Customer{
		TaxID: params.TaxID,
		Name:  params.Name,
		Type:  customerType,
		Role:  params.Role,
		Contact: models.CustomerContact{
			Phone: params.Phone,
			Email: params.Email,
		},
		Password: password,
	}

	return result
}
