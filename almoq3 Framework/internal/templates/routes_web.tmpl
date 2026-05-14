package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterWebRoutes(app *fiber.App) {
	// Serve static files from Vite build
	app.Static("/", "./public")

	// UI Showcase Route (Client-side handles this, but server must serve index.html)
	app.Get("/ui-showcase", func(c *fiber.Ctx) error {
		return c.SendFile("./public/index.html")
	})

	// Fallback to React for any other web routes
	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile("./public/index.html")
	})
}
