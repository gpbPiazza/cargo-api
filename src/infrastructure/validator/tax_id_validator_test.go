package validator

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func Test_isValidTaxID(t *testing.T) {
	suite.Run(t, new(isValidTaxIDSuite))
}

type isValidTaxIDSuite struct {
	suite.Suite
}

func (iv *isValidTaxIDSuite) TestCPFCases() {
	iv.Run("given valid CPF with mask should return true", func() {
		iv.True(isValidTaxID("120.051.489-03"))
	})

	iv.Run("given valid CPF with no mask should return true", func() {
		iv.True(isValidTaxID("12005148903"))
	})

	iv.Run("given valid CPF with wrong mask all dots should return false", func() {
		iv.False(isValidTaxID("120.051.489.03"))
	})

	iv.Run("given valid CPF with wrong mask - at beginning should return false", func() {
		iv.False(isValidTaxID("120-051.489.03"))
	})

	iv.Run("given invalid CPF 3467875434578764345789654 should return false", func() {
		iv.False(isValidTaxID("3467875434578764345789654"))
	})

	iv.Run("given no CPF should return false", func() {
		iv.False(isValidTaxID(""))
	})

	iv.Run("given only letters CPF should return false", func() {
		iv.False(isValidTaxID("AAAAAAAAAAA"))
	})

	iv.Run("given invalid CPF 000.000.000-00 should return false", func() {
		iv.False(isValidTaxID("000.000.000-00"))
	})

	iv.Run("given invalid CPF 111.111.111-11 should return false", func() {
		iv.False(isValidTaxID("111.111.111-11"))
	})

	iv.Run("given invalid CPF 222.222.222-22 should return false", func() {
		iv.False(isValidTaxID("222.222.222-22"))
	})

	iv.Run("given invalid CPF 333.333.333-33 should return false", func() {
		iv.False(isValidTaxID("333.333.333-33"))
	})

	iv.Run("given invalid CPF 444.444.444-44 should return false", func() {
		iv.False(isValidTaxID("444.444.444-44"))
	})

	iv.Run("given invalid CPF 555.555.555-55 should return false", func() {
		iv.False(isValidTaxID("555.555.555-55"))
	})

	iv.Run("given invalid CPF 666.666.666-66 should return false", func() {
		iv.False(isValidTaxID("666.666.666-66"))
	})

	iv.Run("given invalid CPF 777.777.777-77 should return false", func() {
		iv.False(isValidTaxID("777.777.777-77"))
	})

	iv.Run("given invalid CPF 888.888.888-88 should return false", func() {
		iv.False(isValidTaxID("888.888.888-88"))
	})

	iv.Run("given invalid CPF 999.999.999-99 should return false", func() {
		iv.False(isValidTaxID("999.999.999-99"))
	})

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
