package handlers

import (
	"UserCrud/internal/user/adapters/request"
	"UserCrud/internal/user/application/dtos"

	"github.com/gofiber/fiber/v3"
)

func (handler *UserHandler) Create(c fiber.Ctx) error {
	var req request.CreateUserRequest

	if err := c.Bind().Body(&req); err != nil {
		return err
	}

	dto := &dtos.User{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	id, err := handler.service.Create(c.Context(), dto)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(id)
}
