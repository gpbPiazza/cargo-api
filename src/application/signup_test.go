package application

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSignupController_Control(t *testing.T) {
	suite.Run(t, new(SignupControllerSuite))
}

type SignupControllerSuite struct {
	suite.Suite

	controller     SignupController
	costumerParams SignupCustomerParams
	ctx            context.Context
}

func (sc *SignupControllerSuite) SetupTest() {
	sc.controller = SignupController{}
	sc.ctx = context.Background()
	sc.costumerParams = SignupCustomerParams{
		TaxID: "93059283079",
		Name:  "my company",
		Role:  "SHIPPER",
		Phone: "7935919507",
		Email: "my_company@gmail.com",
	}
}

func (sc *SignupControllerSuite) TestShouldReturnErrWhenTaxIDIsNotProvided() {
	sc.costumerParams.TaxID = ""

	sc.Error(sc.controller.Control(sc.ctx, sc.costumerParams))
}

func (sc *SignupControllerSuite) TestShouldReturnErrWhenNameIsNotProvided() {
	sc.costumerParams.Name = ""

	sc.Error(sc.controller.Control(sc.ctx, sc.costumerParams))
}

func (sc *SignupControllerSuite) TestShouldReturnErrWhenRoleIsNotProvided() {
	sc.costumerParams.Role = ""

	sc.Error(sc.controller.Control(sc.ctx, sc.costumerParams))
}

func (sc *SignupControllerSuite) TestShouldReturnErrWhenPhoneIsNotProvided() {
	sc.costumerParams.Phone = ""

	sc.Error(sc.controller.Control(sc.ctx, sc.costumerParams))
}

func (sc *SignupControllerSuite) TestShouldReturnErrWhenEmailIsNotProvided() {
	sc.costumerParams.Phone = ""

	sc.Error(sc.controller.Control(sc.ctx, sc.costumerParams))
}

func (sc *SignupControllerSuite) TestShouldReturnErrWhenRoleIsNotOneOfShipperOrReceiver() {
	sc.costumerParams.Role = "ANY OTHER ROLE"

	sc.Error(sc.controller.Control(sc.ctx, sc.costumerParams))
}

func (sc *SignupControllerSuite) TestShouldReturnNoErrWhenAllIsOk() {
	sc.NoError(sc.controller.Control(sc.ctx, sc.costumerParams))
}
