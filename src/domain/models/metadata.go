package models

import (
	"time"

	"github.com/google/uuid"
)

type Metadata struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (d *Metadata) NewID() {
	d.ID = uuid.New()
}
