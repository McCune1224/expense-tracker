package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) Login(c *fiber.Ctx) error    {
    h.us.GetUser()
}
func (h *Handler) Register(c *fiber.Ctx) error {
}
func (h *Handler) Logout(c *fiber.Ctx) error   {}
