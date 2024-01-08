package application

import (
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
}

func (sc *SignupControllerSuite) SetupTest() {
	sc.controller = SignupController{}
	sc.costumerParams = SignupCustomerParams{
		TaxID: "93059283079",
		Name:  "my company",
		Role:  "SHIPPER",
		Phone: "7935919507",
		Email: "my_company@gmail.com",
	}
}

func (sc *SignupControllerSuite) TestShouldReturnErrWhenTaxIDIsNotProvided() {

}
