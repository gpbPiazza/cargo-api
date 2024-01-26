package suites

import "os"

type IntegrationBaseSuite struct {
	BaseSuite
}

func (ib *IntegrationBaseSuite) shouldRunIntTests() {
	if os.Getenv("INTEGRATION") == "" {
		ib.T().Skip("skiping integration test! Please set env variable 'INTEGRATION=INTEGRATION' to run integration tests")
	}
}

func (ib *IntegrationBaseSuite) SetupSuite() {
	ib.shouldRunIntTests()
}

func (ib *IntegrationBaseSuite) SetupSubTest() {

}
