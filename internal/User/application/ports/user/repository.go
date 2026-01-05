package ports

import (
	"UserCrud/internal/User/domain/entities"
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	GetById(ctx context.Context, id uuid.UUID) (entity *entities.User, err error)
	Create(ctx context.Context, dto *entities.User) (id uuid.UUID, err error)
	Update(ctx context.Context, dto *entities.User) (err error)
	DeleteById(ctx context.Context, id uuid.UUID) (err error)
}
