package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterWebRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./resources/views/welcome.html")
	})
}
