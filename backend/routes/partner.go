package routes

import (
	"backend/handlers"
	"backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func PartnerRouter(app *fiber.App) {
	app.Get("/getPartners", middlewares.JWTMiddleware, handlers.GetAllPartners)
	app.Post("/registerPartner", middlewares.JWTMiddleware, handlers.RegisterNewPartner)
}