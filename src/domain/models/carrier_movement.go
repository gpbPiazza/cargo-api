package models

import "github.com/google/uuid"

type CarrierMovement struct {
	ID   uuid.UUID
	From Location
	To   Location
}
