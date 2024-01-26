package repositories

import (
	"context"
	"testing"

	"github.com/gpbPiazza/cargo-api/src/infrastructure/databases/sql"
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
	ctx := context.Background()
	ctx = sql.ConnectDB(ctx)
}

func (cr *customerRepositorySuite) TestCasesFindCustomerByTaxID() {
	panic("I am not in panic in make test/unit because of suites.IntegrationBaseSuite!")
}
