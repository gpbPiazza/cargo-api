package repositories

import (
	"github.com/google/uuid"
	"github.com/gpbPiazza/cargo-api/src/domain/models"
)

type CarrierMovementRepository struct{}

func (cmr CarrierMovementRepository) FindByScheduleID(ID uuid.UUID)           {}
func (cmr CarrierMovementRepository) FindByLocation(from, to models.Location) {}
