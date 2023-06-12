package models

import (
	"time"

	"github.com/google/uuid"
)

type EventType string

const (
	LOADING_HE   EventType = "LOADING_EVENT"
	UNLOADING_HE EventType = "UNLOADING_EVENT"
	// Sealing other types
	// storing other types
)

// A unique identifier for HandlingEvent is cargoID + completion_time + type event
type HandlingEvent struct {
	ID              uuid.UUID
	CompletionTime  time.Time
	Type            EventType
	Handled         *Cargo
	CarrierMovement *CarrierMovement
}
