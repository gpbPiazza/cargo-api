package application

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSignupApp_Signup(t *testing.T) {
	suite.Run(t, new(SignupAppSuite))
}

type SignupAppSuite struct {
	suite.Suite

	app          SignupApp
	signupParams SignupParams
	ctx          context.Context
}

func (sa *SignupAppSuite) SetupTest() {
	sa.app = SignupApp{}
	sa.ctx = context.Background()
	sa.signupParams = SignupParams{
		TaxID: "93059283079",
		Name:  "my company",
		Role:  "SHIPPER",
		Phone: "7935919507",
		Email: "my_company@gmail.com",
	}
}

func (sa *SignupAppSuite) TestShouldReturnErrWhenTaxIDIsNotProvided() {
	sa.signupParams.TaxID = ""

	sa.Error(sa.app.Signup(sa.ctx, sa.signupParams))
}

func (sa *SignupAppSuite) TestShouldReturnErrWhenTaxIDIsInvalid() {
	sa.signupParams.TaxID = "000.000.000-00"

	sa.Error(sa.app.Signup(sa.ctx, sa.signupParams))
}

func (sa *SignupAppSuite) TestShouldReturnErrWhenNameIsNotProvided() {
	sa.signupParams.Name = ""

	sa.Error(sa.app.Signup(sa.ctx, sa.signupParams))
}

func (sa *SignupAppSuite) TestShouldReturnErrWhenRoleIsNotProvided() {
	sa.signupParams.Role = ""

	sa.Error(sa.app.Signup(sa.ctx, sa.signupParams))
}

func (sa *SignupAppSuite) TestShouldReturnErrWhenPhoneIsNotProvided() {
	sa.signupParams.Phone = ""

	sa.Error(sa.app.Signup(sa.ctx, sa.signupParams))
}

func (sa *SignupAppSuite) TestShouldReturnErrWhenEmailIsNotProvided() {
	sa.signupParams.Email = ""

	sa.Error(sa.app.Signup(sa.ctx, sa.signupParams))
}

func (sa *SignupAppSuite) TestShouldReturnErrWhenEmailIsNotValid() {
	sa.signupParams.Email = "my_email@.com"

	sa.Error(sa.app.Signup(sa.ctx, sa.signupParams))
}

func (sa *SignupAppSuite) TestShouldReturnErrWhenRoleIsNotOneOfShipperOrReceiver() {
	sa.signupParams.Role = "ANY OTHER ROLE"

	sa.Error(sa.app.Signup(sa.ctx, sa.signupParams))
}

func (sa *SignupAppSuite) TestShouldReturnNoErrWhenAllIsOk() {
	sa.NoError(sa.app.Signup(sa.ctx, sa.signupParams))
}
