package application

func NewEventLoggingController() Controller {
	return EventLoggingController{}
}

type EventLoggingController struct{}

func (controller EventLoggingController) Handle()

func (controller EventLoggingController) Logging() {
	// recebe parametros suficientes para
	// buscar uma carga
	// contruir um HandlingEvent com o tipo de handling
	// consultar o histórico de delivery da carga
	// Adiciona no delivery history este novo HandlingEvent
	// salva novo HandlingEvent e salva o delivery history
	// retorna usuário
}
