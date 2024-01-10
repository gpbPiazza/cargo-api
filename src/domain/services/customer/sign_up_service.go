package customer

import "context"

type CustomerSignUp interface {
	Register(ctx context.Context)
}

type CustomerSignUpService struct {
}

// REQUIRIMENTS:
// -- consult by taxID if already exists, return err taxID customer already in use
// -- hash the password
// -- creates a model with the correct data
// --	save the new model in the database
func (cs *CustomerSignUpService) Register(ctx context.Context) {

}
