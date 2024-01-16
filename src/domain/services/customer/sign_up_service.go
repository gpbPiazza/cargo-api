package customer

import "context"

type SignUpService interface {
	Register(ctx context.Context)
}

type signUpService struct {
}

func NewSignUpService() SignUpService {
	return &signUpService{}
}

// REQUIRIMENTS:
// -- consult by taxID if already exists, return err taxID customer already in use
// -- hash the password
// -- creates a model with the correct data
// --	save the new model in the database
func (cs *signUpService) Register(ctx context.Context) {

}
