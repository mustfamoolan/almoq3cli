package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type OrderController struct {
	// Add dependencies here (e.g., Services)
}

func NewOrderController() *OrderController {
	return &OrderController{}
}

// Index - Get all resources
func (ctrl *OrderController) Index(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Index from OrderController",
	})
}

// Show - Get a single resource
func (ctrl *OrderController) Show(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{
		"message": "Show id " + id + " from OrderController",
	})
}

// Store - Create a new resource
func (ctrl *OrderController) Store(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Store in OrderController",
	})
}

// Update - Update an existing resource
func (ctrl *OrderController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{
		"message": "Update id " + id + " in OrderController",
	})
}

// Destroy - Delete a resource
func (ctrl *OrderController) Destroy(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{
		"message": "Destroy id " + id + " in OrderController",
	})
}
