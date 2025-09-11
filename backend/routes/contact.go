package routes

import (
	"backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func ContactsRouter(app *fiber.App) {
	app.Post("/", handlers.CreateContact)
}