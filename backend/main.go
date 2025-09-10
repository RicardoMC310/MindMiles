package main

import (
	"backend/config"
	"backend/routes"
	"backend/utils"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.Init()
	defer config.Close()

	app := fiber.New()
	app.Use(cors.New())

	routes.CompanyRoutes(app)

	port := utils.GetEnv("PORT")
	if port == "" {
		port = "8080" // valor padrão
	}

	err := app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}

}
