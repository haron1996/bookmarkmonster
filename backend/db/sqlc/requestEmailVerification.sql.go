// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: requestEmailVerification.sql

package db

import (
	"context"
	"time"
)

const requestEmailVerification = `-- name: RequestEmailVerification :one
insert into email_verification(email, code, expiry, user_password)
values($1, $2, $3, $4)
on conflict (email) do update set code = excluded.code, expiry = excluded.expiry, user_password = excluded.user_password
returning email, code, expiry, user_password
`

type RequestEmailVerificationParams struct {
	Email        string    `json:"email"`
	Code         string    `json:"code"`
	Expiry       time.Time `json:"expiry"`
	UserPassword string    `json:"user_password"`
}

func (q *Queries) RequestEmailVerification(ctx context.Context, arg RequestEmailVerificationParams) (EmailVerification, error) {
	row := q.db.QueryRow(ctx, requestEmailVerification,
		arg.Email,
		arg.Code,
		arg.Expiry,
		arg.UserPassword,
	)
	var i EmailVerification
	err := row.Scan(
		&i.Email,
		&i.Code,
		&i.Expiry,
		&i.UserPassword,
	)
	return i, err
}