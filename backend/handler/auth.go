package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mccune1224/expense-tracker/model"
)

func (h *Handler) Login(c *fiber.Ctx) error {
	return c.JSON(
		fiber.Map{
			"message": "login",
		},
	)
}

func (h *Handler) Register(c *fiber.Ctx) error {
	registerUserReqBody := struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8,max=32,alphanum"`
	}{}
	if err := c.BodyParser(&registerUserReqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := h.v.Struct(registerUserReqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := h.us.CreateUser(
		model.User{
			Username: registerUserReqBody.Username,
			Email:    registerUserReqBody.Email,
			Password: registerUserReqBody.Password,
		},
	); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func (h *Handler) Logout(c *fiber.Ctx) error {
	return c.JSON(
		fiber.Map{
			"message": "login",
		},
	)
}
