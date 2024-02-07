package repositories

import (
	"fmt"
	"testing"

	"github.com/gpbPiazza/cargo-api/tests/suites"
	"github.com/stretchr/testify/suite"
)

func TestIntegration_customerRepository(t *testing.T) {
	suite.Run(t, new(customerRepositorySuite))
}

type customerRepositorySuite struct {
	suites.IntegrationBaseSuite
}

func (cr *customerRepositorySuite) SetupSubTest() {
}

func (cr *customerRepositorySuite) TestCasesFindCustomerByTaxID() {
	fmt.Println("implement tests here")
}
