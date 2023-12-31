// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: joinWaitlist.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const joinWaitList = `-- name: JoinWaitList :one
insert into waitlist (id, email, name, company_name, comment)
values ($1, $2, $3, $4, $5)
on conflict (email) do update set id = excluded.id, name = excluded.name, company_name = excluded.company_name, comment = excluded.comment
returning id, email, name, company_name, comment, joined
`

type JoinWaitListParams struct {
	ID          string      `json:"id"`
	Email       string      `json:"email"`
	Name        string      `json:"name"`
	CompanyName pgtype.Text `json:"company_name"`
	Comment     pgtype.Text `json:"comment"`
}

func (q *Queries) JoinWaitList(ctx context.Context, arg JoinWaitListParams) (Waitlist, error) {
	row := q.db.QueryRow(ctx, joinWaitList,
		arg.ID,
		arg.Email,
		arg.Name,
		arg.CompanyName,
		arg.Comment,
	)
	var i Waitlist
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Name,
		&i.CompanyName,
		&i.Comment,
		&i.Joined,
	)
	return i, err
}
