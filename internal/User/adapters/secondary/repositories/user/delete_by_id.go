package repositories

import (
	"UserCrud/internal/User/application/errors"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (repo *UserRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	query := "DELETE FROM users WHERE id = $1"

	commandTag, err := repo.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf(" user: %w", apperrors.ErrNotFound)
	}

	return nil
}
