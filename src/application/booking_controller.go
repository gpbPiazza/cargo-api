package application

type BookingController struct{}

func NewBookingController() Controller {
	return BookingController{}
}

func (bc BookingController) Control(req HttpRequest) HttpResponse {
	return HttpResponse{}
}
