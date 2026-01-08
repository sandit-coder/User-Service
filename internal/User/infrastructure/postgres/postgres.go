package postgres

import (
	"UserCrud/internal/user/config"
	"context"

	trmpg "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDb(cfg *config.DBConfig) (*pgxpool.Pool, error) {
	connectionString := config.GetPostgresConnectionString(cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	db, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewTransactionManager(pool *pgxpool.Pool) trm.Manager {
	return manager.Must(trmpg.NewDefaultFactory(pool))
}
