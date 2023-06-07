package models

import "github.com/google/uuid"

type Location struct {
	ID        uuid.UUID
	Latitude  float64
	Logintude float64
	Address   Address
}

type Address struct {
	City                  string
	AdditionalInformation string
	Number                string
	StateCode             string
	Street                string
	Zip                   string
}
