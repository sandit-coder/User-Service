package user

import (
	"UserCrud/internal/User/application/dtos"
	"UserCrud/internal/User/domain/entities"
	"context"

	"github.com/google/uuid"
)

func (service *Service) Update(ctx context.Context, request *dtos.UpdateUserRequest, id uuid.UUID) error {
	err := service.trm.Do(ctx, func(ctx context.Context) error {
		var entity = entities.NewUser(
			uuid.New(),
			request.Email,
			request.FirstName,
			request.LastName)

		if err := service.repo.Update(ctx, entity); err != nil {
			return err
		}
		return nil
	})
	
	return err
}
