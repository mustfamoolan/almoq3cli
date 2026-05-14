package main

import (
	"fmt"
	"log"

	"almoq3-stage6-test/bootstrap"
	"almoq3-stage6-test/config"
	"almoq3-stage6-test/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// 1. Initialize Configuration (FIRST)
	bootstrap.InitializeConfig()

	// 2. Initialize Database
	bootstrap.InitializeDatabase()

	// 3. Initialize Fiber
	app := fiber.New(fiber.Config{
		AppName: config.Global.App.Name + " - almoq3 Framework",
	})

	// 4. Register Routes
	routes.Register(app)

	port := config.Global.App.Port
	if port == "" {
		port = "8080"
	}

	fmt.Printf("🚀 %s Started on port %s\n", config.Global.App.Name, port)
	log.Fatal(app.Listen(":" + port))
}
