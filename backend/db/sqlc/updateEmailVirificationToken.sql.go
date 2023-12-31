// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: updateEmailVirificationToken.sql

package db

import (
	"context"
	"time"
)

const updateEmailVericationToken = `-- name: UpdateEmailVericationToken :exec
update email_verification set expiry = $1 where email = $2
`

type UpdateEmailVericationTokenParams struct {
	Expiry time.Time `json:"expiry"`
	Email  string    `json:"email"`
}

func (q *Queries) UpdateEmailVericationToken(ctx context.Context, arg UpdateEmailVericationTokenParams) error {
	_, err := q.db.Exec(ctx, updateEmailVericationToken, arg.Expiry, arg.Email)
	return err
}
