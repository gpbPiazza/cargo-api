package models

// A unique identifier for HandlingEvent is cargoID + type event
type HandlingEvent struct {
	Handled Cargo
}
