package handlers

import (
	"UserCrud/internal/User/application/ports/user"
)

type UserHandler struct {
	service ports.UserService
}

func NewUserHandler(service ports.UserService) *UserHandler {

	return &UserHandler{service: service}
}
