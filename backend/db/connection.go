package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	util "github.com/kwandapchumba/bookmarkmonster/utils"
)

func ConnectDB() *pgx.Conn {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load config")
	}

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, config.DBString)
	if err != nil {
		log.Fatal(err)
	}

	if err := conn.Ping(ctx); err != nil {
		log.Fatal("db closed")
	}

	return conn
}
