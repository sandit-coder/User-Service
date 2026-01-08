package converters

import (
	"UserCrud/internal/user/application/dtos"
	"UserCrud/internal/user/domain/entities"
)

func ToDtoFromEntity(user *entities.User) *dtos.User {
	return &dtos.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}
