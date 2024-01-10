package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func Test_isValidTaxID(t *testing.T) {
	suite.Run(t, new(isValidTaxIDSuite))
}

type isValidTaxIDSuite struct {
	suite.Suite
}

func (iv *isValidTaxIDSuite) TestCPFCases() {
	tests := []struct {
		name   string
		taxID  string
		assert assert.BoolAssertionFunc
	}{
		{
			name:   "given valid CPF with mask should return true",
			taxID:  "120.051.489-03",
			assert: assert.True,
		},
		{
			name:   "given valid CPF with no mask should return true",
			taxID:  "12005148903",
			assert: assert.True,
		},
		{
			name:   "given valid CPF with wrong mask all dots should return false",
			taxID:  "120.051.489.03",
			assert: assert.False,
		},
		{
			name:   "given valid CPF with wrong mask - at beginning should return false",
			taxID:  "120-051.489.03",
			assert: assert.False,
		},
		{
			name:   "given invalid CPF 3467875434578764345789654 should return false",
			taxID:  "3467875434578764345789654",
			assert: assert.False,
		},
		{
			name:   "given no CPF should return false",
			taxID:  "",
			assert: assert.False,
		},
		{
			name:   "given no CPF should return false",
			taxID:  "AAAAAAAAAAA",
			assert: assert.False,
		},
		{
			name:   "given invalid CPF 000.000.000-00 should return false",
			taxID:  "000.000.000-00",
			assert: assert.False,
		},
		{
			name:   "given invalid CPF 111.111.111-11 should return false",
			taxID:  "111.111.111-11",
			assert: assert.False,
		},
		{
			name:   "given invalid CPF 222.222.222-22 should return false",
			taxID:  "222.222.222-22",
			assert: assert.False,
		},
		{
			name:   "given invalid CPF 333.333.333-33 should return false",
			taxID:  "333.333.333-33",
			assert: assert.False,
		},
		{
			name:   "given invalid CPF 444.444.444-44 should return false",
			taxID:  "444.444.444-44",
			assert: assert.False,
		},
		{
			name:   "given invalid CPF 555.555.555-55 should return false",
			taxID:  "555.555.555-55",
			assert: assert.False,
		},
		{
			name:   "given invalid CPF 666.666.666-66 should return false",
			taxID:  "666.666.666-66",
			assert: assert.False,
		},
		{
			name:   "given invalid CPF 777.777.777-77 should return false",
			taxID:  "777.777.777-77",
			assert: assert.False,
		},
		{
			name:   "given invalid CPF 888.888.888-88 should return false",
			taxID:  "888.888.888-88",
			assert: assert.False,
		},
		{
			name:   "given invalid CPF 999.999.999-99 should return false",
			taxID:  "999.999.999-99",
			assert: assert.False,
		},
		// TODO: IMPLEMENT CPF CALCULATION
		// {
		// 	name:   "given invalid CPF 248.438.034-08 should return false",
		// 	taxID:  "248.438.034-08",
		// 	assert: assert.False,
		// },
		// {
		// 	name:   "given invalid CPF 099.075.865-06 should return false",
		// 	taxID:  "099.075.865-06",
		// 	assert: assert.False,
		// },
		// {
		// 	name:   "given invalid CPF 248 438 034 80 should return false",
		// 	taxID:  "248 438 034 80",
		// 	assert: assert.False,
		// },
		// {
		// 	name:   "given invalid CPF 099-075-865.60 should return false",
		// 	taxID:  "099-075-865.60",
		// 	assert: assert.False,
		// },
	}

	for _, tt := range tests {
		iv.Run(tt.name, func() {
			tt.assert(iv.T(), isValidTaxID(tt.taxID))
		})
	}
}
