package routes

import (
	"backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func UsersRouter(app *fiber.App) {
	app.Get("/getUsers", handlers.GetAllUsers)
	app.Post("/createUser", handlers.CreateUser)
}