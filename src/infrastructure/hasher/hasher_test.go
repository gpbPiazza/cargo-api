package hasher

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

func TestHasher(t *testing.T) {
	suite.Run(t, new(HasherSuite))
}

type HasherSuite struct {
	suite.Suite
	hasherService hasherService
}

func (hs *HasherSuite) SetupSubTest() {
	hs.hasherService = NewHasherService()
}

func (hs *HasherSuite) TestCases() {
	hs.Run("should return err if content is bigget than 72 bytes", func() {
		content := string(make([]byte, 100))

		got, err := hs.hasherService.Hash(content)

		hs.Require().Error(err)
		hs.Empty(got)
		hs.Equal(bcrypt.ErrPasswordTooLong, err)
	})

	hs.Run("should not return err when Hash and validate Hash generated", func() {
		content := "myNotHashedPassword123"

		got, err := hs.hasherService.Hash(content)

		hs.Require().NoError(err)
		hs.Require().NotEmpty(got)
		hs.NoError(hs.hasherService.ValidateHash(got, content))
	})
}
