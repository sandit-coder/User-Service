package handlers

import (
	"UserCrud/internal/User/application/ports/user"

	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	validator *validator.Validate
	service   ports.UserService
}

func NewUserHandler(service ports.UserService) *UserHandler {

	return &UserHandler{service: service}
}
