package application

import (
	"context"

	"github.com/google/uuid"
	"github.com/gpbPiazza/cargo-api/src/application/dtos"
)

func NewEventLoggingController() Controller {
	return EventLoggingController{}
}

type EventLoggingController struct{}

// A camada de aplicação deveria conhcer o HTTP?
func (controller EventLoggingController) Control(req HttpRequest) HttpResponse {
	return HttpResponse{}
}

func (controller EventLoggingController) Logging(ctx context.Context, cargoID uuid.UUID, logEvent dtos.EventLoggingParams) dtos.EventLoggingResponse {

	// recebe parametros suficientes para
	// buscar uma carga

	// contruir um HandlingEvent com o tipo de handling
	// consultar o histórico de delivery da carga
	// Adiciona no delivery history este novo HandlingEvent
	// salva novo HandlingEvent e salva o delivery history
	// retorna usuário
	return dtos.MakeEventLoggingResponse()
}
