package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterAPIRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Example V1 Grouping
	v1 := api.Group("/v1")

	// Resource routes
	v1.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok", "version": "v1"})
	})

	// Add more routes here...
}
