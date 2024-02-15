package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gpbPiazza/cargo-api/src/presentation/http_api/handlers"
)

const customerRoutePrefix = "/customers"

func CustomerRoutes(d Dispatcher) {
	d.ClientRoute.Route(customerRoutePrefix, clientCustomersRoutes(), "client_customers_routes")
	d.InternalRoute.Route(customerRoutePrefix, internalCustomersRoutes(), "internal_customers_routes")
}

func clientCustomersRoutes() func(router fiber.Router) {
	customerHandler := handlers.NewCustomerHandler()

	return func(router fiber.Router) {
		router.Post("", customerHandler.HandleSignup)
	}
}

func internalCustomersRoutes() func(router fiber.Router) {
	return func(router fiber.Router) {
		router.Post("", nil)
	}
}
