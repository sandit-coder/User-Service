package repositories

import (
	"UserCrud/internal/User/application/errors"
	"UserCrud/internal/User/domain/entities"
	"context"
	"database/sql"
	"errors"

	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (repo *UserRepository) GetById(ctx context.Context, tx pgx.Tx, id uuid.UUID) (*entities.User, error) {
	query := "SELECT id, first_name, last_name, email FROM users WHERE id = $1"

	var user entities.User

	if err := repo.db.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}
