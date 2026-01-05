package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (handler *UserHandler) GetById(c fiber.Ctx) error {

	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("id must be a valid UUID")
	}

	user, err := handler.service.GetById(c.Context(), id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
