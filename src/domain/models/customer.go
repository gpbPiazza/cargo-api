package models

import "github.com/google/uuid"

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
	ID    uuid.UUID
	TaxID string
	Name  string
	Type  CustomerType
	Role  CustomerRole
}
