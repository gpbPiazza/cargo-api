package repositories

import (
	"testing"
)

// TODO: MAKE a make comand to only run integration tests and unit tests
// using -short flag -> https://stackoverflow.com/questions/25965584/separating-unit-tests-and-integration-tests-in-go
// https://stackoverflow.com/questions/25965584/separating-unit-tests-and-integration-tests-in-go

func Test_customerRepository_FindByTaxID(t *testing.T) {
	if testing.Short() {
		t.Skip("skiping integration test")
	}
}
