package models

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
	Metadata
	TaxID    string
	Name     string
	Password string
	Type     CustomerType
	Role     CustomerRole
	Contact  CustomerContact
}

// Object value
type CustomerContact struct {
	Phone string
	Email string
}
