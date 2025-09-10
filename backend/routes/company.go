package routes

import (
	"backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func CompanyRoutes(app *fiber.App) {

	app.Get("/getCompanys", handlers.GetAllCompanys)

}