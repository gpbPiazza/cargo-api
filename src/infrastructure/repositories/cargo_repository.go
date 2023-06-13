package repositories

import (
	"github.com/google/uuid"
)

type CargoRepository struct {
}

func (cp CargoRepository) FindByID(ID uuid.UUID) {
}

func (cp CargoRepository) FindByCargoTrackingID(ID string)       {}
func (cp CargoRepository) FindByCustomerID(customerID uuid.UUID) {}
