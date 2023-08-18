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
		log.Panicf("could not load config: %v", err)
	}

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, config.DBString)
	if err != nil {
		log.Panicf("could not create a new pool: %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Panicf("ping unsuccessful: %v", err)
	}

	pool.Config().MaxConnIdleTime = 1 * time.Minute
	pool.Config().MaxConnLifetime = 5 * time.Minute
	pool.Config().MinConns = 150
	pool.Config().MaxConns = 100000
	pool.Config().MaxConnLifetimeJitter = 1 * time.Minute

	return pool
}
