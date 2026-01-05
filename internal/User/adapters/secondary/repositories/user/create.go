package repositories

import (
	"UserCrud/internal/User/application/errors"
	"context"
	"errors"
	"fmt"

	"UserCrud/internal/User/domain/entities"

	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func (repo *UserRepository) Create(ctx context.Context, entity *entities.User) (uuid.UUID, error) {
	query := `
        INSERT INTO users (id, email, first_name, last_name)
        VALUES ($1, $2, $3, $4)
    `

	_, err := repo.db.Exec(ctx, query, entity.ID, entity.Email, entity.FirstName, entity.LastName)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return uuid.Nil, fmt.Errorf("user: %w", apperrors.ErrAlreadyExists)
			}
		}

		return uuid.Nil, fmt.Errorf("repo create user: %w", err)
	}

	return entity.ID, nil
}
