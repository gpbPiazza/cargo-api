package envs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestNewEnvs(t *testing.T) {
	suite.Run(t, new(EnvsSuite))
}

type EnvsSuite struct {
	suite.Suite
}

func (es *EnvsSuite) SetupTest() {
	// tests.MockEnvs()
}

func (es *EnvsSuite) TestShouldPanicWhenSomeEnvWithNotEmptyTagIsNotSeted() {
	os.Clearenv()

	newEnvs := func() { _ = New() }

	es.Panics(newEnvs)
}

func (es *EnvsSuite) TestShouldRunOkWhenAllEnvsAreSeted() {
	newEnvs := func() {
		_ = New()
	}

	es.NotPanics(newEnvs)
}

func (es *EnvsSuite) TestShouldReturnTheSameEnvInstanceReferenceWhenCalledNewTwice() {
	expectedEnv := New()
	newEnv := New()

	es.Equal(expectedEnv, newEnv)
}
