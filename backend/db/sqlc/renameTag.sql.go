// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: renameTag.sql

package db

import (
	"context"
)

const renameTag = `-- name: RenameTag :one
update tag set name = $1 where id = $2 and user_id = $3 and deleted is null returning id, name, user_id, added, updated, deleted
`

type RenameTagParams struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	UserID string `json:"user_id"`
}

func (q *Queries) RenameTag(ctx context.Context, arg RenameTagParams) (Tag, error) {
	row := q.db.QueryRow(ctx, renameTag, arg.Name, arg.ID, arg.UserID)
	var i Tag
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UserID,
		&i.Added,
		&i.Updated,
		&i.Deleted,
	)
	return i, err
}
