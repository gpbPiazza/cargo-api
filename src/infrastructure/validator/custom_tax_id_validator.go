package validator

import (
	"regexp"
	"strings"
)

const (
	cpfLen  = 11
	cnpjLen = 14
)

func isValidTaxID(taxID string) bool {
	if isCPF(taxID) {
		return validateCFP(taxID)
	}
	return validateCNPJ(taxID)
}

func isCPF(taxID string) bool {
	return len(clearMask(taxID)) == cpfLen
}

func clearMask(taxID string) string {
	taxIDCopy := taxID

	taxIDCopy = strings.Replace(taxIDCopy, ".", "", -1)
	taxIDCopy = strings.Replace(taxIDCopy, "-", "", -1)
	taxIDCopy = strings.Replace(taxIDCopy, "/", "", -1)
	taxIDCopy = strings.Replace(taxIDCopy, " ", "", -1)

	return taxIDCopy
}

func validateCFP(cpf string) bool {
	if !isValidCPFFormat(cpf) {
		return false
	}

	withoutMask := clearMask(cpf)

	if isAllEqual(withoutMask) {
		return false
	}

	return true
}

func isAllEqual(taxID string) bool {
	firstElement := taxID[0]
	for i := 1; i < len(taxID); i++ {
		if firstElement != taxID[i] {
			return false
		}
	}

	return true
}

func isValidCPFFormat(taxID string) bool {
	const (
		pattern = `^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`
	)

	return regexp.MustCompile(pattern).MatchString(taxID)
}

func validateCNPJ(cnpj string) bool {
	return false
}

func validateCNPJFormat(taxID string) bool {
	const (
		pattern = `^\d{2}\.?\d{3}\.?\d{3}\/?(:?\d{3}[1-9]|\d{2}[1-9]\d|\d[1-9]\d{2}|[1-9]\d{3})-?\d{2}$`
	)

	return regexp.MustCompile(pattern).MatchString(taxID)
}
