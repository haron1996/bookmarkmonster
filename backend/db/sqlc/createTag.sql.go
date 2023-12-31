// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: createTag.sql

package db

import (
	"context"
)

const createTag = `-- name: CreateTag :one
INSERT INTO tag (id, name, user_id)
VALUES ($1, $2, $3)
RETURNING id, name, user_id, added, updated, deleted
`

type CreateTagParams struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	UserID string `json:"user_id"`
}

func (q *Queries) CreateTag(ctx context.Context, arg CreateTagParams) (Tag, error) {
	row := q.db.QueryRow(ctx, createTag, arg.ID, arg.Name, arg.UserID)
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
