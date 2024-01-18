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

func (v *validateSuite) SetupSubTest() {
	v.structToValidate = structToValidate{
		Name:  "my name",
		Email: "my_email@gmail.com",
	}
}

func (v *validateSuite) TestValidate() {
	v.Run("should return err when required field is not provided", func() {
		v.structToValidate.Name = ""
		v.Error(Validate(context.Background(), v.structToValidate))
	})

	v.Run("should not return err when all fields are provided", func() {
		v.NoError(Validate(context.Background(), v.structToValidate))
	})
}
