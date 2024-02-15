package dtos

import (
	"github.com/google/uuid"
	"github.com/gpbPiazza/cargo-api/src/domain/models"
)

type SignupResponse struct {
	CustomerID   uuid.UUID `json:"customer_id"`
	CustomerName string    `json:"customer_name"`
	AccessToken  string    `json:"access_token"`
}

func NewSignupResponse(customer models.Customer, accessToken string) SignupResponse {
	return SignupResponse{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		AccessToken:  accessToken,
	}
}
