// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: createBookmark.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createBookmark = `-- name: CreateBookmark :one
INSERT INTO bookmark (id, title, bookmark, host, user_id, notes, folder_id)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, title, bookmark, host, favicon, thumbnail, notes, user_id, added, updated, deleted, folder_id, beautified
`

type CreateBookmarkParams struct {
	ID       string      `json:"id"`
	Title    string      `json:"title"`
	Bookmark string      `json:"bookmark"`
	Host     string      `json:"host"`
	UserID   string      `json:"user_id"`
	Notes    pgtype.Text `json:"notes"`
	FolderID pgtype.Text `json:"folder_id"`
}

func (q *Queries) CreateBookmark(ctx context.Context, arg CreateBookmarkParams) (Bookmark, error) {
	row := q.db.QueryRow(ctx, createBookmark,
		arg.ID,
		arg.Title,
		arg.Bookmark,
		arg.Host,
		arg.UserID,
		arg.Notes,
		arg.FolderID,
	)
	var i Bookmark
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Bookmark,
		&i.Host,
		&i.Favicon,
		&i.Thumbnail,
		&i.Notes,
		&i.UserID,
		&i.Added,
		&i.Updated,
		&i.Deleted,
		&i.FolderID,
		&i.Beautified,
	)
	return i, err
}
