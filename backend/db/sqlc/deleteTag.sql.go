// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: deleteTag.sql

package db

import (
	"context"
)

const deleteTag = `-- name: DeleteTag :one
delete from tag where id = $1 returning id, name, user_id, added, updated, deleted
`

func (q *Queries) DeleteTag(ctx context.Context, id string) (Tag, error) {
	row := q.db.QueryRow(ctx, deleteTag, id)
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