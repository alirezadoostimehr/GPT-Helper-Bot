package postgres

import (
	"context"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ConnectionPool struct {
	*pgxpool.Pool
}

func NewConnectionPool(ctx context.Context, conf config.PostgresConfig) (*ConnectionPool, error) {

	pgxConfig, _ := pgxpool.ParseConfig("")
	pgxConfig.ConnConfig.Host = conf.Host
	pgxConfig.ConnConfig.Port = uint16(conf.Port)
	pgxConfig.ConnConfig.Database = conf.Database
	pgxConfig.ConnConfig.User = conf.User
	pgxConfig.ConnConfig.Password = conf.Password

	pool, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		panic(err)
	}

	err = pool.Ping(ctx)
	return &ConnectionPool{pool}, err
}
