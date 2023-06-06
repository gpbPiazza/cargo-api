package models

import "github.com/google/uuid"

type CustomerRole string

const (
	COMPANY_CR CustomerRole = "COMPANY"
	PERSON_CR  CustomerRole = "PERSON"
)

type Customer struct {
	ID           uuid.UUID
	TaxID        string
	Name         string
	CustomerRole CustomerRole
}
