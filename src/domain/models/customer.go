package models

import "github.com/gpbPiazza/cargo-api/src/infrastructure/metadata"

type CustomerType string

const (
	COMPANY_CT CustomerType = "COMPANY"
	PERSON_CT  CustomerType = "PERSON"
)

type CustomerRole string

const (
	SHIPPER_CR  CustomerRole = "SHIPPER"
	RECEIVER_CR CustomerRole = "RECEIVER"
)

type Customer struct {
	metadata.Data
	TaxID   string
	Name    string
	Type    CustomerType
	Role    CustomerRole
	Contact CustomerContact
}

// Object value
type CustomerContact struct {
	Phone string
	Email string
}
