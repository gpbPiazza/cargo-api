package models

import "github.com/google/uuid"

// A unique identifier for HandlingEvent is cargoID + completion_time + type event
type HandlingEvent struct {
	ID              uuid.UUID
	Handled         *Cargo
	CarrierMovement *CarrierMovement
}
