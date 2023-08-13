package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
)

func GetUser(ctx context.Context, email, id string) (*sqlc.Userr, error) {
	q := sqlc.New(ConnectDB())

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
