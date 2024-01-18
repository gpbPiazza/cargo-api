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
	SignupParams SignupParams
	ctx          context.Context
}

func (sa *SignupAppSuite) SetupSubTest() {
	sa.app = SignupApp{}
	sa.ctx = context.Background()
	sa.SignupParams = SignupParams{
		TaxID:    "93059283079",
		Name:     "my company",
		Role:     "SHIPPER",
		Phone:    "7935919507",
		Email:    "my_company@gmail.com",
		Password: "1234567891020",
	}
}

func (sa *SignupAppSuite) TestSignup() {
	sa.Run("should return err when tax id is not provided", func() {
		sa.SignupParams.TaxID = ""

		sa.Error(sa.app.Signup(sa.ctx, sa.SignupParams))
	})

	sa.Run("should return err when tax id is invalid", func() {
		sa.SignupParams.TaxID = "000.000.000-00"

		sa.Error(sa.app.Signup(sa.ctx, sa.SignupParams))
	})

	sa.Run("should return err when name is not provided", func() {
		sa.SignupParams.Name = ""

		sa.Error(sa.app.Signup(sa.ctx, sa.SignupParams))
	})

	sa.Run("should return err when role is not provided", func() {
		sa.SignupParams.Role = ""

		sa.Error(sa.app.Signup(sa.ctx, sa.SignupParams))
	})

	sa.Run("should return err when phone is not provided", func() {
		sa.SignupParams.Phone = ""

		sa.Error(sa.app.Signup(sa.ctx, sa.SignupParams))
	})

	sa.Run("should return err when email is not provided", func() {
		sa.SignupParams.Email = ""

		sa.Error(sa.app.Signup(sa.ctx, sa.SignupParams))
	})

	sa.Run("should return err when email is not valid", func() {
		sa.SignupParams.Email = "my_email@.com"

		sa.Error(sa.app.Signup(sa.ctx, sa.SignupParams))
	})

	sa.Run("should return err when role is not one of shipper or receiver", func() {
		sa.SignupParams.Role = "ANY OTHER ROLE"

		sa.Error(sa.app.Signup(sa.ctx, sa.SignupParams))
	})

	sa.Run("should return err when password is not provided", func() {
		sa.SignupParams.Password = ""

		sa.Error(sa.app.Signup(sa.ctx, sa.SignupParams))
	})

	sa.Run("should return err if password is less than 8 chars", func() {
		sa.SignupParams.Password = "1234"

		sa.Error(sa.app.Signup(sa.ctx, sa.SignupParams))
	})

	sa.Run("should return err if password is greated than 72 chars", func() {
		sa.SignupParams.Password = string(make([]byte, 80))

		sa.Error(sa.app.Signup(sa.ctx, sa.SignupParams))
	})

	sa.Run("should return no err when all is ok", func() {
		sa.NoError(sa.app.Signup(sa.ctx, sa.SignupParams))
	})
}
