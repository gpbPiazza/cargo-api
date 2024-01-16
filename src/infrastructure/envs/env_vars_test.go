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
	es.initMockDataBaseEnv()
	err := os.Setenv("SERVER_PORT", "9190")
	es.NoError(err)
}

func (es *EnvsSuite) initMockDataBaseEnv() {
	err := os.Setenv("DATABASE_NAME", "db_mock_name")
	es.NoError(err)

	err = os.Setenv("DATABASE_USERNAME", "db_mock_username")
	es.NoError(err)

	err = os.Setenv("DATABASE_PASSWORD", "db_mock_password")
	es.NoError(err)

	err = os.Setenv("DATABASE_HOST", "db_mock_host")
	es.NoError(err)

	err = os.Setenv("DATABASE_PORT", "db_mock_port")
	es.NoError(err)

	err = os.Setenv("DATABASE_MAX_OPENS_CONNS", "1")
	es.NoError(err)

	err = os.Setenv("DATABASE_MAX_IDLE_CONNS", "1")
	es.NoError(err)
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
