package repositories

import (
	"UserCrud/internal/user/application/errors"
	"UserCrud/internal/user/domain/entities"
	"UserCrud/internal/user/domain/errors"
	"context"
	"database/sql"
	"errors"

	"fmt"

	"github.com/google/uuid"
)

func (repo *UserRepository) GetById(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	query := "SELECT id, first_name, last_name, email FROM users WHERE id = $1"

	var user entities.User

	if err := repo.db.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.apperrors.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}
