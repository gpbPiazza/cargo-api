package usecases

type TaxIDIdentifier interface {
	IdentifyPersonTaxID(taxID string) bool
}
