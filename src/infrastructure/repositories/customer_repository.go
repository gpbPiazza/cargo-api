package repositories

import "github.com/google/uuid"

type CustomerRepository struct{}

func (cr CustomerRepository) FindByID(ID uuid.UUID)           {}
func (cr CustomerRepository) FindByName(name string)          {}
func (cr CustomerRepository) FindByCargoTrackingID(ID string) {}
