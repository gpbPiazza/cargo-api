package application

import (
	"context"

	"github.com/gpbPiazza/cargo-api/src/domain/usecases"
)

type SignupApp struct {
	signupService usecases.SignupService
}

func (sc *SignupApp) Signup(ctx context.Context, params usecases.SignupParams) error {
	// TODO: call sanitizer.SanitizeInput params

	// chama o serviço de criar um costumer

	// chama o serviço de login

	// retorna um seção autenticada para o usuário

	return nil
}
