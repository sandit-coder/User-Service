package ports

import (
	"UserCrud/internal/user/application/dtos"
	"context"

	"github.com/google/uuid"
)

type UserService interface {
	GetById(ctx context.Context, id uuid.UUID) (entity *dtos.User, err error)
	Create(ctx context.Context, dto *dtos.User) (id uuid.UUID, err error)
	Update(ctx context.Context, dto *dtos.User, id uuid.UUID) (err error)
	DeleteById(ctx context.Context, id uuid.UUID) (err error)
}
