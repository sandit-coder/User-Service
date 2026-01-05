package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (handler *UserHandler) DeleteById(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("id must be a valid UUID")
	}

	if err := handler.service.DeleteById(c.Context(), id); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
