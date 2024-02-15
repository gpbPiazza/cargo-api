package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gpbPiazza/cargo-api/src/application/apps"
	"github.com/gpbPiazza/cargo-api/src/application/usecases"
	domainUseCases "github.com/gpbPiazza/cargo-api/src/domain/usecases"
	"github.com/gpbPiazza/cargo-api/src/presentation/dtos"
)

type customerHandler struct {
	signupApp usecases.SignupApp
}

func NewCustomerHandler() *customerHandler {
	return &customerHandler{
		signupApp: apps.NewSignupApp(),
	}
}

func (ch *customerHandler) HandleSignup(c *fiber.Ctx) error {
	body := domainUseCases.SignupParams{}
	if err := c.BodyParser(&body); err != nil {
		// TODO: implement handlerErrFunction ch.HandleErr(err)
		return err
	}

	customer, accessToken, err := ch.signupApp.Signup(c.UserContext(), body)
	if err != nil {
		// TODO: implement handlerErrFunction ch.HandleErr(err)
		return err
	}

	return c.Status(http.StatusOK).JSON(dtos.NewSignupResponse(customer, accessToken))
}
