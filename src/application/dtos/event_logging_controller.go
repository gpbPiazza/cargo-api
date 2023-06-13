package dtos

type EventLoggingParams struct {
	Type string
}

type EventLoggingResponse struct {
}

func MakeEventLoggingResponse() EventLoggingResponse {
	return EventLoggingResponse{}
}
