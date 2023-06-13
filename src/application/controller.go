package application

import "context"

type HttpRequest struct {
	Ctx context.Context
}

type HttpResponse struct {
	StatusCode int
	Content    any
}

type Controller interface {
	Control(req HttpRequest) HttpResponse
}
