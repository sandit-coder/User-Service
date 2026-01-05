package handlers

import (
	"UserCrud/internal/User/application/dtos"

	"github.com/gofiber/fiber/v3"
)

func (handler *UserHandler) Create(c fiber.Ctx) error {
	var requestDto dtos.CreateUserRequest

	if err := c.Bind().Body(&requestDto); err != nil {
		return err
	}

	id, err := handler.service.Create(c.Context(), &requestDto)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(id)
}
