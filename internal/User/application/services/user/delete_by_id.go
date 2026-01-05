package user

import (
	"context"

	"github.com/google/uuid"
)

func (service *Service) DeleteById(ctx context.Context, id uuid.UUID) error {
	return service.trm.Do(ctx, func(ctx context.Context) error {
		if err := service.repo.DeleteById(ctx, id); err != nil {
			return err
		}
		
		return nil
	})

}
