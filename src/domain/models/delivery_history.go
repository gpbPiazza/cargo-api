package models

import "github.com/google/uuid"

type DeliveryHistory struct {
	ID             uuid.UUID
	HandlingEvents []*HandlingEvent
}
