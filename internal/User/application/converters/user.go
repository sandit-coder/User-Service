package converters

import (
	"UserCrud/internal/User/application/dtos"
	"UserCrud/internal/User/domain/entities"
)

func ToDtoFromEntity(user *entities.User) *dtos.User {
	return &dtos.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}
