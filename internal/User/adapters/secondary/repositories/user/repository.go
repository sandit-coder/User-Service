package repositories

import (
	"UserCrud/internal/User/application/ports/user"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) ports.UserRepository {
	return &UserRepository{
		db: db,
	}
}
