package services

import (
	"context"
	"testing"

	"github.com/gpbPiazza/cargo-api/src/domain/models"
	"github.com/gpbPiazza/cargo-api/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

func Test_authenticatorService_Authenticate(t *testing.T) {
	suite.Run(t, new(authenticatorServiceSuite))
}

type authenticatorServiceSuite struct {
	suite.Suite

	hasherService    *mocks.MockHasher
	tokenizerService *mocks.MockTokenizerService

	notHashedPassword string
	serivce           *authenticatorService
	customer          models.Customer
	ctx               context.Context
}

func (as *authenticatorServiceSuite) SetupSubTest() {
	as.ctx = context.Background()

	ctrl := gomock.NewController(as.T(), gomock.WithOverridableExpectations())
	as.hasherService = mocks.NewMockHasher(ctrl)
	as.tokenizerService = mocks.NewMockTokenizerService(ctrl)

	as.notHashedPassword = "not_hashed_password"
	as.customer = models.Customer{
		TaxID:    "93059283079",
		Name:     "my company",
		Password: "hashed_password",
		Type:     models.COMPANY_CT,
		Role:     models.SHIPPER_CR,
		Contact: models.CustomerContact{
			Email: "my_company@gmail.com",
			Phone: "7935919507",
		},
	}
	as.customer.NewID()

	as.hasherService.
		EXPECT().
		CompareHash(as.customer.Password, as.notHashedPassword).
		AnyTimes().
		Return(nil)

	as.tokenizerService.
		EXPECT().
		Token(as.customer.ID, accessTokenExperitionTime.Seconds()).
		AnyTimes().
		Return("access_token", nil)

	as.serivce = NewAuthenticatorService(as.hasherService, as.tokenizerService)
}

func (as *authenticatorServiceSuite) TestSignup() {
	as.Run("should return access token and no err when all is ok", func() {
		got, err := as.serivce.Authenticate(as.ctx, as.customer, as.notHashedPassword)

		as.NoError(err)
		as.Equal("access_token", got)
	})

	as.Run("should return ErrMismatchedHashAndPassword when CompareHash return err", func() {
		as.hasherService.
			EXPECT().
			CompareHash(as.customer.Password, "wrong password").
			Return(assert.AnError)

		got, err := as.serivce.Authenticate(as.ctx, as.customer, "wrong password")

		as.Equal(ErrMismatchedHashAndPassword, err)
		as.Empty(got)
	})

	as.Run("should return when tokenizerservice returns err", func() {
		as.tokenizerService.
			EXPECT().
			Token(as.customer.ID, accessTokenExperitionTime.Seconds()).
			Return("", assert.AnError)

		got, err := as.serivce.Authenticate(as.ctx, as.customer, as.notHashedPassword)

		as.Error(err)
		as.Empty(got)
	})
}
