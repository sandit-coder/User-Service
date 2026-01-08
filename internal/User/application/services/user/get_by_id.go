package user

import (
	"UserCrud/internal/user/application/converters"
	"UserCrud/internal/user/application/dtos"
	"context"

	"github.com/google/uuid"
)

func (service *Service) GetById(ctx context.Context, id uuid.UUID) (*dtos.User, error) {
	var result *dtos.User

	err := service.trm.Do(ctx, func(ctx context.Context) error {
		entity, err := service.repo.GetById(ctx, id)
		if err != nil {
			return err
		}

		result = converters.ToDtoFromEntity(entity)
		return nil
	})

	return result, err
}
