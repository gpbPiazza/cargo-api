package repositories

import (
	"os"
	"testing"
)

func TestIntegration_customerRepository_FindByTaxID(t *testing.T) {
	if os.Getenv("INTEGRATION") == "" {
		t.Skip("skiping integration test! Please set env variable 'INTEGRATION=INTEGRATION' to run integration tests")
	}

}
