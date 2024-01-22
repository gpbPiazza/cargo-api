package services

import (
	"context"
	"testing"

	"github.com/gpbPiazza/cargo-api/src/domain/models"
	"github.com/gpbPiazza/cargo-api/src/domain/usecases"
	"github.com/gpbPiazza/cargo-api/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

func Test_signupService_Register(t *testing.T) {
	suite.Run(t, new(signupServiceSuite))
}

type signupServiceSuite struct {
	suite.Suite

	serivce                *signupService
	findCustomerRepository *mocks.MockFindCustomerRepository
	hasherService          *mocks.MockHasher
	customerFactory        *mocks.MockCustomerFactory

	SignupParams usecases.SignupParams
	ctx          context.Context
}

func (ss *signupServiceSuite) SetupSubTest() {
	hashedPassword := "hashed_password"
	ss.SignupParams = usecases.SignupParams{
		TaxID:    "93059283079",
		Name:     "my company",
		Role:     models.SHIPPER_CR,
		Phone:    "7935919507",
		Email:    "my_company@gmail.com",
		Password: "1234567891020",
	}
	ss.ctx = context.Background()

	ctrl := gomock.NewController(ss.T(), gomock.WithOverridableExpectations())
	ss.findCustomerRepository = mocks.NewMockFindCustomerRepository(ctrl)
	ss.hasherService = mocks.NewMockHasher(ctrl)
	ss.customerFactory = mocks.NewMockCustomerFactory(ctrl)

	ss.findCustomerRepository.
		EXPECT().
		FindByTaxID(gomock.Any(), ss.SignupParams.TaxID).
		AnyTimes().
		Return(models.Customer{}, nil)

	ss.hasherService.
		EXPECT().
		Hash(ss.SignupParams.Password).
		AnyTimes().
		Return(hashedPassword, nil)

	ss.customerFactory.
		EXPECT().
		Make(ss.SignupParams, hashedPassword).
		AnyTimes().
		Return(models.Customer{})

	ss.serivce = NewSignupService(ss.findCustomerRepository, ss.hasherService, ss.customerFactory)
}

func (ss *signupServiceSuite) TestSignup() {
	ss.Run("should return err when tax id is not provided", func() {
		ss.SignupParams.TaxID = ""

		ss.Error(ss.serivce.Register(ss.ctx, ss.SignupParams))
	})

	ss.Run("should return err when tax id is invalid", func() {
		ss.SignupParams.TaxID = "000.000.000-00"

		ss.Error(ss.serivce.Register(ss.ctx, ss.SignupParams))
	})

	ss.Run("should return err when name is not provided", func() {
		ss.SignupParams.Name = ""

		ss.Error(ss.serivce.Register(ss.ctx, ss.SignupParams))
	})

	ss.Run("should return err when role is not provided", func() {
		ss.SignupParams.Role = ""

		ss.Error(ss.serivce.Register(ss.ctx, ss.SignupParams))
	})

	ss.Run("should return err when phone is not provided", func() {
		ss.SignupParams.Phone = ""

		ss.Error(ss.serivce.Register(ss.ctx, ss.SignupParams))
	})

	ss.Run("should return err when email is not provided", func() {
		ss.SignupParams.Email = ""

		ss.Error(ss.serivce.Register(ss.ctx, ss.SignupParams))
	})

	ss.Run("should return err when email is not valid", func() {
		ss.SignupParams.Email = "my_email@.com"

		ss.Error(ss.serivce.Register(ss.ctx, ss.SignupParams))
	})

	ss.Run("should return err when role is not one of shipper or receiver", func() {
		ss.SignupParams.Role = "ANY OTHER ROLE"

		ss.Error(ss.serivce.Register(ss.ctx, ss.SignupParams))
	})

	ss.Run("should return err when password is not provided", func() {
		ss.SignupParams.Password = ""

		ss.Error(ss.serivce.Register(ss.ctx, ss.SignupParams))
	})

	ss.Run("should return err if password is less than 8 chars", func() {
		ss.SignupParams.Password = "1234"

		ss.Error(ss.serivce.Register(ss.ctx, ss.SignupParams))
	})

	ss.Run("should return err if password is too long than 72 chars", func() {
		ss.SignupParams.Password = string(make([]byte, 80))

		err := ss.serivce.Register(ss.ctx, ss.SignupParams)

		ss.Require().Error(err)
		ss.Equal("Key: 'SignupParams.Password' Error:Field validation for 'Password' failed on the 'lte' tag", err.Error())
	})

	ss.Run("should return err if password bytes length longer than 72 bytes", func() {
		ss.SignupParams.Password = "😁😁😁😁😁😁😁😁😁😁😁😁😁😁😁😁😁😁😁"

		err := ss.serivce.Register(ss.ctx, ss.SignupParams)

		ss.Equal(ErrPasswordTooLong, err)
	})

	ss.Run("should return no err when all is ok", func() {
		ss.NoError(ss.serivce.Register(ss.ctx, ss.SignupParams))
	})

	ss.Run("should return err when findCustomerRepository returns any err that is not errNotFound", func() {
		ss.findCustomerRepository.
			EXPECT().
			FindByTaxID(gomock.Any(), ss.SignupParams.TaxID).
			Return(models.Customer{}, assert.AnError)

		err := ss.serivce.Register(ss.ctx, ss.SignupParams)

		ss.Require().Error(err)
		ss.Equal(assert.AnError, err)
	})

	ss.Run("should return err when hasherService returns any err", func() {
		ss.hasherService.
			EXPECT().
			Hash(ss.SignupParams.Password).
			Return("", assert.AnError)

		err := ss.serivce.Register(ss.ctx, ss.SignupParams)

		ss.Require().Error(err)
		ss.Equal(assert.AnError, err)
	})

	ss.Run("should return err when findCustomerRepository founds a customer", func() {
		ss.findCustomerRepository.
			EXPECT().
			FindByTaxID(gomock.Any(), ss.SignupParams.TaxID).
			Return(models.Customer{TaxID: ss.SignupParams.TaxID}, nil)

		got := ss.serivce.Register(ss.ctx, ss.SignupParams)

		ss.Error(got)
		ss.Equal(ErrCustomerAlreadyRegistered, got)
	})
}
