package user

import (
	"UserCrud/internal/user/application/dtos"
	"UserCrud/internal/user/domain/entities"
	"context"

	"github.com/google/uuid"
)

func (service *Service) Update(ctx context.Context, request *dtos.User, id uuid.UUID) error {
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
