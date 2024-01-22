package validator

import (
	"context"
	"testing"

	"github.com/gpbPiazza/cargo-api/tests/suites"
	"github.com/stretchr/testify/suite"
)

type structToValidate struct {
	Name                   string `json:"first_user_name" validate:"required"`
	Email                  string `json:"my_custom_tag_json_email" validate:"required"`
	PropertyWithoutJSONTag string `json:"-" validate:"required"`
}

func TestValidate(t *testing.T) {
	suite.Run(t, new(validateSuite))
}

type validateSuite struct {
	suites.BaseSuite
	structToValidate structToValidate
}

func (v *validateSuite) SetupSubTest() {
	v.structToValidate = structToValidate{
		Name:                   "my name",
		Email:                  "my_email@gmail.com",
		PropertyWithoutJSONTag: "a field",
	}
}

func (v *validateSuite) TestValidate() {
	v.Run("should return err when required field is not provided", func() {
		v.structToValidate.Name = ""
		err := Validate(context.Background(), v.structToValidate)
		v.Error(err)
	})

	v.Run("should return err message with json tag when field has json tag", func() {
		v.structToValidate.Email = ""
		err := Validate(context.Background(), v.structToValidate)
		v.Require().Error(err)
		v.ErrorContainsAll(err, "my_custom_tag_json_email", "required")
	})

	v.Run("should return err message with struct name when field has not json tag", func() {
		v.structToValidate.PropertyWithoutJSONTag = ""
		err := Validate(context.Background(), v.structToValidate)
		v.Require().Error(err)
		v.ErrorContainsAll(err, "PropertyWithoutJSONTag", "required")
	})

	v.Run("should not return err when all fields are provided", func() {
		err := Validate(context.Background(), v.structToValidate)

		v.NoError(err)
	})
}
