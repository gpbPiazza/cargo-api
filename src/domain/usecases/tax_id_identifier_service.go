package usecases

type TaxIDIdentifierService interface {
	IdentifyPersonTaxID(taxID string) bool
}
