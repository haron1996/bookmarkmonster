// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: getUsers.sql

package db

import (
	"context"
)

const getUsers = `-- name: GetUsers :many
select id, name, email, email_verified, picture, account_password, created_at, last_login, refresh_token, deleted, signup_mode from userr order by last_login desc
`

func (q *Queries) GetUsers(ctx context.Context) ([]Userr, error) {
	rows, err := q.db.Query(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Userr{}
	for rows.Next() {
		var i Userr
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.EmailVerified,
			&i.Picture,
			&i.AccountPassword,
			&i.CreatedAt,
			&i.LastLogin,
			&i.RefreshToken,
			&i.Deleted,
			&i.SignupMode,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
