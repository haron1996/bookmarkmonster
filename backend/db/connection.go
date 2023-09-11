package db

import (
	_ "github.com/jackc/pgx/v5/stdlib"
)

// func ConnectDB() *pgxpool.Pool {

// 	config, err := utils.LoadConfig(".")
// 	if err != nil {
// 		log.Panicf("could not load config: %v", err)
// 	}

// 	ctx := context.Background()

// 	pool, err := pgxpool.New(ctx, config.DBString)
// 	if err != nil {
// 		log.Panicf("could not create a new pool: %v", err)
// 	}

// 	if err := pool.Ping(ctx); err != nil {
// 		log.Printf("ping unsuccessful: %v", err)
// 	}

// 	pool.Config().MaxConnIdleTime = 1 * time.Second
// 	pool.Config().MaxConnLifetime = 30 * time.Second
// 	pool.Config().MinConns = 0
// 	pool.Config().MaxConns = 1500
// 	pool.Config().HealthCheckPeriod = 1 * time.Second

// 	return pool
// }
