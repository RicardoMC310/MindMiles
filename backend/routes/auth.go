package routes

import (
	"backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app *fiber.App) {
	app.Post("/login", handlers.Login)
}