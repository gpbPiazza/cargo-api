package customer

import (
	"testing"

	"github.com/gpbPiazza/cargo-api/src/domain/models"
	"github.com/gpbPiazza/cargo-api/src/domain/usecases"
	"github.com/gpbPiazza/cargo-api/tests/mocks"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

func TestCustomerFactory(t *testing.T) {
	suite.Run(t, new(CustomerFactorySuite))
}

type CustomerFactorySuite struct {
	suite.Suite

	factory    customerFactory
	identifier *mocks.MockTaxIDIdentifier
	params     usecases.SignupParams
	customer   models.Customer
	password   string
}

func (cf *CustomerFactorySuite) SetupSubTest() {
	ctrl := gomock.NewController(cf.T())
	cf.identifier = mocks.NewMockTaxIDIdentifier(ctrl)

	cf.params = usecases.SignupParams{
		TaxID:    "93059283079",
		Name:     "my company",
		Role:     models.SHIPPER_CR,
		Phone:    "7935919507",
		Email:    "my_company@gmail.com",
		Password: "1234567891020",
	}
	cf.password = "my_hashed_passowrd"
	cf.customer = models.Customer{
		TaxID: cf.params.TaxID,
		Name:  cf.params.Name,
		Type:  models.COMPANY_CT,
		Role:  cf.params.Role,
		Contact: models.CustomerContact{
			Phone: cf.params.Phone,
			Email: cf.params.Email,
		},
		Password: cf.password,
	}

	cf.factory = *NewCustomerFactory(cf.identifier)
}

func (cf *CustomerFactorySuite) TestCasesMake() {
	cf.Run("should return a customer with type company when TaxIDIdentifier returns false", func() {
		expected := cf.customer
		expected.Type = models.COMPANY_CT

		cf.identifier.EXPECT().IdentifyPersonTaxID(cf.params.TaxID).Return(false)

		got := cf.factory.Make(cf.params, cf.password)

		cf.Require().NotEmpty(got)
		cf.Equal(models.COMPANY_CT, got.Type)
		cf.Equal(expected, got)
	})

	cf.Run("should return a customer with type person when TaxIDIdentifier returns true", func() {
		expected := cf.customer
		expected.Type = models.PERSON_CT

		cf.identifier.EXPECT().IdentifyPersonTaxID(cf.params.TaxID).Return(true)

		got := cf.factory.Make(cf.params, cf.password)

		cf.Require().NotEmpty(got)
		cf.Equal(models.PERSON_CT, got.Type)
		cf.Equal(expected, got)
	})
}
