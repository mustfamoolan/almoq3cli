package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	// Register all route files
	RegisterWebRoutes(app)
	RegisterAPIRoutes(app)
}
