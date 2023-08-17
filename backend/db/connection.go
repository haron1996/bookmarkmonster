package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	util "github.com/kwandapchumba/bookmarkmonster/utils"
)

func ConnectDB() *pgxpool.Pool {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load config")
	}

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, config.DBString)
	if err != nil {
		log.Fatal(err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	pool.Config().MaxConnIdleTime = 30 * time.Minute
	pool.Config().MaxConnLifetime = 5 * time.Minute
	pool.Config().MinConns = 10
	pool.Config().MaxConns = 50

	return pool
}
