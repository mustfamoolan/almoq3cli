package main

import (
	"fmt"
	"log"

	"almoq3-config-app/bootstrap"
	"almoq3-config-app/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// 1. Initialize Configuration (FIRST)
	bootstrap.InitializeConfig()

	// 2. Initialize Fiber
	app := fiber.New(fiber.Config{
		AppName: config.Global.App.Name + " - almoq3 Framework",
	})

	// Basic route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("almoq3 Framework Started")
	})

	port := config.Global.App.Port
	if port == "" {
		port = "8080"
	}

	fmt.Printf("🚀 %s Started on port %s\n", config.Global.App.Name, port)
	log.Fatal(app.Listen(":" + port))
}
