package handlers

import (
	"UserCrud/internal/User/application/dtos"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (handler *UserHandler) Update(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("id must be a valid UUID")
	}

	var req dtos.UpdateUserRequest

	if err := c.Bind().Body(&req); err != nil {
		return err
	}

	if err := handler.service.Update(c.Context(), &req, id); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
