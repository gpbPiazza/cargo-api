package models

import "github.com/google/uuid"

// Cargo aggregate
// Cargo agregate is everything that would noe exist bufo the parciular Cargo,
// which would include Role, DeliveryHistory and DeliveryEspecification.
// Each of this domains only are needed when the Cargo is present by him self.

type Role struct {
	// CUSTOMER ID TO HIS CARGO
	Customers map[uuid.UUID]Cargo
}

type Cargo struct {
	ID                     uuid.UUID
	TrackingID             uuid.UUID
	Role                   Role
	DeliveryHistory        DeliveryHistory
	DeliveryEspecification DeliveryEspecification
}
