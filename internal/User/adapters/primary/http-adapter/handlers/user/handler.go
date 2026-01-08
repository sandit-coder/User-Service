package handlers

import (
	"UserCrud/internal/user/application/ports/user"
)

type UserHandler struct {
	service ports.UserService
}

func NewUserHandler(service ports.UserService) *UserHandler {

	return &UserHandler{service: service}
}
