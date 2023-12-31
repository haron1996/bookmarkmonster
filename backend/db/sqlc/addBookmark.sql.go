// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: addBookmark.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addBookmark = `-- name: AddBookmark :one
INSERT INTO bookmark (id, title, bookmark, host, favicon, thumbnail, user_id)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, title, bookmark, host, favicon, thumbnail, notes, user_id, added, updated, deleted, folder_id, beautified, fromchrome
`

type AddBookmarkParams struct {
	ID        string      `json:"id"`
	Title     string      `json:"title"`
	Bookmark  string      `json:"bookmark"`
	Host      string      `json:"host"`
	Favicon   pgtype.Text `json:"favicon"`
	Thumbnail pgtype.Text `json:"thumbnail"`
	UserID    string      `json:"user_id"`
}

func (q *Queries) AddBookmark(ctx context.Context, arg AddBookmarkParams) (Bookmark, error) {
	row := q.db.QueryRow(ctx, addBookmark,
		arg.ID,
		arg.Title,
		arg.Bookmark,
		arg.Host,
		arg.Favicon,
		arg.Thumbnail,
		arg.UserID,
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
		&i.Fromchrome,
	)
	return i, err
}
