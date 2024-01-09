package validator

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type structToValidate struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
}

func TestValidate(t *testing.T) {
	suite.Run(t, new(validateSuite))
}

type validateSuite struct {
	suite.Suite
	structToValidate structToValidate
}

func (v *validateSuite) SetupTest() {
	v.structToValidate = structToValidate{
		Name:  "my name",
		Email: "my_email@gmail.com",
	}
}

func (v *validateSuite) TestShouldReturnErrWhenRequiredFieldIsNotProvided() {
	v.structToValidate.Name = ""
	v.Error(Validate(context.Background(), v.structToValidate))
}

func (v *validateSuite) TestShouldNotReturnErrWhenAllFieldsAreProvided() {
	v.NoError(Validate(context.Background(), v.structToValidate))
}
