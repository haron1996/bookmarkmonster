package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

func GetUser(ctx context.Context, email, id string) (*sqlc.Userr, error) {

	config, err := utils.LoadConfig(".")
	if err != nil {
		return nil, fmt.Errorf("could not load config: %v", err)
	}

	pool, err := pgxpool.New(ctx, config.DBString)
	if err != nil {
		return nil, fmt.Errorf("could not create new pool: %v", err)
	}

	defer pool.Close()

	q := sqlc.New(pool)

	getUserParams := sqlc.GetUserByEmailAndIDParams{
		Email: email,
		ID:    id,
	}

	var pgErr *pgconn.PgError

	user, err := q.GetUserByEmailAndID(ctx, getUserParams)
	if err != nil {
		if errors.As(err, &pgErr) {
			return nil, fmt.Errorf("could not get user: %v", err)
		} else if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user not found: %v", err)
		} else {
			return nil, fmt.Errorf("could not get user: %v", err)
		}
	}

	return &user, nil
}
