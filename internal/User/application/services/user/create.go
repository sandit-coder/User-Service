package user

import (
	"UserCrud/internal/User/application/dtos"
	"UserCrud/internal/User/domain/entities"
	"context"

	"github.com/google/uuid"
)

func (service *Service) Create(ctx context.Context, request *dtos.CreateUserRequest) (uuid.UUID, error) {
	var createdID uuid.UUID

	err := service.trm.Do(ctx, func(ctx context.Context) error {
		entity := entities.NewUser(
			uuid.New(),
			request.Email,
			request.FirstName,
			request.LastName,
		)

		id, err := service.repo.Create(ctx, entity)
		if err != nil {
			return err
		}

		createdID = id
		return nil
	})

	return createdID, err
}
