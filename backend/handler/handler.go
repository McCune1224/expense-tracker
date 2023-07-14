package handler

import "github.com/gofiber/fiber/v2"

type Handler struct{}

func New() *Handler {
    return &Handler{}
}

func (h *Handler) AttachRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello World",
		})
	})
}
