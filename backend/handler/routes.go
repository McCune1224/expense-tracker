package handler

import (
	"fmt"

	jwtware "github.com/gofiber/contrib/jwt"
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
	h.authRoutes(app)
	api := app.Group("/api")
	h.userRoutes(api)
	h.transactionRoutes(api)
	h.categoryRoutes(api)
}

func (h *Handler) authRoutes(app fiber.Router) {
	auth := app.Group("/auth")
	auth.Post("/login", h.Login)

	auth.Post("/register", h.Register)
	protectedAuth := auth.Group("/logout")
	protectedAuth.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("JWT_SECRET")},
	}))
	auth.Post("/logout", todo)
}

func (h *Handler) userRoutes(router fiber.Router) {
}
func (h *Handler) transactionRoutes(router fiber.Router) {}
func (h *Handler) categoryRoutes(router fiber.Router)    {}
