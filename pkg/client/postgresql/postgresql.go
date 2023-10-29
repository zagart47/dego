package postgresql

import (
	"context"
	"dego/config"
	"dego/utils"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Client interface {
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, attempts int) (pool *pgxpool.Pool, err error) {
	server := config.NewConfig()
	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s", server.DB, server.DBUser, server.DBPwd, server.DBHost, server.DBPort, server.DBName)
	err = utils.ConnectWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.New(ctx, dsn)
		if err != nil {
			return err
		}
		return nil
	}, attempts, 5*time.Second)

	if err != nil {
		return nil, err
	}
	return pool, nil
}
