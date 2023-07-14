package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func todo(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"TODO": fmt.Sprintf("Implement Route: %s",
			c.Path(),
		),
	})
}

func (h *Handler) AttachRoutes(app *fiber.App) {
	// Root routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello World",
		})
	})

	api := app.Group("/api")
	authRoutes(api)
}

func authRoutes(app fiber.Router) {
	auth := app.Group("/auth", todo)
	auth.Get("/", todo)
	auth.Post("/login", todo)
	auth.Post("/logout", todo)
	auth.Post("/register", todo)
}
func userRoutes(router fiber.Router)        {}
func transactionRoutes(router fiber.Router) {}
func categoryRoutes(router fiber.Router)    {}
