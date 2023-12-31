// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: restoreTag.sql

package db

import (
	"context"
)

const restoreTag = `-- name: RestoreTag :one
update tag set deleted = null where id = $1 returning id, name, user_id, added, updated, deleted
`

func (q *Queries) RestoreTag(ctx context.Context, id string) (Tag, error) {
	row := q.db.QueryRow(ctx, restoreTag, id)
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
