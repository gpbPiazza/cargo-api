package application

import (
	"context"
	"testing"

	"github.com/gpbPiazza/cargo-api/src/domain/usecases"
	"github.com/stretchr/testify/suite"
)

func TestSignupApp_Signup(t *testing.T) {
	suite.Run(t, new(SignupAppSuite))
}

type SignupAppSuite struct {
	suite.Suite

	app          SignupApp
	SignupParams usecases.SignupParams
	ctx          context.Context
}

func (sa *SignupAppSuite) SetupSubTest() {
	sa.app = SignupApp{}
	sa.ctx = context.Background()
	sa.SignupParams = usecases.SignupParams{
		TaxID:    "93059283079",
		Name:     "my company",
		Role:     "SHIPPER",
		Phone:    "7935919507",
		Email:    "my_company@gmail.com",
		Password: "1234567891020",
	}
}

func (sa *SignupAppSuite) TestSignup() {

}
