package dtos

type EventLoggingRequest struct {
	Type string
}

type EventLoggingResponse struct {
}

func MakeEventLoggingResponse() EventLoggingResponse {
	return EventLoggingResponse{}
}
