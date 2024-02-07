package application

import (
	"context"

	"github.com/gpbPiazza/cargo-api/src/domain/usecases"
)

type SignupApp struct {
	signupService usecases.SignupService
}

func (sa *SignupApp) Signup(ctx context.Context, params usecases.SignupParams) error {
	// TODO: call sanitizer.SanitizeInput params

	if err := sa.signupService.Register(ctx, params); err != nil {
		return err
	}

	// chama o serviço de login

	// retorna um seção autenticada para o usuário

	return nil
}
