package metadata

import (
	"time"

	"github.com/google/uuid"
)

type Data struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (d *Data) NewID() {
	d.ID = uuid.New()
}
