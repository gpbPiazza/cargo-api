package application

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSignUpApp_SignUp(t *testing.T) {
	suite.Run(t, new(SignUpAppSuite))
}

type SignUpAppSuite struct {
	suite.Suite

	app          SignUpApp
	SignUpParams SignUpParams
	ctx          context.Context
}

func (sa *SignUpAppSuite) SetupTest() {
	sa.app = SignUpApp{}
	sa.ctx = context.Background()
	sa.SignUpParams = SignUpParams{
		TaxID: "93059283079",
		Name:  "my company",
		Role:  "SHIPPER",
		Phone: "7935919507",
		Email: "my_company@gmail.com",
	}
}

func (sa *SignUpAppSuite) TestShouldReturnErrWhenTaxIDIsNotProvided() {
	sa.SignUpParams.TaxID = ""

	sa.Error(sa.app.SignUp(sa.ctx, sa.SignUpParams))
}

func (sa *SignUpAppSuite) TestShouldReturnErrWhenTaxIDIsInvalid() {
	sa.SignUpParams.TaxID = "000.000.000-00"

	sa.Error(sa.app.SignUp(sa.ctx, sa.SignUpParams))
}

func (sa *SignUpAppSuite) TestShouldReturnErrWhenNameIsNotProvided() {
	sa.SignUpParams.Name = ""

	sa.Error(sa.app.SignUp(sa.ctx, sa.SignUpParams))
}

func (sa *SignUpAppSuite) TestShouldReturnErrWhenRoleIsNotProvided() {
	sa.SignUpParams.Role = ""

	sa.Error(sa.app.SignUp(sa.ctx, sa.SignUpParams))
}

func (sa *SignUpAppSuite) TestShouldReturnErrWhenPhoneIsNotProvided() {
	sa.SignUpParams.Phone = ""

	sa.Error(sa.app.SignUp(sa.ctx, sa.SignUpParams))
}

func (sa *SignUpAppSuite) TestShouldReturnErrWhenEmailIsNotProvided() {
	sa.SignUpParams.Email = ""

	sa.Error(sa.app.SignUp(sa.ctx, sa.SignUpParams))
}

func (sa *SignUpAppSuite) TestShouldReturnErrWhenEmailIsNotValid() {
	sa.SignUpParams.Email = "my_email@.com"

	sa.Error(sa.app.SignUp(sa.ctx, sa.SignUpParams))
}

func (sa *SignUpAppSuite) TestShouldReturnErrWhenRoleIsNotOneOfShipperOrReceiver() {
	sa.SignUpParams.Role = "ANY OTHER ROLE"

	sa.Error(sa.app.SignUp(sa.ctx, sa.SignUpParams))
}

func (sa *SignUpAppSuite) TestShouldReturnNoErrWhenAllIsOk() {
	sa.NoError(sa.app.SignUp(sa.ctx, sa.SignUpParams))
}
